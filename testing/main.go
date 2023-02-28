package main

import (
	"encoding/csv"
	"fmt"

	// "gobaidumap"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
	func IP2Addr(ipAddress string) (result string) {
		IPToAddress, err := gobaidumap.GetAddressViaIP(ipAddress)
		if err != nil {
			fmt.Println(err)
		} else {
			//println(From ip to address - address, IPtoAddress.Content.Address)
			//		fmt.Println("从ip到地址-地址", IPToAddress.Content.Address)
			result = IPToAddress.Content.Address
		}
		return
	}
*/
func WriteFile(content []string, name string, FileServer string) {
	f, err := os.OpenFile("./"+FileServer+"/"+name+".xls", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal("WriteFile: ", err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)

	writer.Write(content)

	writer.Flush()
}

func timer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/timer.html")
		fmt.Println("timer GET")
		log.Println(t.Execute(w, nil))
	}
}

func navigator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/navigator.html")
		fmt.Println("navigator GET")
		log.Println(t.Execute(w, nil))
	}
}

/*
func iframe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/iframe.html")
		fmt.Println("iframe GET")
		log.Println(t.Execute(w, nil))
	}
}

func picassauth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/picassauth.html")
		fmt.Println("picassauth GET")
		log.Println(t.Execute(w, nil))
	}
}
*/

func stanford(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/stanford.html")
		fmt.Println("stanford GET")
		log.Println(t.Execute(w, nil))
	}
}

func oregonState(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/oregonState.html")
		fmt.Println("oregonState GET")
		log.Println(t.Execute(w, nil))
	}
}

func auburn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/auburn.html")
		fmt.Println("auburn GET")
		log.Println(t.Execute(w, nil))
	}
}

func alaska(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/alaska.html")
		fmt.Println("alaska GET")
		log.Println(t.Execute(w, nil))
	}
}

func texas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/texas.html")
		fmt.Println("texas GET")
		log.Println(t.Execute(w, nil))
	}
}

func pennState(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/pennState.html")
		fmt.Println("pennState GET")
		log.Println(t.Execute(w, nil))
	}
}

func northDakota(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/northDakota.html")
		fmt.Println("northDakota GET")
		log.Println(t.Execute(w, nil))
	}
}

func colorado(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/colorado.html")
		fmt.Println("colorado GET")
		log.Println(t.Execute(w, nil))
	}
}

func dartmouth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/darmouth.html")
		fmt.Println("dartmouth GET")
		log.Println(t.Execute(w, nil))
	}
}

func wisconsin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/wisconsin.html")
		fmt.Println("wisconsin GET")
		log.Println(t.Execute(w, nil))
	}
}

func florida(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/florida.html")
		fmt.Println("florida GET")
		log.Println(t.Execute(w, nil))
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// client attempting to access index.html...
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/index.html")
		fmt.Println("index GET")
		log.Println(t.Execute(w, nil))
	}
	// client attempting to send information...
	// this is where data is collected?
	if r.Method == "POST" {
		fmt.Println("index POST")
		//Get user IP
		loginip := strings.Split(r.RemoteAddr, ":")[0]
		fmt.Println("ip:", loginip)
		// IP2Addr(loginip)		supposed to return a string, but return value is not used so whats the point?
		//processing form
		Addr := r.ParseForm()
		fmt.Println("Addr:", Addr)
		fmt.Println("name:", r.Form["name"])

		// servers
		fmt.Println("stanford:", r.Form["stanford"])
		fmt.Println("oregonState:", r.Form["oregonState"])
		fmt.Println("auburn:", r.Form["auburn"])
		fmt.Println("alaska:", r.Form["alaska"])
		fmt.Println("texas:", r.Form["texas"])
		fmt.Println("pennState:", r.Form["pennState"])
		fmt.Println("northDakota:", r.Form["northDakota"])
		fmt.Println("colorado:", r.Form["colorado"])
		fmt.Println("dartmouth:", r.Form["dartmouth"])
		fmt.Println("wisconsin:", r.Form["wisconsin"])
		fmt.Println("florida:", r.Form["florida"])

		// navigator
		fmt.Println("deviceType:", r.Form["deviceType"])
		fmt.Println("OSname:", r.Form["OSname"])
		fmt.Println("browserName:", r.Form["browserName"])
		fmt.Println("browserVer:", r.Form["browserVer"])
		fmt.Println("adaptType:", r.Form["adaptType"])

		// write the data
		// WriteFile(content []string,name string,FileServer string)
		data := []string{"", r.Form["baidu"][0], r.Form["sina"][0], r.Form["nju"][0], r.Form["iqiyi"][0], r.Form["douban"][0], r.Form["so"][0], r.Form["youku"][0], r.Form["qidian"][0], r.Form["2345"][0], r.Form["dianping"][0], r.Form["seu"][0]}
		WriteFile(data, r.Form["name"][0], "hash")
		data2 := []string{r.Form["name"][0], loginip, r.Form["deviceType"][0], r.Form["OSname"][0], r.Form["browserName"][0], r.Form["browserVer"][0], r.Form["adaptType"][0]}
		WriteFile(data2, r.Form["name"][0], "ip")
	}
}

func main() {
	fmt.Println("connecting 127.0.0.1:9000 ......")
	fs := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))

	// when these paths are requested, call this function to do stuff...
	http.HandleFunc("/", index)
	http.HandleFunc("/timer.html", timer)
	http.HandleFunc("/navigator.html", navigator)
	// http.HandleFunc("/picassauth.html", picassauth)
	// http.HandleFunc("/iframe.html", iframe)
	http.HandleFunc("/stanford.html", stanford)
	http.HandleFunc("/oregonState.html", oregonState)
	http.HandleFunc("/auburn.html", auburn)
	http.HandleFunc("/alaska.html", alaska)
	http.HandleFunc("/texas.html", texas)
	http.HandleFunc("/pennState.html", pennState)
	http.HandleFunc("/northDakota.html", northDakota)
	http.HandleFunc("/colorado.html", colorado)
	http.HandleFunc("/dartmouth.html", dartmouth)
	http.HandleFunc("/wisconsin.html", wisconsin)
	http.HandleFunc("/florida.html", florida)

	// start http server with given address and a handler. handler is usually nil, which
	// means to use DefaultServeMux
	// Handle and HandleFunc add handlers to DefaultServeMux
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}