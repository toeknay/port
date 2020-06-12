package main

import (
    //"fmt"
    "net/http"
    //"time"
    //"encoding/json"
    "github.com/labstack/echo"
)



      /* Form handlers */


func addDeptForm(c echo.Context) (err error) {
  d := new(Department)
  if err = c.Bind(d); err != nil {
    return
  }
  if d.Dept_name == "" {
    return echo.ErrUnauthorized
  }
  err2 := addDept(d.Dept_name)
  if err2 == 0 {
		return echo.ErrUnauthorized
	}
  return echo.NewHTTPError(http.StatusOK, "Department Added")
}

func addDeptManForm(c echo.Context) (err error) {
  dm := new(Department_Managment)
  if err = c.Bind(dm); err != nil {
    return
  }
  if (dm.Did == "" || dm.Eid == "" || dm.Start_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := addDeptMan(dm.Did,dm.Eid,dm.Start_date)
  if err2 == 0 {
		return echo.ErrUnauthorized
	}
  return echo.NewHTTPError(http.StatusOK, "Department Manager Added")
}

func termDeptManForm(c echo.Context) (err error) {
  dm := new(Department_Managment)
  if err = c.Bind(dm); err != nil {
    return
  }
  if (dm.Dmid == "" || dm.End_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := termDeptMan(dm.Dmid,dm.End_date)
  if err2 == 0 {
		return echo.ErrUnauthorized
	}

  return echo.NewHTTPError(http.StatusOK, "Department Manager Terminated")
}

func addEmpForm(c echo.Context) (err error) {
  e := new(Employee)
  if err = c.Bind(e); err != nil {
    return
  }
  if (e.Did == "" || e.Employee_title == "" || e.Employee_fname == "" || e.Employee_lname == "" || e.Employee_salary == "" || e.Employee_dob == "" || e.Start_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := addEmp(e.Did,e.Employee_title,e.Employee_fname,e.Employee_lname,e.Employee_salary,e.Employee_dob,e.Start_date)
  if err2 == 0 {
		return echo.ErrUnauthorized
	}
  return echo.NewHTTPError(http.StatusOK, "Employee Added")
}

func termEmpForm(c echo.Context) (err error) {
  e := new(Employee)
  if err = c.Bind(e); err != nil {
    return
  }
  if (e.Eid == "" || e.End_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := termEmp(e.Eid,e.End_date)
  if err2 == 0 {
		return echo.ErrUnauthorized
	}

  return echo.NewHTTPError(http.StatusOK, "Employee Terminated")
}

func addEmpSupForm(c echo.Context) (err error) {
  es := new(Employee_Supervision)
  if err = c.Bind(es); err != nil {
    return
  }
  if (es.Eid == "" || es.Sid == "" || es.Start_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := addEmpSup(es.Eid,es.Sid,es.Start_date)
  if err2 == 0 {
    return echo.ErrUnauthorized
  }
  err3 := changeEmpSup(es.Eid)
  if err3 == 0 {
		return echo.ErrUnauthorized
	}
  return echo.NewHTTPError(http.StatusOK, "Employee Supervision Added")
}

func termEmpSupForm(c echo.Context) (err error) {
  es := new(Employee_Supervision)
  if err = c.Bind(es); err != nil {
    return
  }
  if (es.Esid == "" || es.End_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := termEmpSup(es.Esid,es.End_date)
  if err2 == 0 {
		return echo.ErrUnauthorized
	}
  return echo.NewHTTPError(http.StatusOK, "Employee Supervision Terminated")
}

func addProForm(c echo.Context) (err error) {
  p := new(Project)
  if err = c.Bind(p); err != nil {
    return
  }
  if (p.Did == "" || p.Eid == "" || p.Project_desc == "" || p.Start_date == "" || p.Due_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := addPro(p.Did,p.Eid,p.Project_desc,p.Start_date,p.Due_date)
  if err2 == 0 {
		return echo.ErrUnauthorized
	}
  return echo.NewHTTPError(http.StatusOK, "Project Added")
}

func termProForm(c echo.Context) (err error) {
  p := new(Project)
  if err = c.Bind(p); err != nil {
    return
  }
  if (p.Pid == "" || p.End_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := termPro(p.Pid,p.End_date)
  if err2 == 0 {
    return echo.ErrUnauthorized
  }
  return echo.NewHTTPError(http.StatusOK, "Project Terminated")
}

func addEmpProForm(c echo.Context) (err error) {
  ep := new(Employee_Project)
  if err = c.Bind(ep); err != nil {
    return
  }
  if (ep.Eid == "" || ep.Pid == "" || ep.Project_role == "" || ep.Est_hours == "" || ep.Start_date == "" || ep.Due_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := addEmpPro(ep.Eid,ep.Pid,ep.Project_role,ep.Est_hours,ep.Start_date,ep.Due_date)
  if err2 == 0 {
		return echo.ErrUnauthorized
	}
  return echo.NewHTTPError(http.StatusOK, "Employee Project Added")
}

func termEmpProForm(c echo.Context) (err error) {
  ep := new(Employee_Project)
  if err = c.Bind(ep); err != nil {
    return
  }
  if (ep.Epid == "" || ep.End_date == "") {
    return echo.ErrUnauthorized
  }
  err2 := termEmpPro(ep.Epid,ep.End_date)
  if err2 == 0 {
    return echo.ErrUnauthorized
  }
  return echo.NewHTTPError(http.StatusOK, "Employee Project Terminated")
}

func addEmpProWrkForm(c echo.Context) (err error) {
  epw := new(Employee_Project_Work)
  if err = c.Bind(epw); err != nil {
    return
  }
  if (epw.Epid == "") {
    return echo.ErrUnauthorized
  }
  err2 := addEmpProWrk(epw.Epid)
  if err2 == 0 {
    return echo.ErrUnauthorized
  }
  return echo.NewHTTPError(http.StatusOK, "Employee Work Added")
}

func termEmpProWrkForm(c echo.Context) (err error) {
  epw := new(Employee_Project_Work)
  if err = c.Bind(epw); err != nil {
    return
  }
  if (epw.Epwid == "") {
    return echo.ErrUnauthorized
  }
  err2 := termEmpProWrk(epw.Epwid)
  if err2 == 0 {
    return echo.ErrUnauthorized
  }
  return echo.NewHTTPError(http.StatusOK, "Employee Work Added")
}


      /* Make tables and option sets */


func getDataRows(c echo.Context, dataType string) error {
  var data []string
  var rows string
  var head string = "<table class=\"table table-striped\"> <thead><tr>"
  switch dataType {
    case "dept":
        result := getDepts()
        data = append(data, "Did", "Dept_name")
        for i := 0; i < len(result); i++ {
          rows += "<tr>"
          rows += "<td>" + result[i].Did + "</td>"
          rows += "<td>" + result[i].Dept_name + "</td>"
          rows += "</tr>"
        }
    case "emp":
        result := getEmps()
        data = append(data, "Eid", "Did", "Esid", "Employee_title", "Employee_fname", "Employee_lname", "Employee_salary", "Employee_dob", "Start_date" , "End_date")
        for i := 0; i < len(result); i++ {
          rows += "<tr>"
          rows += "<td>" + result[i].Eid + "</td>"
          rows += "<td>" + result[i].Did + "</td>"
          rows += "<td>" + result[i].Esid + "</td>"
          rows += "<td>" + result[i].Employee_title + "</td>"
          rows += "<td>" + result[i].Employee_fname + "</td>"
          rows += "<td>" + result[i].Employee_lname + "</td>"
          rows += "<td>" + result[i].Employee_salary + "</td>"
          rows += "<td>" + result[i].Employee_dob + "</td>"
          rows += "<td>" + result[i].Start_date + "</td>"
          rows += "<td>" + result[i].End_date + "</td>"
          rows += "</tr>"
        }
    case "pro":
        result := getPros()
        data = append(data, "Pid", "Did", "Eid", "Project_desc", "Start_date", "Due_date", "End_date")
        for i := 0; i < len(result); i++ {
          rows += "<tr>"
          rows += "<td>" + result[i].Pid + "</td>"
          rows += "<td>" + result[i].Did + "</td>"
          rows += "<td>" + result[i].Eid + "</td>"
          rows += "<td>" + result[i].Project_desc + "</td>"
          rows += "<td>" + result[i].Start_date + "</td>"
          rows += "<td>" + result[i].Due_date + "</td>"
          rows += "<td>" + result[i].End_date + "</td>"
          rows += "</tr>"
        }
    case "empsup":
        result := getEmpSups()
        data = append(data, "Esid", "Eid", "Start_date", "End_date", "Sid")
        for i := 0; i < len(result); i++ {
          rows += "<tr>"
          rows += "<td>" + result[i].Esid + "</td>"
          rows += "<td>" + result[i].Eid + "</td>"
          rows += "<td>" + result[i].Start_date + "</td>"
          rows += "<td>" + result[i].End_date + "</td>"
          rows += "<td>" + result[i].Sid + "</td>"
          rows += "</tr>"
        }
    case "emppro":
        result := getEmpPros()
        data = append(data, "Epid", "Eid", "Pid", "Project_role", "Est_hours", "Start_date", "Due_date", "End_date")
        for i := 0; i < len(result); i++ {
          rows += "<tr>"
          rows += "<td>" + result[i].Epid + "</td>"
          rows += "<td>" + result[i].Eid + "</td>"
          rows += "<td>" + result[i].Pid + "</td>"
          rows += "<td>" + result[i].Project_role + "</td>"
          rows += "<td>" + result[i].Est_hours + "</td>"
          rows += "<td>" + result[i].Start_date + "</td>"
          rows += "<td>" + result[i].Due_date + "</td>"
          rows += "<td>" + result[i].End_date + "</td>"
          rows += "</tr>"
        }
    case "empprowrk":
        result := getEmpProWrks()
        data = append(data, "Epwid", "Epid", "Start_time", "End_time")
        for i := 0; i < len(result); i++ {
          rows += "<tr>"
          rows += "<td>" + result[i].Epwid + "</td>"
          rows += "<td>" + result[i].Epid + "</td>"
          rows += "<td>" + result[i].Start_time + "</td>"
          rows += "<td>" + result[i].End_time + "</td>"
          rows += "</tr>"
        }
    case "deptman":
        result := getDeptMans()
        data = append(data, "Dmid", "Did", "Eid", "Start_date", "End_date")
        for i := 0; i < len(result); i++ {
          rows += "<tr>"
          rows += "<td>" + result[i].Dmid + "</td>"
          rows += "<td>" + result[i].Did + "</td>"
          rows += "<td>" + result[i].Eid + "</td>"
          rows += "<td>" + result[i].Start_date + "</td>"
          rows += "<td>" + result[i].End_date + "</td>"
          rows += "</tr>"
        }
  }
  for j := 0; j < len(data); j++ {
    hrow := "<th>" + data[j] + "</th>"
    head += hrow
  }
  head += "</tr></thead><tbody>"
  table := head + rows + "</tbody></table>"
  return c.HTML(http.StatusOK, table)
}

func getDataOpts(c echo.Context, dataType string, label string, formId string) error {
  var rows string
  var head string = "<div class=\"col-12 form-group\"><label>"
  switch dataType {
    case "dept":
        result := getDepts()
        for i := 0; i < len(result); i++ {
          rows += "<option value=\""
          //Option Value
          rows += result[i].Did
          rows += "\">"
          //Option Label
          rows += result[i].Dept_name
          rows += "</option>"
        }
    case "emp":
        result := getEmps()
        for i := 0; i < len(result); i++ {
          if result[i].End_date == "0000-00-00" {
            rows += "<option value=\""
            //Option Value
            rows += result[i].Eid
            rows += "\">"
            //Option Label
            rows += result[i].Employee_fname + " " + result[i].Employee_lname
            rows += "</option>"
          }
        }
    case "pro":
        result := getPros()
        for i := 0; i < len(result); i++ {
          if result[i].End_date == "0000-00-00" {
            rows += "<option value=\""
            //Option Value
            rows += result[i].Pid
            rows += "\">"
            //Option Label
            rows += result[i].Project_desc
            rows += "</option>"
          }
        }
    case "empsup":
        result := getEmpSups()
        for i := 0; i < len(result); i++ {
          if result[i].End_date == "0000-00-00" {
            emp := getEmp(result[i].Eid)
            empName := emp.Employee_fname + " " + emp.Employee_lname
            sup := getEmp(result[i].Sid)
            supName := sup.Employee_fname + " " + sup.Employee_lname
            rows += "<option value=\""
            //Option Value
            rows += result[i].Esid
            rows += "\">"
            //Option Label
            rows += empName + " / Supervisor: " + supName
            rows += "</option>"
          }
        }
    case "emppro":
        result := getEmpPros()
        for i := 0; i < len(result); i++ {
          if result[i].End_date == "0000-00-00" {
            emp := getEmp(result[i].Eid)
            empName := emp.Employee_fname + " " + emp.Employee_lname
            pro := getPro(result[i].Pid)
            proDesc := pro.Project_desc
            rows += "<option value=\""
            //Option Value
            rows += result[i].Epid
            rows += "\">"
            //Option Label
            rows += empName + " / " + proDesc + " / " + result[i].Project_role
            rows += "</option>"
          }
        }
    case "empprowrk":
        result := getEmpProWrks()
        for i := 0; i < len(result); i++ {
          if result[i].End_time == "0000-00-00 00:00:00" {

            empPro := getEmpPro(result[i].Epid)
            proRole := empPro.Project_role
            emp := getEmp(empPro.Eid)
            empName := emp.Employee_fname + " " + emp.Employee_lname
            rows += "<option value=\""
            //Option Value
            rows += result[i].Epwid
            rows += "\">"
            //Option Label
            rows += empName + " / " + proRole
            rows += "</option>"
          }
        }
    case "deptman":
        result := getDeptMans()
        for i := 0; i < len(result); i++ {
          if result[i].End_date == "0000-00-00" {
            emp := getEmp(result[i].Eid)
            empName := emp.Employee_fname + " " + emp.Employee_lname
            dept := getDept(result[i].Did)
            deptName := dept.Dept_name
            rows += "<option value=\""
            //Option Value
            rows += result[i].Dmid
            rows += "\">"
            //Option Label
            rows += deptName + " / " + empName
            rows += "</option>"
          }
        }
  }
  head += label
  head += "</label><select class=\"form-control valid\" name=\""
  head += formId
  head += "\" id=\""
  head += formId
  head += "\"><option value=\"\">-- Select One --</option>"
  table := head + rows + "</select></div>"
  return c.HTML(http.StatusOK, table)
}
