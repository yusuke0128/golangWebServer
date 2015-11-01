package main
import ( "fmt"
         "net/http"
         "log"
         "html"
)

func hellowWorldServer(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w, "Hello, World\n")
}

func Display(w http.ResponseWriter, r *http.Request){
 hoge := r.FormValue("hoge")
 fmt.Fprintf(w, "%s",html.EscapeString(hoge))
}

func main(){
http.HandleFunc("/", hellowWorldServer)
http.HandleFunc("/disp", Display)
log.Printf("Start Go HTTP Server")
 err := http.ListenAndServe(":4000", nil)
 if err != nil {
       log.Fatal("ListenAndServe: ", err)
    }
} 
