package main

import (
	"crypto/rand"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"authorname"`
	Website    string `json:"website"`
}

var courses []Course

func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

	// //fmt.Println(rand.Text())

}

// this is a standardd syntax to declare a controller, writer & reader should be positioned as it is.
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h2>This is the home page</h2>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get id value
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No value found with this courseId")
}

func addACourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty data is provided.")
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("Empty data is provided inside json.")
	}

	// rand/math is redundant and not used, also logic below does'nt provide you a string
	/* rand.Seed(time.Now().UnixMicro())
	vall := strconv.Itoa(rand.Intn(100))
	fmt.Println(vall) */

	course.CourseId = rand.Text()

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty data is provided.")
	}

	var course1 Course
	json.NewDecoder(r.Body).Decode(&course1)

	if course1.isEmpty() {
		json.NewEncoder(w).Encode("Empty data is provided inside json.")
	}

	for key, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:key], courses[key+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given courseId.")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	params := mux.Vars(r)

	for key, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:key], courses[key+1:]...)
			json.NewEncoder(w).Encode(`Course with courseId ${params["id"]} has been deleted successfully`)
			return
		}
	}

	json.NewEncoder(w).Encode("No matching courseId found.")
}
