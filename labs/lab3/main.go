package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("secret-key"))
)

var tmpl = template.Must(template.ParseFiles("templates/login.html", "templates/register.html", "templates/notes.html"))

type Note struct {
	Title   string
	Content string
	Created string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	if auth, _ := session.Values["authenticated"].(bool); auth {
		http.Redirect(w, r, "/notes", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if findUser(username, password) {
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
			}
			session.Save(r, w)
			http.Redirect(w, r, "/notes", http.StatusSeeOther)
			return
		} else {
			fmt.Fprintf(w, "Invalid credentials")
		}
		return
	}
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if writeUser(username, password) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	tmpl.ExecuteTemplate(w, "register.html", nil)
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	created := session.Values["username"].(string)
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		writeNote(Note{title, content, created})
		http.Redirect(w, r, "/notes", http.StatusSeeOther)
		return
	}
	notes := getNotes(created)
	tmpl.ExecuteTemplate(w, "notes.html", notes)
}

func writeNote(note Note) {
	file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s %s %s\n", note.Title, note.Content, note.Created))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func writeUser(username string, password string) bool {
	file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()
	loginUsed := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		if username == values[0] {
			loginUsed = true
			break
		}
	}
	if !loginUsed {
		_, err := file.WriteString(username + " " + password + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return false
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return !loginUsed
}

func findUser(username string, password string) bool {
	file, err := os.Open("users.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	status := false
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		if username == values[0] && password == values[1] {
			status = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return status
}

func getNotes(username string) []Note {
	file, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return make([]Note, 0)
	}
	defer file.Close()
	notes := make([]Note, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		if username == values[2] {
			notes = append(notes, Note{values[0], values[1], values[2]})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return notes
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/notes", notesHandler)
	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
