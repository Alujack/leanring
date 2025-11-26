package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Connect to database
	var err error
	dsn := "postgres://user:user123@localhost:5432/school_db?sslmode=disable"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot reach the database: ", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
	showMenu()
}

// ============= MENU ============
func showMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n====== School Management Console ======")
		fmt.Println("1. Register Student")
		fmt.Println("2. Register Teacher")
		fmt.Println("3. Create Course")
		fmt.Println("4. Enroll Student to Course")
		fmt.Println("5. Assign Teacher to Course")
		fmt.Println("6. Set Teacher Timeslot")
		fmt.Println("7. View Students")
		fmt.Println("8. View Teachers")
		fmt.Println("9. View Courses")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			registerStudent()
		case 2:
			registerTeacher()
		case 3:
			createCourse()
		case 4:
			enrollStudent()
		case 5:
			assignTeacher()
		case 6:
			setTeacherTimeslot()
		case 7:
			viewStudents()
		case 8:
			viewTeachers()
		case 9:
			viewCourses()
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

// ============= CRUD FUNCTIONS =============

// STUDENT REGISTER
func registerStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter student name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	_, err := db.Exec(`INSERT INTO students (name) VALUES ($1)`, name)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	fmt.Println("✅ Student registered:", name)
}

// TEACHER REGISTER
func registerTeacher() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter teacher name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	_, err := db.Exec(`INSERT INTO teachers (name) VALUES ($1)`, name)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	fmt.Println("✅ Teacher registered:", name)
}

// CREATE COURSE
func createCourse() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter course name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	_, err := db.Exec(`INSERT INTO courses (title) VALUES ($1)`, name)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	fmt.Println("✅ Course created:", name)
}

// ENROLL STUDENT TO COURSE
func enrollStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter student ID: ")
	sidStr, _ := reader.ReadString('\n')
	sid, _ := strconv.Atoi(strings.TrimSpace(sidStr))

	fmt.Print("Enter course ID: ")
	cidStr, _ := reader.ReadString('\n')
	cid, _ := strconv.Atoi(strings.TrimSpace(cidStr))

	_, err := db.Exec(`INSERT INTO enrollments (student_id, course_id) VALUES ($1, $2)`, sid, cid)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	fmt.Println("📚 Student enrolled successfully!")
}

// ASSIGN TEACHER TO COURSE
func assignTeacher() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter teacher ID: ")
	tidStr, _ := reader.ReadString('\n')
	tid, _ := strconv.Atoi(strings.TrimSpace(tidStr))

	fmt.Print("Enter course ID: ")
	cidStr, _ := reader.ReadString('\n')
	cid, _ := strconv.Atoi(strings.TrimSpace(cidStr))

	// Insert mapping into teacher_courses
	_, err := db.Exec(`INSERT INTO teacher_courses (teacher_id, course_id) VALUES ($1, $2)`, tid, cid)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("👨‍🏫 Teacher assigned to course!")
}


// SET TEACHER TIMESLOT
func setTeacherTimeslot() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter teacher ID: ")
	tidStr, _ := reader.ReadString('\n')
	tid, _ := strconv.Atoi(strings.TrimSpace(tidStr))

	fmt.Print("Enter timeslot (e.g. Mon 10AM-12PM): ")
	slot, _ := reader.ReadString('\n')
	slot = strings.TrimSpace(slot)

	_, err := db.Exec(`UPDATE teachers SET timeslot=$1 WHERE id=$2`, slot, tid)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("⏰ Timeslot updated!")
}

// VIEW STUDENTS
func viewStudents() {
	rows, err := db.Query(`SELECT id, name FROM students ORDER BY id`)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n📌 Students:")
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("%d: %s\n", id, name)
	}
}

// VIEW TEACHERS
func viewTeachers() {
	rows, err := db.Query(`SELECT id, name, timeslot FROM teachers ORDER BY id`)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n👨‍🏫 Teachers:")
	for rows.Next() {
		var id int
		var name, slot string
		rows.Scan(&id, &name, &slot)
		fmt.Printf("%d: %s | Timeslot: %s\n", id, name, slot)
	}
}

// VIEW COURSES
func viewCourses() {
	rows, err := db.Query(`
		SELECT 
			c.id, 
			c.title, 
			t.name
		FROM courses c
		LEFT JOIN teacher_courses tc ON tc.course_id = c.id
		LEFT JOIN teachers t ON t.id = tc.teacher_id
		ORDER BY c.id
	`)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n📘 Courses:")
	for rows.Next() {
		var courseID int
		var courseTitle, teacherName sql.NullString

		err := rows.Scan(&courseID, &courseTitle, &teacherName)
		if err != nil {
			fmt.Println("Scan error:", err)
			continue
		}

		title := ""
		if courseTitle.Valid {
			title = courseTitle.String
		}

		tName := "No teacher"
		if teacherName.Valid {
			tName = teacherName.String
		}

		fmt.Printf("%d: %s | Teacher: %s\n", courseID, title, tName)
	}
}

