  $("#addDept-btn").on("click", function(){
      var dept_name = $("#dept-add-name").val();

      $.ajax({
        type: "POST",
        url: "/addDept",
        data: JSON.stringify({
          dept_name: dept_name
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#addDeptMan-btn").on("click", function(){
      var did = $("#optDept-addDeptMan").val();
      var eid = $("#optEmp-addDeptMan").val();
      var start_date = $("#dept-man-add-date").val();

      $.ajax({
        type: "POST",
        url: "/addDeptMan",
        data: JSON.stringify({
          did: did,
          eid: eid,
          start_date: start_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#termDeptMan-btn").on("click", function(){
      var dmid = $("#optDeptMan-termDeptMan").val();
      var end_date = $("#dept-man-term-date").val();

      $.ajax({
        type: "POST",
        url: "/termDeptMan",
        data: JSON.stringify({
          dmid: dmid,
          end_date: end_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#addEmp-btn").on("click", function(){
      var employee_fname = $("#emp-add-fname").val();
      var employee_lname = $("#emp-add-lname").val();
      var did = $("#optDept-addEmp").val();
      var employee_title = $("#emp-add-title").val();
      var employee_salary = $("#emp-add-salary").val();
      var employee_dob = $("#emp-add-dob").val();
      var start_date = $("#emp-add-start-date").val();

      $.ajax({
        type: "POST",
        url: "/addEmp",
        data: JSON.stringify({
          employee_fname: employee_fname,
          employee_lname: employee_lname,
          did: did,
          employee_title: employee_title,
          employee_salary: employee_salary,
          employee_dob: employee_dob,
          start_date: start_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#termEmp-btn").on("click", function(){
      var eid = $("#optEmp-termEmp").val();
      var end_date = $("#emp-term-date").val();

      $.ajax({
        type: "POST",
        url: "/termEmp",
        data: JSON.stringify({
          eid: eid,
          end_date: end_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#addEmpSup-btn").on("click", function(){
      var eid = $("#optEmp-addEmpSup1").val();
      var sid = $("#optEmp-addEmpSup2").val();
      var start_date = $("#emp-sup-add-start-date").val();

      $.ajax({
        type: "POST",
        url: "/addEmpSup",
        data: JSON.stringify({
          eid: eid,
          sid: sid,
          start_date: start_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#termEmpSup-btn").on("click", function(){
      var esid = $("#optEmpSup-termEmpSup").val();
      var end_date = $("#emp-sup-term-date").val();

      $.ajax({
        type: "POST",
        url: "/termEmpSup",
        data: JSON.stringify({
          esid: esid,
          end_date: end_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#addPro-btn").on("click", function(){
      var did = $("#optDept-addPro").val();
      var eid = $("#optEmp-addPro").val();
      var project_desc = $("#pro-add-desc").val();
      var start_date = $("#pro-add-start-date").val();
      var due_date = $("#pro-add-due-date").val();

      $.ajax({
        type: "POST",
        url: "/addPro",
        data: JSON.stringify({
          did: did,
          eid: eid,
          project_desc: project_desc,
          start_date: start_date,
          due_date: due_date,
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#termPro-btn").on("click", function(){
      var pid = $("#optPro-termPro").val();
      var end_date = $("#pro-term-end-date").val();

      $.ajax({
        type: "POST",
        url: "/termPro",
        data: JSON.stringify({
          pid: pid,
          end_date: end_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#addEmpPro-btn").on("click", function(){
      var eid = $("#optEmp-addEmpPro").val();
      var pid = $("#optPro-addEmpPro").val();
      var project_role = $("#emp-pro-add-role").val();
      var est_hours = $("#emp-pro-add-hour").val();
      var start_date = $("#emp-pro-add-start-date").val();
      var due_date = $("#emp-pro-add-due-date").val();

      $.ajax({
        type: "POST",
        url: "/addEmpPro",
        data: JSON.stringify({
          eid: eid,
          pid: pid,
          project_role: project_role,
          est_hours: est_hours,
          start_date: start_date,
          due_date: due_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#termEmpPro-btn").on("click", function(){
      var epid = $("#optEmpPro-termEmpPro").val();
      var end_date = $("#emp-pro-term-end-date").val();

      $.ajax({
        type: "POST",
        url: "/termEmpPro",
        data: JSON.stringify({
          epid: epid,
          end_date: end_date
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#addEmpProWrk-btn").on("click", function(){
      var epid = $("#optEmpPro-addEmpProWrk").val();

      $.ajax({
        type: "POST",
        url: "/addEmpProWrk",
        data: JSON.stringify({
          epid: epid
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
  $("#termEmpProWrk-btn").on("click", function(){
      var epwid = $("#optEmpProWrk-termEmpProWrk").val();

      $.ajax({
        type: "POST",
        url: "/termEmpProWrk",
        data: JSON.stringify({
          epwid: epwid
        }),
        processData: false,
        contentType: "application/json",
        dataType: "json"
      }).done(
        function(data){
          window.location.reload();
        }
      ).fail(
        function(data){
          $(".form-result").text("POST request failed!");
        }
      )
  });
