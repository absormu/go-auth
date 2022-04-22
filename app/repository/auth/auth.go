package auth

import (
	"context"
	"strconv"

	"github.com/absormu/go-auth/app/entity"
	md "github.com/absormu/go-auth/app/middleware"
	cm "github.com/absormu/go-auth/pkg/configuration"
	db "github.com/absormu/go-auth/pkg/mariadb"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	sdk "gitlab.com/d3386/library"
)

func GetAuth(c echo.Context, params map[string]string) (user entity.User, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"params": params,
	}).Info("repository: Auth-GetAuth")

	db := db.MariaDBInit()

	defer db.Close()

	query := "SELECT id, name, email, user_contact_id, role_id FROM user"
	var condition string
	// Combine where clause
	clause := false
	for key, value := range params {

		if clause == false {
			condition += " WHERE "
		} else {
			condition += " AND "
		}
		condition += "is_deleted = 0"
		clause = true

		if clause == false {
			condition += " WHERE "
		} else {
			condition += " AND "
		}
		condition += "user" + "." + key + " = '" + value + "'"
		clause = true
	}

	query += condition

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: Auth-GetAuth-Query")

	result, err := db.Query(query)
	if err != nil {
		return
	}

	defer result.Close()

	for result.Next() {
		if e = result.Scan(&user.ID, &user.Name, &user.Email, &user.UserContactID, &user.RoleID); e != nil {
			return
		}
	}

	return
}

func Signup(c echo.Context, paramsSeller map[string]interface{}, paramsUser map[string]interface{}) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", paramsSeller).Info("repository: PostSeller")

	db := db.MariaDBInit()

	defer db.Close()

	ctx := context.Background()
	tx, e := db.BeginTx(ctx, nil)
	if e != nil {
		return
	}

	defer tx.Rollback()

	// insert seller
	querySeller := "INSERT INTO user_contact("
	var fieldsSeller = ""
	var valuesSeller = ""
	i := 0
	var valueSellerstr []interface{}
	for key, value := range paramsSeller {
		if (key != "email") && (key != "password") && (key != "company_id") && (key != "role_id") {
			fieldsSeller += "`" + key + "`"
			valuesSeller += "" + "?" + ""
			valueSellerstr = append(valueSellerstr, value)
			if (len(paramsSeller) - 1) > i {
				fieldsSeller += ", "
				valuesSeller += ", "
			}
			i++
		}
	}

	querySeller += fieldsSeller + ") VALUES(" + valuesSeller + ")"

	logger.WithFields(logrus.Fields{"query": querySeller}).Info("repository: CreateSeller-query")

	resultSeller, e := tx.ExecContext(ctx, querySeller, valueSellerstr...)
	if e != nil {
		return
	}
	sellerID, e := resultSeller.LastInsertId()
	if e != nil {
		return
	}

	// insert user
	queryUser := "INSERT INTO user("
	var fieldsUser = ""
	var valuesUser = ""
	j := 0
	var valueUserstr []interface{}
	for key, value := range paramsUser {
		if key != "telephone" && (key != "address") && (key != "seller_id") {
			fieldsUser += "`" + key + "`"
			valuesUser += "" + "?" + ""
			valueUserstr = append(valueUserstr, value)
			if (len(paramsUser) - 1) > j {
				fieldsUser += ", "
				valuesUser += ", "
			}
			j++
		}
	}

	sellerIdStr := strconv.FormatInt(sellerID, 10)
	queryUser += fieldsUser + ", user_contact_id) VALUES(" + valuesUser + "," + sellerIdStr + ")"

	logger.WithFields(logrus.Fields{"query": queryUser}).Info("repository: CreateUser-query")

	_, e = tx.ExecContext(ctx, queryUser, valueUserstr...)
	if e != nil {
		return
	}

	// Commit the transaction.
	if e = tx.Commit(); e != nil {
		return
	}

	return
}

func SignupNotification(c echo.Context, req entity.Email) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", req).Info("repository: SignupNotification")

	header := map[string]string{}
	header["Content-Type"] = "application/json; charset=utf-8"
	_, e = sdk.RawPostRequest(logger, cm.Config.ServiceEmail, cm.Config.Timeout, req, header)
	if e != nil {
		return
	}

	return
}

func GetAuthEmail(c echo.Context, params map[string]string) (user entity.User, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"params": params,
	}).Info("repository: Auth-GetEmail")

	db := db.MariaDBInit()

	defer db.Close()

	query := "SELECT id, name, email, user_contact_id, role_id, password FROM user"
	var condition string
	// Combine where clause
	clause := false
	for key, value := range params {

		if clause == false {
			condition += " WHERE "
		} else {
			condition += " AND "
		}
		condition += "user" + "." + key + " = '" + value + "'"
		clause = true
	}

	query += condition

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: Auth-GetEmail-Query")

	result, err := db.Query(query)
	if err != nil {
		return
	}

	defer result.Close()

	for result.Next() {
		if e = result.Scan(&user.ID, &user.Name, &user.Email, &user.UserContactID, &user.RoleID, &user.Password); e != nil {
			return
		}
	}

	return
}
