package main
import ( "fmt"
         "net/http"
         "log"
         "html"
         "os"
         "syscall"
)
func daemon (nochdir, noclose int) int {
var ret uintptr
var err syscall.Errno

ret,_,err = syscall.Syscall(syscall.SYS_FORK,0,0,0)
if err != 0 { return -1	}
 switch (ret) {
  case 0: break
  default: os.Exit (0)
  }
 pid,_ := syscall.Setsid()
 if pid == -1 {
  return -1
 }
 if (nochdir == 0) { os.Chdir("/") }
 if noclose == 0 {
  f,e := os.Open("/dev/null")
   if e == nil {
     fd := int(f.Fd ())
     syscall.Dup2 (fd,int(os.Stdin.Fd()))
     syscall.Dup2 (fd,int(os.Stdout.Fd()))
     syscall.Dup2 (fd,int(os.Stderr.Fd()))
  }
}
	return 0
}
func hellowWorldServer(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w, "Hello, World\n")
}

func Display(w http.ResponseWriter, r *http.Request){
 hoge := r.FormValue("hoge")
 fmt.Fprintf(w, "%s",html.EscapeString(hoge))
}

func main(){
errcd := daemon(0,0)
 if errcd != 0 {
  fmt.Println("daemon err!!")
  os.Exit(1)
 }
os.Chdir("/home/yusuke/golangWebServer")
http.HandleFunc("/", hellowWorldServer)
http.HandleFunc("/disp", Display)
log.Printf("Start Go HTTP Server")
 err := http.ListenAndServe(":8080", nil)
 if err != nil {
  os.Exit(2)
 }
} 
