package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func encrypt(data []byte) ([]byte, error) {
	key := []byte("6z2n4Fd9eriJGVL9du3dIaTS3MVx1BBj")
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	nonce := []byte("eHNwwe9MiEdd")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	return ciphertext, nil
}

func decrypt(data []byte) ([]byte, error) {
	key := []byte("6z2n4Fd9eriJGVL9du3dIaTS3MVx1BBj")
	nonce := []byte("eHNwwe9MiEdd")
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	bin, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}
	return bin, nil
}

func encripFile(ch <-chan string, wait *sync.WaitGroup, option int) {
	var data []byte
	for file := range ch {
		fmt.Println(file)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if option == 1 {
			data, err = encrypt(content)
		} else {
			data, err = decrypt(content)
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		file, err := os.Create(file)
		err = file.Truncate(0)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		file.Write(data)

		file.Close()
	}
	wait.Done()
}

//Faz uma busca Recursiva por Caminhos
func getPath(path string) []string {
	var paths []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		if f.IsDir() {
			path := fmt.Sprintf("%s/%s", path, f.Name())
			re := getPath(path)
			paths = append(paths, re...)
		} else {
			//fmt.Println(fmt.Sprintf("%s\\%s", path, f.Name()))
			paths = append(paths, fmt.Sprintf("%s/%s", path, f.Name()))
		}

	}

	return paths
}

func main() {
	//fmt.Println(time.Now())

	works := 3  //Quantidade de Trabalhadores
	option := 0 //opção de entrafa
	var waitGrupo sync.WaitGroup
	waitGrupo.Add(works)
	path_default := "./teste" // Caminho de Origem

	data := getPath(path_default) //Buscar caminhos a partir do caminho de origem
	fmt.Println("Quantidades de arquivos a serem Criptografados ", len(data))
	fmt.Println("Digite 1 Para Criptografar 0 para descriptogravar")
	fmt.Print(">> ")
	fmt.Scanf("%d", &option)
	//Inicia os Trabalhadores
	ch := make(chan string, 10)
	for i := 0; i < works; i++ {
		go encripFile(ch, &waitGrupo, option)
	}
	//Passa os caminhos para o canal
	for _, x := range data {
		ch <- x
	}
	close(ch) //Fecha Canal finaliza as goroutines

	waitGrupo.Wait()
	fmt.Println("Processo Finalizado")

	//fmt.Println(time.Now())
}
