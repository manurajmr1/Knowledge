package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/", SendaMail)          // set router
	err := http.ListenAndServe(":9002", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func SendaMail(w http.ResponseWriter, re *http.Request) {
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}

	ch := make(chan int, 2)

	for i := 1; i < 15; i++ {
		r := NewRequest("testaccnt"+strconv.Itoa(i)+"@mailinator.com", "Hello Junk!", "Hello, World!")
		err := r.ParseTemplate("templates/template.html", templateData)
		if err == nil {
			go r.SendaEmail(ch)
		} else {
			fmt.Println(err)
		}
	}
	for i := 1; i < 15; i++ {
		fmt.Println(<-ch)
	}
	//
	// for msg := range ch {
	// 			 fmt.Println(msg)
	// 	}
	// fmt.Println("Hi%d",<-ch)
	fmt.Fprintf(w, "mail send") // send data to client side
}

func (r *Request) SendaEmail(ch chan int) (bool, error) {

	m := gomail.NewMessage()
	m.SetHeader("From", "manuraj.mr@fingent.com")
	m.SetHeader("To", r.to)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", r.subject)
	m.SetBody("text/html", r.body)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 465, "username", "pass")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	ch <- 1
	// fmt.Fprintf(w, "mail send") // send data to client side

	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()

	return nil
}

//Request struct
type Request struct {
	from    string
	to      string
	subject string
	body    string
}

func NewRequest(to string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}
