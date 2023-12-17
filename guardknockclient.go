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

 
 
    // Замените URL на целевой REST API
    //url := "http://95.140.159.65:3004/"
      url := cfg.Section("Main").Key("url").String()
    // Замените вашими учетными данными
 //   username := "olfox2"
  //  password := "tuxpux7"

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
 exec.Command("D:\golang\guardknockclient\test").Run()
 
 fmt.Println("Ответ от сервера:", string(body))
}
