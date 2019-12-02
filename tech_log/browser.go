package main


import(
	"fmt"
	"log"
	"net/http"
	"os/exec"
)


func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/start", startBrowser)
	http.HandleFunc("/stop", stopBrowser)
	http.HandleFunc("/cleanup", cleanup)
	http.HandleFunc("/geturl", geturl)
	http.ListenAndServe(":8080", nil)

}


func index(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w, "Hello Guest!")
}


func startBrowser(w http.ResponseWriter, req *http.Request){

	values, ok := req.URL.Query()["browser"]
	if !ok || len(values)<1{
		log.Println("browser Parameter is missing!")
		return
	}
	browser:= values[0]

	values, ok = req.URL.Query()["url"]
	if !ok || len(values)<1{
		log.Println("url Parameter is missing!")
		return
	}
	url:= values[0]
	if browser=="chrome"{
		cmd := exec.Command("mkdir", "/home/akash/Desktop/temp_chrome_testing")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)

		cmd = exec.Command("google-chrome", url, "--user-data-dir=/home/akash/Desktop/temp_chrome_testing")
		err = cmd.Run()
		log.Printf("Command finished with error: %v", err)
	}else if browser=="firefox" {

		cmd := exec.Command("mkdir", "/home/akash/Desktop/temp_firefox_dir")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)

		cmd =exec.Command("firefox", "--new-instance", "-P", "temp-user", url)
		err = cmd.Run()
		log.Printf("Command finished with error: %v", err)
	}
}

func stopBrowser(w http.ResponseWriter, req *http.Request){
	values, ok := req.URL.Query()["browser"]
	if !ok || len(values)<1{
		log.Println("browser Parameter is missing!")
		return
	}
	browser:= values[0]

	if browser=="firefox"{
		cmd:=exec.Command("killall", browser)
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)
	} else if browser== "chrome"{
		//cmd:=exec.Command("pkill", "--oldest", browser)
		//cmd:=exec.Command("killall", "google-chrome-stable")
		cmd:=exec.Command("killall", "chrome")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)	
	}
}

func cleanup(w http.ResponseWriter, req *http.Request){
	values, ok := req.URL.Query()["browser"]
	if !ok || len(values)<1{
		log.Println("browser Parameter is missing!")
		return
	}
	browser:= values[0]
	if browser=="chrome"{

		cmd:=exec.Command("rm", "-rf", "/home/akash/Desktop/temp_chrome_testing/")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)

	}else if browser=="firefox"{

		cmd:=exec.Command("rm", "-rf", "/home/akash/Desktop/temp_firefox_dir/")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)

	}


}

func geturl(w http.ResponseWriter, req *http.Request){
	values, ok := req.URL.Query()["browser"]
	if !ok || len(values)<1{
		log.Println("browser Parameter is missing!")
		return
	}
	browser:= values[0]
	fmt.Println(browser)
}



///start?browser=firefox&url=https://example.com
//stop?browser=firefox