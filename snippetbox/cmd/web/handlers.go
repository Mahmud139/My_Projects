package main

import (
	"errors"
	"fmt"
	//"html/template"
	"net/http"
	"strconv"
	// "strings"
	// "unicode/utf8"

	"mahmud139/snippetbox/pkg/forms"
	"mahmud139/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	//http.NotFound(w, r)
	// 	app.notFound(w)
	// 	return
	// }

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})

	/*
		data := &templateData{Snippets: s}
		// for _, snippet := range s {
		// 	fmt.Fprintf(w, "%v\n", snippet)
		// }

		//initialize a slice containing the paths to the two files. Note that the home.page.tmpl
		//must be the first file in the slice.
		files := []string{
			"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/home.page.tmpl",
			"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/base.layout.tmpl",
			"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/footer.partial.tmpl",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			//app.errorLog.Println(err.Error())
			// http.Error(w, "Internal Server Error", 500)
			//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			app.serverError(w, err)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			// app.errorLog.Println(err.Error())
			// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			app.serverError(w, err)
		}
		//w.Write([]byte("Hello from SnippetBox")) */
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.page.tmpl", nil)
}

func (app *application) userProfile(w http.ResponseWriter, r *http.Request) {
	userID := app.session.GetInt(r, "authenticatedUserID")

	user, err := app.users.Get(userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// fmt.Fprintf(w, "%+v", user)
	app.render(w, r, "profile.page.tmpl", &templateData{
		User: user,
	})
}

func (app *application) changePasswordForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "password.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) changePassword(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("currentPassword", "newPassword", "newPasswordConfirmation")
	form.MinLength("newPassword", 10)
	if form.Get("newPassword") != form.Get("newPasswordConfirmation") {
		form.Errors.Add("newPasswordConfirmation", "Passwords do not match")
	}

	if !form.Valid() {
		app.render(w, r, "password.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	userID := app.session.GetInt(r, "authenticatedUserID")
	err = app.users.ChangePassword(userID, form.Get("currentPassword"), form.Get("newPassword"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("currentPassword", "Current password is incorrect")
			app.render(w, r, "password.page.tmpl", &templateData{
				Form: form,
			})
		} else if err != nil {
			app.serverError(w, err)
		}
		return
	}
	app.session.Put(r, "flash", "Your password has been updated!")
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})

	/*
		data := &templateData{Snippet: s}

		files := []string {
			"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/show.page.tmpl",
			"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/base.layout.tmpl",
			"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/footer.partial.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
			return
		}

		err = ts.Execute(w, data)
		if err != nil {
			app.serverError(w, err)
		}

		//w.Write([]byte("Display a new snippet"))
		//fmt.Fprintf(w,"Display with specific snippet with ID %d...", id)
		//fmt.Fprintf(w, "%v", s) */
}

func (app *application) deleteSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	err = app.snippets.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Create a new snippet..."))
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	// w.Header().Set("Allow", "POST")
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	// w.WriteHeader(405)
	// 	// w.Write([]byte("Method not Allowed!"))
	// 	// http.Error(w, "Method not Allowed!", 405)
	// 	// http.Error(w,"Method not Allowed!", http.StatusMethodNotAllowed)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	// First we call r.ParseForm() which adds any data in POST request bodies
	// to the r.PostForm map. This also works in the same way for PUT and PATCH
	// requests. If there are any errors, we use our app.ClientError helper to send
	// a 400 Bad Request response to the user.
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	/*
		// Use the r.PostForm.Get() method to retrieve the relevant data fields
		// from the r.PostForm map.
		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		expires := r.PostForm.Get("expires")

		// Initialize a map to hold any validation errors.
		errors := make(map[string]string)
		// Check that the title field is not blank and is not more than 100 characters
		// long. If it fails either of those checks, add a message to the errors
		// map using the field name as the key.
		if strings.TrimSpace(title) == "" {
			errors["title"] = "This field cannot be blank"
		} else if utf8.RuneCountInString(title) > 100 {
			errors["title"] = "This field is too long (maximum is 100 characters"
		}

		if strings.TrimSpace(content) == "" {
			errors["content"] = "This field cannot be blank"
		}

		if strings.TrimSpace(expires) == "" {
			errors["expires"] = "This field cannot be blank"
		} else if expires != "365" && expires != "30" && expires != "7" {
			errors["expires"] = "This field is invalid"
		}

		if len(errors) > 0 {
			app.render(w, r, "create.page.tmpl", &templateData{
				FormErrors: errors,
				FormData: r.PostForm,
			})
			return
		} */

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "30", "7")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	// Because the form data (with type url.Values) has been anonymously embedded
	// in the form.Form struct, we can use the Get() method to retrieve
	// the validated value for a particular form field.
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Snippet successfully created")
	//w.Write([]byte("Create a new snippet"))
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
	/*The HTTP response status code 303 See Other is a way to redirect web applications
	to a new URI, particularly after a HTTP POST has been performed*/
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRx)
	form.MinLength("password", 10)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{
			Form: form,
		})
	}

	// Try to create a new user record in the database. If the email already exists
	// add an error message to the form and re-display it.
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Otherwise add a confirmation flash message to the session confirming that
	// their signup worked and asking them to log in.
	app.session.Put(r, "flash", "Your signup was successful, Please login.")

	// And redirect the user to the login page.
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

	//fmt.Fprintln(w, "Create a new user")
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
	//fmt.Fprintln(w, "Display the user login form")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Authenticate and login the user")
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Check whether the credentials are valid. If they're not, add a generic error
	// message to the form failures map and re-display the login page.
	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
	}

	app.session.Put(r, "authenticatedUserID", id)
	// Use the PopString method to retrieve and remove a value from the session 
	// data in one step. If no matching key exists this will return the empty string.
	path := app.session.PopString(r, "redirectPathAfterLogin")
	if path != "" {
		http.Redirect(w, r, path, http.StatusSeeOther)
		return
	}
	
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	app.session.Remove(r, "authenticatedUserID")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
