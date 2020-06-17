package postgres

var queryInsertNewEmployeeInfo = `
  INSERT INTO
  ws_employee (name, age gender)
  VALUES ($1, $2, $3)
`

var queryInsertEmployeeDepartmentInfo = `
  INSERT INTO
  ws_employee_department(id, department, position, manager)
  VALUES ($1, $2, $3, $4)
`
