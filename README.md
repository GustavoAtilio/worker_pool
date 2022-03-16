# Padrão Worker Pool em Golang
![GitHub](https://img.shields.io/github/license/GustavoAtilio/worker_pool)

## Como Funciona
#### Algoritmo Utiliza o padrão worker pool do golang para criptografar arquivos
#### Usando função recursiva para pegar os caminhos 
```
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

```
## Como Utilizar em seu exemplo
 
#### Set a variavel path_default := "./teste"
```
go run main.go 
```
