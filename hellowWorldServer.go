package main
import ( "fmt"
         "net/http"
         "log"
)

func hellowWorldServer(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w, "Hello, World\n")
}

func main(){
http.HandleFunc("/", hellowWorldServer)
log.Printf("Start Go HTTP Server")
 err := http.ListenAndServe(":4000", nil)
 if err != nil {
       log.Fatal("ListenAndServe: ", err)
    }
} 
