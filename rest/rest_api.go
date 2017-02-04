package rest

import (
	"net/http"
	"time"

	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/varunamachi/orekng/data"
	"github.com/varunamachi/orekng/olog"
	"gopkg.in/hlandau/passlib.v1"
)

const ky = "orek_2232redsfaj3234edsa"

func logIfError(err error) (errOut error) {
	if err != nil {
		olog.PrintError("REST", err)
	}
	return
}

func str(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func getAllUsers(ctx echo.Context) (err error) {
	users, err := data.GetStore().GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get User List",
				Message:   "Failed to fetch list of user",
				Error:     str(err)})

	} else {
		err = logIfError(ctx.JSON(http.StatusOK, users))
	}
	return err
}

func getUser(ctx echo.Context) (err error) {
	userName := ctx.Param("userName")
	email := ctx.Param("email")
	var user *data.User
	if len(userName) != 0 {
		user, err = data.GetStore().GetUser(userName)
	} else {
		user, err = data.GetStore().GetUserWithEmail(email)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get User Details",
				Message:   "Failed to fetch user details",
				Error:     str(err)})

	} else {
		err = logIfError(ctx.JSON(http.StatusOK, user))
	}
	return err
}

// func getUserWithEmail(ctx echo.Context) (err error) {
// 	email := ctx.Param("email")
// 	user, err := data.GetStore().GetUserWithEmail(email)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError,
// 			Result{
// 				Operation: "Get User Details",
// 				Message:   "Failed to fetch user details",
// 				Error: str(err)})
// 	} else {
// 		err = logIfError(ctx.JSON(http.StatusOK, user))
// 	}
// 	return err
// }

func createUser(ctx echo.Context) (err error) {
	var user data.User
	err = ctx.Bind(&user)
	if err == nil {
		err = data.GetStore().CreateUser(&user)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create User",
				Message:   "User creation failed",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create User",
				Message:   "User creation successful",
				Error:     ""})
	}
	return err
}

func updateUser(ctx echo.Context) (err error) {
	var user data.User
	err = ctx.Bind(&user)
	if err == nil {
		err = data.GetStore().UpdateUser(&user)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update User",
				Message:   "User update failed",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update User",
				Message:   "User update successful",
				Error:     ""})
	}
	return err
}

func deleteUser(ctx echo.Context) (err error) {
	userName := ctx.Param("userName")
	err = data.GetStore().DeleteUser(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Delete User",
				Message:   "User deletion failed",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete User",
				Message:   "User deletion successful",
				Error:     ""})
	}
	return err
}

func getAllEndpoints(ctx echo.Context) (err error) {
	endpoints, err := data.GetStore().GetAllEndpoints()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Endpoint",
				Message:   "Listing endpoint failed",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, endpoints))
	}
	return err
}

func getEndpoint(ctx echo.Context) (err error) {
	endpointID := ctx.Param("endpointID")
	endpoint, err := data.GetStore().GetEndpoint(endpointID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get Endpoint Details",
				Message:   "Failed to fetch endpoint details",
				Error:     str(err)})

	} else {
		err = logIfError(ctx.JSON(http.StatusOK, endpoint))
	}
	return err
}

func createEndpoint(ctx echo.Context) (err error) {
	var endpoint data.Endpoint
	err = ctx.Bind(&endpoint)
	if err == nil {
		err = data.GetStore().CreateEndpoint(&endpoint)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create Endpoint",
				Message:   "Failed to create endpoint",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create Endpoint",
				Message:   "Endpoint created succesffully",
				Error:     str(err)})
	}
	return err
}

func updateEndpoint(ctx echo.Context) (err error) {
	var endpoint data.Endpoint
	err = ctx.Bind(&endpoint)
	if err == nil {
		err = data.GetStore().UpdateEndpoint(&endpoint)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update Endpoint",
				Message:   "Failed to update endpoint",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update Endpoint",
				Message:   "Endpoint updated succesffully",
				Error:     str(err)})
	}
	return err
}

func deleteEndpoint(ctx echo.Context) (err error) {
	endpointID := ctx.Param("endpointID")
	err = data.GetStore().DeleteEndpoint(endpointID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Dalete Endpoint",
				Message:   "Failed to delete endpoint",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete Endpoint",
				Message:   "Endpoint delete succesffully",
				Error:     str(err)})
	}
	return err
}

func getAllVariables(ctx echo.Context) (err error) {
	variables, err := data.GetStore().GetAllVariables()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Variables",
				Message:   "Failed to list all variables",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, variables))
	}
	return err
}

func getVariablesForEndpoint(ctx echo.Context) (err error) {
	endpointID := ctx.Param("endpointID")
	variables, err := data.GetStore().GetVariablesForEndpoint(endpointID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Variables For Endpoint",
				Message:   "Failed to list all variables associated with an endpoint",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, variables))
	}
	return err
}

func getVariable(ctx echo.Context) (err error) {
	variableID := ctx.Param("variableID")
	variable, err := data.GetStore().GetVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get Variable Details",
				Message:   "Failed to fetch information about a variable",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, variable))
	}
	return err
}

func createVariable(ctx echo.Context) (err error) {
	var variable data.Variable
	err = ctx.Bind(&variable)
	if err == nil {
		err = data.GetStore().CreateVariable(&variable)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create Variable",
				Message:   "Failed to create variable",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create Variable",
				Message:   "Variable created succesffully",
				Error:     str(err)})
	}
	return err
}

func updateVariable(ctx echo.Context) (err error) {
	var variable data.Variable
	err = ctx.Bind(&variable)
	if err == nil {
		err = data.GetStore().UpdateVariable(&variable)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update Variable",
				Message:   "Failed to update variable",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update Variable",
				Message:   "Variable updated succesffully",
				Error:     str(err)})
	}
	return err
}

func deleteVariable(ctx echo.Context) (err error) {
	variableID := ctx.Param("variableID")
	err = data.GetStore().DeleteVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Dalete Variable",
				Message:   "Failed to delete variable",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete Variable",
				Message:   "Variable delete succesffully",
				Error:     str(err)})
	}
	return err
}

func getAllParameters(ctx echo.Context) (err error) {
	parameters, err := data.GetStore().GetAllParameters()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Parameters",
				Message:   "Failed to list all parameters",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, parameters))
	}
	return err
}

func getParametersForEndpoint(ctx echo.Context) (err error) {
	endpointID := ctx.Param("endpointID")
	parameters, err := data.GetStore().GetParametersForEndpoint(endpointID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Parameters For Endpoint",
				Message:   "Failed to list all parameters associated with an endpoint",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, parameters))
	}
	return err
}

func getParameter(ctx echo.Context) (err error) {
	parameterID := ctx.Param("parameterID")
	parameter, err := data.GetStore().GetParameter(parameterID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get Parameter Details",
				Message:   "Failed to fetch information about a parameter",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, parameter))
	}
	return err
}

func createParameter(ctx echo.Context) (err error) {
	var parameter data.Parameter
	err = ctx.Bind(&parameter)
	if err == nil {
		err = data.GetStore().CreateParameter(&parameter)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create Parameter",
				Message:   "Failed to create parameter",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create Parameter",
				Message:   "Parameter created succesffully",
				Error:     str(err)})
	}
	return err
}

func updateParameter(ctx echo.Context) (err error) {
	var parameter data.Parameter
	err = ctx.Bind(&parameter)
	if err == nil {
		err = data.GetStore().UpdateParameter(&parameter)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update Parameter",
				Message:   "Failed to update parameter",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update Parameter",
				Message:   "Parameter updated succesffully",
				Error:     str(err)})
	}
	return err
}

func deleteParameter(ctx echo.Context) (err error) {
	parameterID := ctx.Param("parameterID")
	err = data.GetStore().DeleteParameter(parameterID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Dalete Parameter",
				Message:   "Failed to delete parameter",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete Parameter",
				Message:   "Parameter delete succesffully",
				Error:     str(err)})
	}
	return err
}

func getAllUserGroups(ctx echo.Context) (err error) {
	userGroups, err := data.GetStore().GetAllUserGroups()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get UserGroup List",
				Message:   "Failed to fetch list of userGroup",
				Error:     str(err)})

	} else {
		err = logIfError(ctx.JSON(http.StatusOK, userGroups))
	}
	return err
}

func getUserGroup(ctx echo.Context) (err error) {
	groupID := ctx.Param("groupID")
	userGroup, err := data.GetStore().GetUserGroup(groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get UserGroup Details",
				Message:   "Failed to fetch userGroup details",
				Error:     str(err)})

	} else {
		err = logIfError(ctx.JSON(http.StatusOK, userGroup))
	}
	return err
}

func createUserGroup(ctx echo.Context) (err error) {
	var userGroup data.UserGroup
	err = ctx.Bind(&userGroup)
	if err == nil {
		err = data.GetStore().CreateUserGroup(&userGroup)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create UserGroup",
				Message:   "UserGroup creation failed",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create UserGroup",
				Message:   "UserGroup creation successful",
				Error:     ""})
	}
	return err
}

func updateUserGroup(ctx echo.Context) (err error) {
	var userGroup data.UserGroup
	err = ctx.Bind(&userGroup)
	if err == nil {
		err = data.GetStore().UpdateUserGroup(&userGroup)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update UserGroup",
				Message:   "UserGroup update failed",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update UserGroup",
				Message:   "UserGroup update successful",
				Error:     ""})
	}
	return err
}

func deleteUserGroup(ctx echo.Context) (err error) {
	groupID := ctx.Param("groupID")
	err = data.GetStore().DeleteUserGroup(groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Delete UserGroup",
				Message:   "UserGroup deletion failed",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete UserGroup",
				Message:   "UserGroup deletion successful",
				Error:     ""})
	}
	return err
}

func addUserToGroup(ctx echo.Context) (err error) {
	userName := ctx.Param("userName")
	groupID := ctx.Param("groupID")
	err = data.GetStore().AddUserToGroup(userName, groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Associate User to Group",
				Message:   "Failed to associate user to a group",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Associate User to Group",
				Message:   "User associated with a group successfully",
				Error:     ""})
	}
	return err
}

func removeUserFromGroup(ctx echo.Context) (err error) {
	userName := ctx.Param("userName")
	groupID := ctx.Param("groupID")
	err = data.GetStore().RemoveUserFromGroup(userName, groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Disassociate User to Group",
				Message:   "Failed to disaassociate user to a group",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Associate User to Group",
				Message:   "User associated with a group successfully",
				Error:     ""})
	}
	return err
}

func getUsersInGroup(ctx echo.Context) (err error) {
	groupID := ctx.Param("groupID")
	users, err := data.GetStore().GetUsersInGroup(groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List User in Group",
				Message:   "Failed to fetch users associated with a group",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, users))
	}
	return err
}

func getGroupsForUser(ctx echo.Context) (err error) {
	userName := ctx.Param("userName")
	groups, err := data.GetStore().GetGroupsForUser(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Groups of User",
				Message:   "Failed to fetch groups to which a user is  associated",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, groups))
	}
	return err
}

func addVariableValue(ctx echo.Context) (err error) {
	varValue := struct {
		VariableID string `json:"variableID"`
		Value      string `json:"value"`
	}{}
	err = ctx.Bind(&varValue)
	if err == nil {
		err = data.GetStore().AddVariableValue(varValue.VariableID,
			varValue.Value)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Add Variable Value",
				Message:   "Failed add a value for a variable",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Add Variable Value",
				Message:   "Value for a variable added",
				Error:     str(err)})
	}
	return err
}

func getValuesForVariable(ctx echo.Context) (err error) {
	variableID := ctx.Param("variableID")
	values, err := data.GetStore().GetValuesForVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Fetch Variable Values",
				Message:   "Failed fetch values of a variable",
				Error:     str(err)})
	} else {
		err = logIfError(ctx.JSON(http.StatusOK, values))
	}
	return err
}

func clearValuesForVariable(ctx echo.Context) (err error) {
	variableID := ctx.Param("variableID")
	err = data.GetStore().ClearValuesForVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Clear Variable Values",
				Message:   "Failed clear the values of a variable",
				Error:     str(err)})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Clear Variable Values",
				Message:   "Cleared values of a variable",
				Error:     ""})
	}
	logIfError(err)
	return err
}

func setPassword(ctx echo.Context) (err error) {
	//This needs to run only if the current logged in user is an admin
	//OR email user the link to set the password
	userName := ctx.FormValue("userName")
	password := ctx.FormValue("password")
	var hash string
	hash, err = passlib.Hash(password)
	if err == nil {
		err = data.GetStore().SetPasswordHash(userName, hash)
		if err == nil {
			ctx.JSON(http.StatusOK,
				Result{
					Operation: "SetPassword",
					Message:   "Password set successfully",
					Error:     str(err)})
		}
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "SetPassword",
				Message:   "Error occured while setting password",
				Error:     ""})
	}
	return err
}

func updatePassword(ctx echo.Context) (err error) {
	userName := ctx.FormValue("userName")
	oldPassword := ctx.FormValue("oldPassword")
	password := ctx.FormValue("password")
	err = varifyPassword(userName, oldPassword)
	if err == nil {
		var hash string
		hash, err = passlib.Hash(password)
		if err == nil {
			err = data.GetStore().SetPasswordHash(userName, hash)
			if err == nil {
				ctx.JSON(http.StatusInternalServerError,
					Result{
						Operation: "UpdatePassword",
						Message:   "Password updated successfully",
						Error:     str(err)})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError,
				Result{
					Operation: "UpdatePassword",
					Message:   "Error occured while updating password",
					Error:     str(err)})
		}
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "UpdatePassword",
				Message:   "Error occured while updating password",
				Error:     str(err)})
	}
	return err
}

func varifyPassword(userName, password string) (err error) {
	var hash string
	hash, err = data.GetStore().GetPasswordHash(userName)
	if err == nil {
		var newHash string
		newHash, err = passlib.Verify(password, hash)
		if err == nil && newHash != "" {
			err = data.GetStore().SetPasswordHash(userName, newHash)
		}
	}
	logIfError(err)
	return err
}

func login(ctx echo.Context) (err error) {
	userName := ctx.FormValue("username")
	password := ctx.FormValue("password")
	if err = varifyPassword(userName, password); err == nil {
		// 	userExists, checkError := data.GetStore().UserExists(userName)
		// 	if checkError == nil && !userExists {
		// 		ctx.JSON(http.StatusUnauthorized,
		// 			Result{
		// 				Operation: "login",
		// 				Message:   "Create password",
		// 				Error: str(err)ors.New("NoPasswordFound"),
		// 			})
		// 		err = nil
		// 	} else {
		// 		if !userExists {
		// 			olog.Error("REST", "User %s does not exist", userName)
		// 			err = echo.ErrUnauthorized
		// 		} else {
		// 			// err = checkError
		// 		}
		// 	}
		// } else {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = userName
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(ky))
		if err != nil {
			return err
		}
		err = ctx.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
		olog.Print("REST", "## %s", t)
	} else {
		ctx.JSON(http.StatusUnauthorized,
			Result{
				Operation: "Login",
				Message:   "Invalid user name or password",
				Error:     str(err)})
	}
	logIfError(err)
	return err
}

func defaultHandler(ctx echo.Context) (err error) {
	err = ctx.HTML(http.StatusOK, "<b> Hello from Orek </b>")
	return err
}

//Serve - maps a route to handler function and starts the server
func Serve(port int) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	v0 := e.Group("/v0")

	v0.POST("/login", login)
	in0 := v0.Group("/in")
	in0.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(ky),
	}))

	in0.GET("/users", getAllUsers)
	in0.GET("/users/:userName", getUser)
	in0.GET("/users/email/:email", getUser)
	in0.POST("/users", createUser)
	in0.PUT("/users", updateUser)
	in0.DELETE("/users/:userName", deleteUser)

	in0.GET("/endpoints", getAllEndpoints)
	in0.GET("/endpoints/:endpointID", getEndpoint)
	in0.POST("/endpoints", createEndpoint)
	in0.PUT("/endpoints", updateEndpoint)
	in0.DELETE("/endpoints/:endpointID", deleteEndpoint)

	in0.GET("/variables", getAllVariables)
	in0.GET("/endpoints/:endpointID/variables", getVariablesForEndpoint)
	in0.GET("/variables/:variableID", getVariable)
	in0.POST("/variables", createVariable)
	in0.PUT("/variables", updateVariable)
	in0.DELETE("/variables/:endpointID", deleteVariable)

	in0.GET("/parameters", getAllParameters)
	in0.GET("/endpoints/:endpointID/parameters", getParametersForEndpoint)
	in0.GET("/parameters/:parameterID", getParameter)
	in0.POST("/parameters", createParameter)
	in0.PUT("/parameters", updateParameter)
	in0.DELETE("/parameters/:parameterID", deleteParameter)

	in0.GET("/groups", getAllUserGroups)
	in0.GET("/groups/:groupID", getUserGroup)
	in0.POST("/groups", createUserGroup)
	in0.PUT("/groups", updateUserGroup)
	in0.DELETE("/groups/:groupID", deleteUserGroup)
	in0.PUT("/groups/:groupID/users/:userName", addUserToGroup)
	in0.DELETE("/groups/:groupID/users/:userName", removeUserFromGroup)
	in0.GET("/groups/:groupID/users", getUsersInGroup)
	in0.GET("/users/:userName/groups", getGroupsForUser)

	in0.POST("/varialbes/:variableID/values", addVariableValue)
	in0.GET("/variables/:variableID/values", getValuesForVariable)
	in0.DELETE("/variables/:variableID/values", clearValuesForVariable)
	in0.POST("/manageAuth", setPassword)
	in0.PUT("/manageAuth", updatePassword)

	portStr := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(portStr))
}
