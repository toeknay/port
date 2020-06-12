package main

import (
  "fmt"
)

//Setting up data objects
type User struct {
			Uid          string `json:"uid"`
			User_email   string `json:"user_email"`
			Username     string `json:"username"`
			User_pass    string `json:"user_pass"`
      Start_date   string `json:"start_date"`
			Can_login    string `json:"can_login"`
}

type Users struct {
			Users []User `json:"user"`
}

//Employee Managment data objects
type Department struct {
			Did              string    `json:"did"`
			Dept_name        string    `json:"dept_name"`
}

type Departments []Department

type Employee struct {
			Eid              string    `json:"eid"`
			Did              string    `json:"did"`
      Esid             string    `json:"esid"`
      Employee_title   string    `json:"employee_title"`
      Employee_fname   string    `json:"employee_fname"`
      Employee_lname   string    `json:"employee_lname"`
      Employee_salary  string    `json:"employee_salary"`
      Employee_dob     string    `json:"employee_dob"`
      Start_date       string    `json:"start_date"`
      End_date         string    `json:"end_date"`
}

type Employees []Employee

type Department_Managment struct {
      Dmid             string    `json:"dmid"`
      Did              string    `json:"did"`
      Eid              string    `json:"eid"`
      Start_date       string    `json:"start_date"`
      End_date         string    `json:"end_date"`
}

type Department_Managments []Department_Managment

type Employee_Supervision struct {
			Esid             string    `json:"esid"`
			Eid              string    `json:"eid"`
      Start_date       string    `json:"start_date"`
      End_date         string    `json:"end_date"`
      Sid              string    `json:"sid"`
}

type Employee_Supervisions []Employee_Supervision

type Project struct {
			Pid              string    `json:"pid"`
			Did              string    `json:"did"`
      Eid              string    `json:"eid"`
      Project_desc     string    `json:"project_desc"`
      Start_date       string    `json:"start_date"`
      Due_date         string    `json:"due_date"`
      End_date         string    `json:"end_date"`
}

type Projects []Project

type Employee_Project struct {
			Epid             string    `json:"epid"`
			Eid              string    `json:"eid"`
      Pid              string    `json:"pid"`
      Project_role     string    `json:"project_role"`
      Est_hours        string    `json:"est_hours"`
      Start_date       string    `json:"start_date"`
      Due_date         string    `json:"due_date"`
      End_date         string    `json:"end_date"`
}

type Employee_Projects []Employee_Project

type Employee_Project_Work struct {
			Epwid            string    `json:"epwid"`
			Epid             string    `json:"epid"`
      Start_time       string    `json:"start_time"`
      End_time         string    `json:"end_time"`
}

type Employee_Project_Works []Employee_Project_Work




          /* Employee Managment Demo */


      /* GET functions */

func getUsers() Users {
  con := CreateCon()
  sqlStatement := "SELECT uid, user_email, username, user_pass, start_date, can_login FROM Users ORDER BY uid"

  rows, err := con.Query(sqlStatement)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  result := Users{}
  for rows.Next() {
    user := User{}
    err2 := rows.Scan(&user.Uid, &user.User_email, &user.Username, &user.User_pass, &user.Start_date, &user.Can_login)
    if err2 != nil {
      fmt.Println(err2)
    }
    result.Users = append(result.Users, user)
  }
  return result
}

func getUser(username string) User {
  con := CreateCon()
  sqlStatement := "SELECT uid, user_email, username, user_pass, start_date, can_login FROM Users WHERE username = ?"
  var uid string
  var user_email string
  var user_pass string
  var start_date string
  var can_login string

  err := con.QueryRow(sqlStatement, username).Scan(&uid, &user_email, &username, &user_pass, &start_date, &can_login)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  result := User{Uid: uid, User_email: user_email, Username: username, User_pass: user_pass, Start_date: start_date, Can_login: can_login}
  return result
}

func getUserPass(username string) string {
  con := CreateCon()
  sqlStatement := "SELECT user_pass FROM Users WHERE username = ?"
  var user_pass string

  err := con.QueryRow(sqlStatement, username).Scan(&user_pass)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  result := user_pass
  return result
}

func addUser(user_email string, username string, user_pass string) int {
  con := CreateCon()
  sqlStatement := "INSERT INTO Users(user_email, username, user_pass, start_date, can_login) VALUES( ?, ?, ?, CURDATE(), 1)"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(user_email, username, user_pass)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

//Employee Managment Demo
func getDepts() Departments {
  con := CreateCon()
  sqlStatement := "SELECT did, dept_name FROM Departments ORDER BY did"

  rows, err := con.Query(sqlStatement)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  dresult := Departments{}
  for rows.Next() {
    dept := Department{}
    err2 := rows.Scan(&dept.Did, &dept.Dept_name)
    if err2 != nil {
      fmt.Println(err2)
    }
    dresult = append(dresult, dept)
  }
  return dresult
}

func getDept(did string) Department {
  con := CreateCon()
  sqlStatement := "SELECT did, dept_name FROM Departments WHERE did = ?"
  var dept_name string

  err := con.QueryRow(sqlStatement, did).Scan(&did, &dept_name )
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  result := Department{Did: did, Dept_name: dept_name }
  return result
}

func getEmps() Employees {
  con := CreateCon()
  sqlStatement := "SELECT eid, did, esid, employee_title, employee_fname, employee_lname, employee_salary, employee_dob, start_date, end_date FROM Employees ORDER BY eid"

  rows, err := con.Query(sqlStatement)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  eresult := Employees{}
  for rows.Next() {
    emp := Employee{}
    err2 := rows.Scan(&emp.Eid, &emp.Did, &emp.Esid, &emp.Employee_title, &emp.Employee_fname, &emp.Employee_lname, &emp.Employee_salary, &emp.Employee_dob, &emp.Start_date, &emp.End_date)
    if err2 != nil {
      fmt.Println(err2)
    }
    eresult = append(eresult, emp)
  }
  return eresult
}

func getEmp(eid string) Employee {
  con := CreateCon()
  sqlStatement := "SELECT eid, did, esid, employee_title, employee_fname, employee_lname, employee_salary, employee_dob, start_date, end_date FROM Employees WHERE eid = ?"
  var did string
  var esid string
  var employee_title string
  var employee_fname string
  var employee_lname string
  var employee_salary string
  var employee_dob string
  var start_date string
  var end_date string

  err := con.QueryRow(sqlStatement, eid).Scan(&eid, &did, &esid, &employee_title, &employee_fname, &employee_lname, &employee_salary, &employee_dob, &start_date, &end_date )
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  result := Employee{Eid: eid, Did: did, Esid: esid, Employee_title: employee_title, Employee_fname: employee_fname, Employee_lname: employee_lname, Employee_salary: employee_salary, Employee_dob: employee_dob, Start_date: start_date, End_date: end_date }
  return result
}



func getEmpSup(esid string) Employee_Supervision {
  con := CreateCon()
  sqlStatement := "SELECT esid, eid, start_date, end_date, sid FROM Employee_Supervision WHERE esid = ?"
  var eid string
  var start_date string
  var end_date string
  var sid string

  err := con.QueryRow(sqlStatement, esid).Scan(&esid, &eid, &start_date, &end_date, &sid)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  result := Employee_Supervision{Esid: esid, Eid: eid, Start_date: start_date, End_date: end_date, Sid: sid }
  return result
}


func getLastEmpSup() Employee_Supervision {
  con := CreateCon()
  sqlStatement := "SELECT esid, eid, start_date, end_date, sid FROM Employee_Supervision ORDER BY esid DESC"
  var eid string
  var esid string
  var start_date string
  var end_date string
  var sid string

  err := con.QueryRow(sqlStatement).Scan(&esid, &eid, &start_date, &end_date, &sid )
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  result := Employee_Supervision{Esid: esid, Eid: eid, Start_date: start_date, End_date: end_date, Sid: sid }
  return result
}


func getDeptMans() Department_Managments {
  con := CreateCon()
  sqlStatement := "SELECT dmid, did, eid, start_date, end_date FROM Department_managed ORDER BY dmid"

  rows, err := con.Query(sqlStatement)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  dmresult := Department_Managments{}
  for rows.Next() {
    demptm := Department_Managment{}
    err2 := rows.Scan(&demptm.Dmid, &demptm.Did, &demptm.Eid, &demptm.Start_date, &demptm.End_date)
    if err2 != nil {
      fmt.Println(err2)
    }
    dmresult = append(dmresult, demptm)
    //result.Departments = append(result.Departments, dept)
    //fmt.Println(user)
  }
  return dmresult
}

func getEmpSups() Employee_Supervisions {
  con := CreateCon()
  sqlStatement := "SELECT esid, eid, start_date, end_date, sid FROM Employee_Supervision ORDER BY esid"

  rows, err := con.Query(sqlStatement)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  esresult := Employee_Supervisions{}
  for rows.Next() {
    emps := Employee_Supervision{}
    err2 := rows.Scan(&emps.Esid, &emps.Eid, &emps.Start_date, &emps.End_date, &emps.Sid)
    if err2 != nil {
      fmt.Println(err2)
    }
    esresult = append(esresult, emps)
  }
  return esresult
}

func getPros() Projects {
  con := CreateCon()
  sqlStatement := "SELECT pid, did, eid, project_desc, start_date, due_date, end_date FROM Projects ORDER BY pid"

  rows, err := con.Query(sqlStatement)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  presult := Projects{}
  for rows.Next() {
    pro := Project{}
    err2 := rows.Scan(&pro.Pid, &pro.Did, &pro.Eid, &pro.Project_desc, &pro.Start_date, &pro.Due_date, &pro.End_date)
    if err2 != nil {
      fmt.Println(err2)
    }
    presult = append(presult, pro)
  }
  return presult
}

func getPro(pid string) Project {
  con := CreateCon()
  sqlStatement := "SELECT pid, did, eid, project_desc, start_date, due_date, end_date FROM Projects WHERE pid = ?"
  var did string
  var eid string
  var project_desc string
  var start_date string
  var due_date string
  var end_date string

  err := con.QueryRow(sqlStatement, pid).Scan(&pid, &did, &eid, &project_desc, &start_date, &due_date, &end_date )
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  result := Project{Pid: pid, Did: did, Eid: eid, Project_desc: project_desc, Start_date: start_date, Due_date: due_date, End_date: end_date }
  return result
}


func getEmpPros() Employee_Projects {
  con := CreateCon()
  sqlStatement := "SELECT epid, eid, pid, project_role, est_hours, start_date, due_date, end_date FROM Employee_Projects ORDER BY epid"

  rows, err := con.Query(sqlStatement)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  epresult := Employee_Projects{}
  for rows.Next() {
    empp := Employee_Project{}
    err2 := rows.Scan(&empp.Epid, &empp.Eid, &empp.Pid, &empp.Project_role, &empp.Est_hours, &empp.Start_date, &empp.Due_date, &empp.End_date)
    if err2 != nil {
      fmt.Println(err2)
    }
    epresult = append(epresult, empp)
  }
  return epresult
}

func getEmpPro(epid string) Employee_Project {
  con := CreateCon()
  sqlStatement := "SELECT epid, eid, pid, project_role, est_hours, start_date, due_date, end_date FROM Employee_Projects WHERE epid = ?"
  var eid string
  var pid string
  var project_role string
  var est_hours string
  var start_date string
  var due_date string
  var end_date string

  err := con.QueryRow(sqlStatement, epid).Scan(&epid, &eid, &pid, &project_role, &est_hours, &start_date, &due_date, &end_date )
  if err != nil {
     fmt.Println(err)
  }

  defer con.Close()

  result := Employee_Project{Epid: epid, Eid: eid, Pid: pid, Project_role: project_role, Est_hours: est_hours, Start_date: start_date, Due_date: due_date, End_date: end_date }
  return result
}

func getEmpProWrks() Employee_Project_Works {
  con := CreateCon()
  sqlStatement := "SELECT epwid, epid, start_time, end_time FROM Employee_Project_Works ORDER BY epwid"

  rows, err := con.Query(sqlStatement)
  if err != nil {
     fmt.Println(err)
  }
  defer con.Close()
  epwresult := Employee_Project_Works{}
  for rows.Next() {
    emppw := Employee_Project_Work{}
    err2 := rows.Scan(&emppw.Epwid, &emppw.Epid, &emppw.Start_time, &emppw.End_time)
    if err2 != nil {
      fmt.Println(err2)
    }
    epwresult = append(epwresult, emppw)
  }
  return epwresult
}

      /* ADD functions */

func addDept(dept_name string) int {
  con := CreateCon()
  sqlStatement := "INSERT INTO Departments(dept_name) VALUES( ? )"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(dept_name)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func addDeptMan(did string, eid string, start_date string) int {
  con := CreateCon()
  sqlStatement := "INSERT INTO Department_managed(did, eid, start_date, end_date) VALUES( ?, ?, ?, '0000-00-00' )"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
      //TODO or should it be fmt.Println(err) like the rest?
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(did,eid,start_date)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func termDeptMan(dmid string, end_date string) int {
  con := CreateCon()
  sqlStatement := "UPDATE Department_managed SET end_date = ? WHERE dmid = ?"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(end_date,dmid)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func addEmp(did string, employee_title string, employee_fname string, employee_lname string, employee_salary string, employee_dob string, start_date string) int {
  con := CreateCon()
  sqlStatement := "INSERT INTO Employees(did, esid, employee_title, employee_fname, employee_lname, employee_salary, employee_dob, start_date, end_date) VALUES( ?, 10, ?, ?, ?, ?, ?, ?, '0000-00-00' )"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
      //TODO or should it be fmt.Println(err) like the rest?
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(did,employee_title,employee_fname,employee_lname,employee_salary,employee_dob,start_date)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func termEmp(eid string, end_date string) int {
  con := CreateCon()
  sqlStatement := "UPDATE Employees SET end_date = ? WHERE eid = ?"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(end_date,eid)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func addEmpSup(empId string, supId string, startDate string) int {
  con := CreateCon()
  sqlStatement := "INSERT INTO Employee_Supervision(eid, start_date, end_date, sid) VALUES( ?, ?, '0000-00-00', ? )"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(empId,startDate,supId)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func changeEmpSup(empId string) int {
  con := CreateCon()
  empsup := getLastEmpSup()
  esid := empsup.Esid

  fmt.Println(esid)

  sqlStatement := "UPDATE Employees SET esid = ? WHERE eid = ?"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(esid,empId)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func termEmpSup(esid string, end_date string) int {
  con := CreateCon()
  sqlStatement := "UPDATE Employee_Supervision SET end_date = ? WHERE esid = ?"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
      //TODO or should it be fmt.Println(err) like the rest?
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(end_date,esid)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func addPro(dept string, req string, proDesc string, startDate string, dueDate string) int {
  con := CreateCon()
  sqlStatement := "INSERT INTO Projects(did, eid, project_desc, start_date, due_date, end_date) VALUES( ?, ?, ?, ?, ?, '0000-00-00' )"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(dept,req,proDesc,startDate,dueDate)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func termPro(pid string, end_date string) int {
  con := CreateCon()
  sqlStatement := "UPDATE Projects SET end_date = ? WHERE pid = ?"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
      //TODO or should it be fmt.Println(err) like the rest?
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(end_date,pid)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func addEmpPro(emp string, pro string, proRole string, estHour string, startDate string, dueDate string) int {
  con := CreateCon()
  sqlStatement := "INSERT INTO Employee_Projects(eid, pid, project_role, est_hours, start_date, due_date, end_date) VALUES( ?, ?, ?, ?, ?, ?, '0000-00-00' )"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(emp,pro,proRole,estHour,startDate,dueDate)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func termEmpPro(epid string, end_date string) int {
  con := CreateCon()
  sqlStatement := "UPDATE Employee_Projects SET end_date = ? WHERE epid = ?"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(end_date,epid)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func addEmpProWrk(empPro string) int {
  con := CreateCon()
  sqlStatement := "INSERT INTO Employee_Project_Works(epid, end_time) VALUES( ?, '0000-00-00 00:00:00' )"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(empPro)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}

func termEmpProWrk(epwid string) int {
  con := CreateCon()
  sqlStatement := "UPDATE Employee_Project_Works SET end_time = CURRENT_TIMESTAMP WHERE epwid = ?"
  stmt, err := con.Prepare(sqlStatement)

	if err != nil {
				fmt.Print(err.Error())
	}

	defer con.Close()
	result, err2 := stmt.Exec(epwid)

	// Exit if we get an error
	if err2 != nil {
		return 0
	}
  fmt.Println(result)
  return 1
}
