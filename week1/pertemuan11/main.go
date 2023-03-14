package main 

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
    _"github.com/go-sql-driver/mysql"
	"log"
)

type Contact struct {
    Id int
    Name string
    Phone string
}


func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_gocrud")
    if err != nil {
        return nil, err
    }
    return db, nil
}

// type M map[string]interface{}
var tmpl = template.Must(template.ParseGlob("views/*"))
// var tmpl = template.Must(template.ParseFiles("views/index.html"))

// func sqlQueryRow() {
//     var db, err = connect()
//     if err != nil {
//         fmt.Println(err.Error())
//         return
//     }
//     defer db.Close()

//     var result = Contact{}
//     var id = 1
//     err = db.
//         QueryRow("select name, phone from tb_contact where id = ?", id).
//         Scan(&result.name, &result.phone)
//     if err != nil {
//         fmt.Println(err.Error())
//         return
//     }

//     fmt.Printf("name: %s\nphone: %d\n", result.name, result.phone)
// }

func Index(w http.ResponseWriter, r *http.Request) {
    db, err := connect()
    dbRows, err := db.Query("SELECT * FROM tb_contact")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
	// defer db.Close()

	contactResult := []Contact{}

    for dbRows.Next() {
        each := Contact{}
        err = dbRows.Scan(&each.Id, &each.Name, &each.Phone)
        if err != nil {
            panic(err.Error())
        }
        contactResult = append(contactResult, each)
    }
    // var tmpl = template.Must(template.ParseFiles("views/index.html"))
    tmpl.ExecuteTemplate(w, "Index", contactResult)
    // tmpl.Execute(w, "index", contactResult)
    defer db.Close()
}

func createContact(w http.ResponseWriter, r *http.Request) {
	// db, err := connect()
    // dbRows, err := db.Query("SELECT * FROM tb_contact")
    // if err != nil {
    //     fmt.Println(err.Error())
    //     return
    // }
	tmpl.ExecuteTemplate(w, "createContact", nil)
}


func contactEdit(w http.ResponseWriter, r *http.Request) {
	// db, err := connect()
    // dbRows, err := db.Query("SELECT * FROM tb_contact")
    // if err != nil {
    //     fmt.Println(err.Error())
    //     return
    // }
	tmpl.ExecuteTemplate(w, "contactEdit", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
	if r.Method == "POST" {
		name := r.FormValue("name")
		phone := r.FormValue("phone")
		insForm, err := db.Prepare("INSERT INTO tb_contact (name, phone) VALUES (?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, phone)
		log.Println("Insert Data: name " + name + " | phone " + phone)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db, err := connect()
	if err != nil {
        fmt.Println(err.Error())
        return
    }

    nId := r.URL.Query().Get("id")
    dbRows, err := db.Query("SELECT * FROM tb_contact WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }

	each := Contact{}

	for dbRows.Next() {
        err = dbRows.Scan(&each.Id, &each.Name, &each.Phone)
        if err != nil {
            panic(err.Error())
        }
	}
	tmpl.ExecuteTemplate(w, "contactEdit", each)
    defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
    db, err := connect()
	if err != nil {
        fmt.Println(err.Error())
        return
	}
	
    if r.Method == "POST" {
        id := r.FormValue("uid")
        name := r.FormValue("name")
        phone := r.FormValue("phone")
        insForm, err := db.Prepare("UPDATE tb_contact SET name=?, phone=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, phone, id)
        log.Println("UPDATE: Name: " + name + " | Phone: " + phone)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
        fmt.Println(err.Error())
        return
	}

    contactFound := r.URL.Query().Get("id")
    deleteForm, err := db.Prepare("DELETE FROM tb_contact WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    deleteForm.Exec(contactFound)
    log.Println("DELETED")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {
    // db, err := connect()
    // dbRows, err := db.Query("SELECT * FROM tb_contact")
    // if err != nil {
    //     fmt.Println(err.Error())
    //     return
    // }
    // contactResult := []Contact{}
    // for dbRows.Next() {
    //     each := Contact{}
    //     err = dbRows.Scan(&each.Id, &each.Name, &each.Phone)
    //     if err != nil {
    //         panic(err.Error())
    //     }
    //     contactResult = append(contactResult, each)
    // }

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// var data = M{"name": "Bat"}
	// 	var tmpl = template.Must(template.ParseFiles(
	// 		"views/index.html",
	// 	))

	// 	var err = tmpl.ExecuteTemplate(w, "Index", contactResult)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

    http.HandleFunc("/", Index)
	http.HandleFunc("/contact/create", createContact)
	http.HandleFunc("/contact/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
