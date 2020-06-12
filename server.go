package main

import (
  "io"
  "html/template"
	"net/http"
  "crypto/subtle"

  //_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

)

//Setting up echo templates
type Template struct {
  templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

  if viewContext, isMap := data.(map[string]interface{}); isMap {
    viewContext["reverse"] = c.Echo().Reverse
  }

  return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
// Start echo Server
	e := echo.New()

// Load and Render templates
  t := &Template{
    templates: template.Must(template.ParseGlob("html/*.html")),
  }

  e.Renderer = t

// Middleware Setup
  e.Use(middleware.Recover())
  e.Use(middleware.Logger())

// HTTP to HTTPS Redirect
  e.Pre(middleware.HTTPSRedirect())

// Load statics at root
  e.Static("/", "html")

// Routing
  e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "")
	})

  r := e.Group("/demo.html")
  r.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
  	if subtle.ConstantTimeCompare([]byte(username), []byte(username)) == 1 &&
  		subtle.ConstantTimeCompare([]byte(password), []byte(getUserPass(username))) == 1 {
  		return true, nil
  	}
  	return false, nil
  }))
  r.GET("", func(c echo.Context) error {
		return c.Render(http.StatusOK, "demo.html", "")
	})

// Form routes
  e.POST("/addDept", addDeptForm)
  e.POST("/addDeptMan", addDeptManForm)
  e.POST("/termDeptMan", termDeptManForm)
  e.POST("/addEmp", addEmpForm)
  e.POST("/termEmp", termEmpForm)
  e.POST("/addEmpSup", addEmpSupForm)
  e.POST("/termEmpSup", termEmpSupForm)
  e.POST("/addPro", addProForm)
  e.POST("/termPro", termProForm)
  e.POST("/addEmpPro", addEmpProForm)
  e.POST("/termEmpPro", termEmpProForm)
  e.POST("/addEmpProWrk", addEmpProWrkForm)
  e.POST("/termEmpProWrk", termEmpProWrkForm)


// Tables routes
  e.GET("/dept.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataRows(c, "dept"))
  }).Name = "dept"

  e.GET("/emp.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataRows(c, "emp"))
  }).Name = "emp"

  e.GET("/pro.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataRows(c, "pro"))
  }).Name = "pro"

  e.GET("/emp-pro.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataRows(c, "emppro"))
  }).Name = "emppro"

  e.GET("/emp-pro-wrk.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataRows(c, "empprowrk"))
  }).Name = "empprowrk"

  e.GET("/emp-sup.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataRows(c, "empsup"))
  }).Name = "empsup"

  e.GET("/dept-man.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataRows(c, "deptman"))
  }).Name = "deptman"

  //Option set routes
  e.GET("/optDept-addDeptMan.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "dept", "Department:", "optDept-addDeptMan"))
  }).Name = "optDept-addDeptMan"

  e.GET("/optEmp-addDeptMan.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "emp", "Manager:", "optEmp-addDeptMan"))
  }).Name = "optEmp-addDeptMan"

  e.GET("/optDeptMan-termDeptMan.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "deptman", "Department Manager:", "optDeptMan-termDeptMan"))
  }).Name = "optDeptMan-termDeptMann"

  e.GET("/optDept-addEmp.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "dept", "Department:", "optDept-addEmp"))
  }).Name = "optDept-addEmp"

  e.GET("/optEmp-termEmp.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "emp", "Employee:", "optEmp-termEmp"))
  }).Name = "optEmp-termEmp"

  e.GET("/optEmp-addEmpSup1.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "emp", "Employee:", "optEmp-addEmpSup1"))
  }).Name = "optEmp-addEmpSup1"

  e.GET("/optEmp-addEmpSup2.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "emp", "Supervisor:", "optEmp-addEmpSup2"))
  }).Name = "optEmp-addEmpSup2"

  e.GET("/optEmpSup-termEmpSup.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "empsup", "Supervision:", "optEmpSup-termEmpSup"))
  }).Name = "optEmpSup-termEmpSup"

  e.GET("/optDept-addPro.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "dept", "Deparment:", "optDept-addPro"))
  }).Name = "optDept-addPro"

  e.GET("/optEmp-addPro.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "emp", "Requester:", "optEmp-addPro"))
  }).Name = "optEmp-addPro"

  e.GET("/optPro-termPro.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "pro", "Projects:", "optPro-termPro"))
  }).Name = "optPro-termPro"

  e.GET("/optEmp-addEmpPro.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "emp", "Employee:", "optEmp-addEmpPro"))
  }).Name = "optEmp-addEmpPro"

  e.GET("/optPro-addEmpPro.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "pro", "Project:", "optPro-addEmpPro"))
  }).Name = "optPro-addEmpPro"

  e.GET("/optEmpPro-termEmpPro.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "emppro", "Employee Project:", "optEmpPro-termEmpPro"))
  }).Name = "optEmpPro-termEmpPro"

  e.GET("/optEmpPro-addEmpProWrk.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "emppro", "Employee Project:", "optEmpPro-addEmpProWrk"))
  }).Name = "optEmpPro-addEmpProWrk"

  e.GET("/optEmpProWrk-termEmpProWrk.html", func(c echo.Context) error {
      return c.Render(http.StatusOK, "empty.html", getDataOpts(c, "empprowrk", "Employee Project Work:", "optEmpProWrk-termEmpProWrk"))
  }).Name = "optEmpProWrk-termEmpProWrk"



// Start servers
  go func(c *echo.Echo){
    e.Logger.Fatal(e.Start(":80"))
  }(e)
  e.Logger.Fatal(e.StartTLS(":443", "cert.pem", "key.pem"))
}
