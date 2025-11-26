package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Student struct {
    ID       int
    Name     string
    Courses  []string
}

type Teacher struct {
    ID       int
    Name     string
    Courses  []string
    TimeSlot string
}

type Course struct {
    Name      string
    TeacherID int
}

var students []Student
var teachers []Teacher
var courses []Course

var studentIDCounter = 1
var teacherIDCounter = 1

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\n===== School Management Console =====")
        fmt.Println("1. Register Student")
        fmt.Println("2. Register Teacher")
        fmt.Println("3. Create Course")
        fmt.Println("4. Enroll Student to Course")
        fmt.Println("5. Assign Teacher to Course")
        fmt.Println("6. Teacher Set Time Slot")
        fmt.Println("7. View All Data")
        fmt.Println("0. Exit")

        fmt.Print("Choose: ")
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            registerStudent(reader)
        case "2":
            registerTeacher(reader)
        case "3":
            createCourse(reader)
        case "4":
            enrollStudent(reader)
        case "5":
            assignTeacher(reader)
        case "6":
            setTeacherTimeSlot(reader)
        case "7":
            viewAllData()
        case "0":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid option.")
        }
    }
}

// 1. Register Student
func registerStudent(reader *bufio.Reader) {
    fmt.Print("Enter student name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    student := Student{ID: studentIDCounter, Name: name}
    studentIDCounter++
    students = append(students, student)

    fmt.Println("Student registered successfully.")
}

// 2. Register Teacher
func registerTeacher(reader *bufio.Reader) {
    fmt.Print("Enter teacher name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    teacher := Teacher{ID: teacherIDCounter, Name: name}
    teacherIDCounter++
    teachers = append(teachers, teacher)

    fmt.Println("Teacher registered successfully.")
}

// 3. Create Course
func createCourse(reader *bufio.Reader) {
    fmt.Print("Enter course name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    course := Course{Name: name}
    courses = append(courses, course)

    fmt.Println("Course created.")
}

// 4. Enroll Student into Course
func enrollStudent(reader *bufio.Reader) {
    if len(students) == 0 || len(courses) == 0 {
        fmt.Println("No students or courses available.")
        return
    }

    fmt.Println("Students:")
    for _, s := range students {
        fmt.Printf("%d. %s\n", s.ID, s.Name)
    }

    fmt.Print("Enter student ID: ")
    var sid int
    fmt.Scan(&sid)

    var student *Student
    for i := range students {
        if students[i].ID == sid {
            student = &students[i]
        }
    }
    if student == nil {
        fmt.Println("Student not found.")
        return
    }

    fmt.Println("Courses:")
    for i, c := range courses {
        fmt.Printf("%d. %s\n", i+1, c.Name)
    }

    fmt.Print("Enter course number: ")
    var cindex int
    fmt.Scan(&cindex)
    cindex--

    if cindex < 0 || cindex >= len(courses) {
        fmt.Println("Invalid course.")
        return
    }

    student.Courses = append(student.Courses, courses[cindex].Name)
    fmt.Println("Student enrolled.")
}

// 5. Assign Teacher to Course
func assignTeacher(reader *bufio.Reader) {
    if len(teachers) == 0 || len(courses) == 0 {
        fmt.Println("No teachers or courses available.")
        return
    }

    fmt.Println("Teachers:")
    for _, t := range teachers {
        fmt.Printf("%d. %s\n", t.ID, t.Name)
    }

    fmt.Print("Enter teacher ID: ")
    var tid int
    fmt.Scan(&tid)

    var teacher *Teacher
    for i := range teachers {
        if teachers[i].ID == tid {
            teacher = &teachers[i]
        }
    }

    if teacher == nil {
        fmt.Println("Teacher not found.")
        return
    }

    fmt.Println("Courses:")
    for i, c := range courses {
        fmt.Printf("%d. %s\n", i+1, c.Name)
    }

    fmt.Print("Choose course number: ")
    var cindex int
    fmt.Scan(&cindex)
    cindex--

    if cindex < 0 || cindex >= len(courses) {
        fmt.Println("Invalid course.")
        return
    }

    courses[cindex].TeacherID = teacher.ID
    teacher.Courses = append(teacher.Courses, courses[cindex].Name)

    fmt.Println("Teacher assigned to course.")
}

// 6. Teacher sets time slot
func setTeacherTimeSlot(reader *bufio.Reader) {
    fmt.Println("Teachers:")
    for _, t := range teachers {
        fmt.Printf("%d. %s\n", t.ID, t.Name)
    }

    fmt.Print("Enter teacher ID: ")
    var tid int
    fmt.Scan(&tid)

    var teacher *Teacher
    for i := range teachers {
        if teachers[i].ID == tid {
            teacher = &teachers[i]
        }
    }

    if teacher == nil {
        fmt.Println("Teacher not found.")
        return
    }

    reader.ReadString('\n')
    fmt.Print("Enter time slot (e.g., Mon 2-4 PM): ")
    slot, _ := reader.ReadString('\n')
    teacher.TimeSlot = strings.TrimSpace(slot)

    fmt.Println("Time slot saved.")
}

// 7. View All Data
func viewAllData() {
    fmt.Println("\n--- Students ---")
    for _, s := range students {
        fmt.Printf("ID: %d, Name: %s, Courses: %v\n", s.ID, s.Name, s.Courses)
    }

    fmt.Println("\n--- Teachers ---")
    for _, t := range teachers {
        fmt.Printf("ID: %d, Name: %s, Courses: %v, TimeSlot: %s\n", t.ID, t.Name, t.Courses, t.TimeSlot)
    }

    fmt.Println("\n--- Courses ---")
    for _, c := range courses {
        fmt.Printf("Name: %s, TeacherID: %d\n", c.Name, c.TeacherID)
    }
}