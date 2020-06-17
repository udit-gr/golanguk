package postgres

import (
	"context"
	"database/sql"
	"log"
)

func InsertEmployeeDetails(dbDriver *sql.DB, employeeDetails EmployeeDetails) (err error) {

	ctx := context.Background()
	tx, err := dbDriver.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("[ERR][InsertEmployeeDetails] Error while starting the transaction, Err : %+v", err)
		return
	}

	// inserting employee info
	_, err = tx.ExecContext(ctx, queryInsertEmployeeInfo, employeeDetails.Name, employeeDetails.Age, employeeDetails.Gender)
	if err != nil {
		log.Printf("[ERR][InsertEmployeeDetails] Error while inserting employee info, Err : %+v", err)
		tx.Rollback()
		return
	}

	// getting total number of employees
	row := tx.QueryRow(queryGetTotalEmployees)
	var lastEmployeeID int
	err = row.Scan(&lastEmployeeID)
	if err != nil {
		log.Printf("[ERR][InsertEmployeeDetails] Error while getting total number of employees, Err : %+v", err)
		tx.Rollback()
		return
	}

	// insert employee department info
	_, err = tx.ExecContext(ctx, queryInsertEmployeeDepartmentInfo, lastEmployeeID, employeeDetails.Department, employeeDetails.Position, employeeDetails.Manager)
	if err != nil {
		log.Printf("[ERR][InsertEmployeeDetails] Error while inserting employee department info, Err : %+v", err)
		tx.Rollback()
		return
	}

	// commiting the change
	err = tx.Commit()
	if err != nil {
		log.Printf("[ERR][InsertEmployeeDetails] Error while committing the transaction, Err : %+v", err)
		return
	}

	return nil
}
