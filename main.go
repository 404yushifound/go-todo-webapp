package main

import (
	"fmt"
	"net/http"
)

var DailyTasks = []string{
	"🌞 Morning Routine",
	"💻 Work on Projects",
	"🎬 Watch a Movie",
	"🌿 Enjoy a Walk and Dessert",
}

func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/add", addTaskHandler)

	fmt.Println("🌸 Server running at http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi Ayushi! 🌷 Welcome to your To-Do World ✨")
	fmt.Fprintln(w, "👉 Visit /tasks to see your list!")
	fmt.Fprintln(w, "➕ Visit /add to add a new task!")
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📝 Your Daily Tasks:")
	for i, task := range DailyTasks {
		fmt.Fprintf(w, "%d. %s\n", i+1, task)
	}
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Show the form
		fmt.Fprintln(w, `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Add Task</title>
		</head>
		<body>
			<h2>📝 Add a New Task</h2>
			<form method="POST" action="/add">
				<label>Enter a new task:</label><br>
				<input type="text" name="task" />
				<input type="submit" value="Add Task" />
			</form>
		</body>
	</html>`)
		fmt.Fprintln(w, "🌟 Add a new task to your list!")

	} else if r.Method == http.MethodPost {
		// Handle the form submission
		r.ParseForm()
		newTask := r.FormValue("task")
		if newTask != "" {
			DailyTasks = append(DailyTasks, newTask)
		}
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
	}
}
