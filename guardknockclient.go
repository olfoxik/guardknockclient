package main


import "gopkg.in/ini.v1"
import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/base64"
import "os/exec"
 

func main() {

   cfg, err := ini.Load("config.ini")
   username := cfg.Section("Main").Key("username").String()
   password := cfg.Section("Main").Key("password").String()
 
   url := cfg.Section("Main").Key("url").String()

 
   fmt.Printf("Welcom guardknockclient v 1.0 \n")
   fmt.Printf("CONNECT ....")    

    // Создание HTTP клиента
    client := &http.Client{}

    // Формирование запроса
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Ошибка при создании запроса:", err)
        return
    }

    // Добавление заголовка для Basic Auth
    auth := username + ":" + password
    basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
    req.Header.Add("Authorization", basicAuth)

    // Выполнение запроса
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Ошибка при выполнении запроса:", err)
        return
    }
    defer resp.Body.Close()

    // Чтение ответа
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Ошибка при чтении ответа:", err)
        return
    }

    // Вывод ответа
 fmt.Println("Ответ от сервера:", string(body))

cmd := exec.Command("cmd", "/c", "knock.exe -d 10 " + string(body))
stdoutStderr, err := cmd.CombinedOutput()

fmt.Printf("OK \n")    

fmt.Println(stdoutStderr, err) 
fmt.Println("Ответ от сервера:", string(body))
}
