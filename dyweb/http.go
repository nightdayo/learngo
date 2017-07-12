package main
import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)
type Infromation struct{  
    Name string  
}

type Postinfo struct{  
    Postname string
    Postage string  
}

func main() {  
    http.HandleFunc("/info", infoHandler) 
    http.HandleFunc("/posttest", postHandler) 
    
    err := http.ListenAndServe(":9090", nil)  
    if err != nil {  
        log.Fatal("ListenAndServe: ", err.Error())  
    }  
}  
func infoHandler(w http.ResponseWriter, r *http.Request) {  
      info := new(Infromation)
    if r.Method == "GET" {  
        info.Name = "AV"  
        t, err := template.ParseFiles("switch.html")  
        if err != nil {  
            http.Error(w, err.Error(),http.StatusInternalServerError)  
            return  
        }  
            t.Execute(w, info)  
            return  
    }   
    if r.Method == "POST" {  
      fmt.Println("click")  
        info.Name = "B"  
        t, err := template.ParseFiles("switch.html")  
        if err != nil {  
            http.Error(w, err.Error(),http.StatusInternalServerError)  
            return  
        }  
            t.Execute(w, info)  
        return  
    }  
}  

func postHandler(w http.ResponseWriter, r *http.Request) {  
      info := new(Postinfo)
    if r.Method == "GET" {  
    	fmt.Println("build posttest") 
        info.Postname = "Your Name"  
        info.Postage  = "Your  age"
        t, err := template.ParseFiles("posttest.html")  
        if err != nil {  
            http.Error(w, err.Error(),http.StatusInternalServerError)  
            fmt.Println(err)
            return  
        }             
         t.Execute(w, info)  

            return  
    }   
    if r.Method == "POST" {  
      fmt.Println("click")  
        info.Postname = "Mike"  
        info.Postage  = "19"
        t, err := template.ParseFiles("posttest.html")  
        if err != nil {  
            http.Error(w, err.Error(),http.StatusInternalServerError)
            fmt.Println(err)
            return  
        }  
            t.Execute(w, info)  
        return  
    }  
}  