package rest

import (
	"net/http"

	"log"

	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	irisjwt "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/varunamachi/orekng/data"
)

const ky = "orek_2232redsfaj3234edsa"

func logIfError(err error) {
	if err != nil {
		log.Printf("Error:REST: %v", err)
	}
}

func getAllUsers(ctx *iris.Context) {
	users, err := data.GetDataStore().GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get User List",
				Message:   "Failed to fetch list of user",
				Error:     err})

	} else {
		logIfError(ctx.JSON(http.StatusOK, users))
	}
}

func getUser(ctx *iris.Context) {
	userName := ctx.Param("userName")
	user, err := data.GetDataStore().GetUser(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get User Details",
				Message:   "Failed to fetch user details",
				Error:     err})

	} else {
		logIfError(ctx.JSON(http.StatusOK, user))
	}
}

func getUserWithEmail(ctx *iris.Context) {
	email := ctx.Param("email")
	user, err := data.GetDataStore().GetUserWithEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get User Details",
				Message:   "Failed to fetch user details",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, user))
	}

}

func createUser(ctx *iris.Context) {
	var user data.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		err = data.GetDataStore().CreateUser(&user)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create User",
				Message:   "User creation failed",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create User",
				Message:   "User creation successful",
				Error:     nil})
	}

}

func updateUser(ctx *iris.Context) {
	var user data.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		err = data.GetDataStore().UpdateUser(&user)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update User",
				Message:   "User update failed",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update User",
				Message:   "User update successful",
				Error:     nil})
	}
}

func deleteUser(ctx *iris.Context) {
	userName := ctx.Param("userName")
	err := data.GetDataStore().DeleteUser(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Delete User",
				Message:   "User deletion failed",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete User",
				Message:   "User deletion successful",
				Error:     nil})
	}
}

func getAllEndpoints(ctx *iris.Context) {
	endpoints, err := data.GetDataStore().GetAllEndpoints()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Endpoint",
				Message:   "Listing endpoint failed",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, endpoints))
	}
}

func getEndpoint(ctx *iris.Context) {
	endpointID := ctx.Param("endpointID")
	endpoint, err := data.GetDataStore().GetEndpoint(endpointID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get Endpoint Details",
				Message:   "Failed to fetch endpoint details",
				Error:     err})

	} else {
		logIfError(ctx.JSON(http.StatusOK, endpoint))
	}
}

func createEndpoint(ctx *iris.Context) {
	var endpoint data.Endpoint
	err := ctx.ReadJSON(&endpoint)
	if err != nil {
		err = data.GetDataStore().CreateEndpoint(&endpoint)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create Endpoint",
				Message:   "Failed to create endpoint",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create Endpoint",
				Message:   "Endpoint created succesffully",
				Error:     err})
	}
}

func updateEndpoint(ctx *iris.Context) {
	var endpoint data.Endpoint
	err := ctx.ReadJSON(&endpoint)
	if err != nil {
		err = data.GetDataStore().UpdateEndpoint(&endpoint)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update Endpoint",
				Message:   "Failed to update endpoint",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update Endpoint",
				Message:   "Endpoint updated succesffully",
				Error:     err})
	}
}

func deleteEndpoint(ctx *iris.Context) {
	endpointID := ctx.Param("endpointID")
	err := data.GetDataStore().DeleteEndpoint(endpointID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Dalete Endpoint",
				Message:   "Failed to delete endpoint",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete Endpoint",
				Message:   "Endpoint delete succesffully",
				Error:     err})
	}
}

func getAllVariables(ctx *iris.Context) {
	variables, err := data.GetDataStore().GetAllVariables()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Variables",
				Message:   "Failed to list all variables",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, variables))
	}
}

func getVariablesForEndpoint(ctx *iris.Context) {
	endpointID := ctx.Param("endpointID")
	variables, err := data.GetDataStore().GetVariablesForEndpoint(endpointID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Variables For Endpoint",
				Message:   "Failed to list all variables associated with an endpoint",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, variables))
	}
}

func getVariable(ctx *iris.Context) {
	variableID := ctx.Param("variableID")
	variable, err := data.GetDataStore().GetVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get Variable Details",
				Message:   "Failed to fetch information about a variable",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, variable))
	}

}

func createVariable(ctx *iris.Context) {
	var variable data.Variable
	err := ctx.ReadJSON(&variable)
	if err != nil {
		err = data.GetDataStore().CreateVariable(&variable)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create Variable",
				Message:   "Failed to create variable",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create Variable",
				Message:   "Variable created succesffully",
				Error:     err})
	}
}

func updateVariable(ctx *iris.Context) {
	var variable data.Variable
	err := ctx.ReadJSON(&variable)
	if err != nil {
		err = data.GetDataStore().UpdateVariable(&variable)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update Variable",
				Message:   "Failed to update variable",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update Variable",
				Message:   "Variable updated succesffully",
				Error:     err})
	}
}

func deleteVariable(ctx *iris.Context) {
	variableID := ctx.Param("variableID")
	err := data.GetDataStore().DeleteVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Dalete Variable",
				Message:   "Failed to delete variable",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete Variable",
				Message:   "Variable delete succesffully",
				Error:     err})
	}
}

func getAllParameters(ctx *iris.Context) {
	parameters, err := data.GetDataStore().GetAllParameters()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Parameters",
				Message:   "Failed to list all parameters",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, parameters))
	}
}

func getParametersForEndpoint(ctx *iris.Context) {
	endpointID := ctx.Param("endpointID")
	parameters, err := data.GetDataStore().GetParametersForEndpoint(endpointID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Parameters For Endpoint",
				Message:   "Failed to list all parameters associated with an endpoint",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, parameters))
	}
}

func getParameter(ctx *iris.Context) {
	parameterID := ctx.Param("parameterID")
	parameter, err := data.GetDataStore().GetParameter(parameterID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get Parameter Details",
				Message:   "Failed to fetch information about a parameter",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, parameter))
	}

}

func createParameter(ctx *iris.Context) {
	var parameter data.Parameter
	err := ctx.ReadJSON(&parameter)
	if err != nil {
		err = data.GetDataStore().CreateParameter(&parameter)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create Parameter",
				Message:   "Failed to create parameter",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create Parameter",
				Message:   "Parameter created succesffully",
				Error:     err})
	}
}

func updateParameter(ctx *iris.Context) {
	var parameter data.Parameter
	err := ctx.ReadJSON(&parameter)
	if err != nil {
		err = data.GetDataStore().UpdateParameter(&parameter)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update Parameter",
				Message:   "Failed to update parameter",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update Parameter",
				Message:   "Parameter updated succesffully",
				Error:     err})
	}
}

func deleteParameter(ctx *iris.Context) {
	parameterID := ctx.Param("parameterID")
	err := data.GetDataStore().DeleteParameter(parameterID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Dalete Parameter",
				Message:   "Failed to delete parameter",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete Parameter",
				Message:   "Parameter delete succesffully",
				Error:     err})
	}
}

// func getAllUserGroups(ctx *iris.Context) {

// }

// func getUserGroup(ctx *iris.Context) {

// }

// func createUserGroup(ctx *iris.Context) {

// }

// func updateUserGroup(ctx *iris.Context) {

// }

// func deleteUserGroup(ctx *iris.Context) {

// }

func getAllUserGroups(ctx *iris.Context) {
	userGroups, err := data.GetDataStore().GetAllUserGroups()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get UserGroup List",
				Message:   "Failed to fetch list of userGroup",
				Error:     err})

	} else {
		logIfError(ctx.JSON(http.StatusOK, userGroups))
	}
}

func getUserGroup(ctx *iris.Context) {
	groupID := ctx.Param("groupID")
	userGroup, err := data.GetDataStore().GetUserGroup(groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Get UserGroup Details",
				Message:   "Failed to fetch userGroup details",
				Error:     err})

	} else {
		logIfError(ctx.JSON(http.StatusOK, userGroup))
	}
}

func createUserGroup(ctx *iris.Context) {
	var userGroup data.UserGroup
	err := ctx.ReadJSON(&userGroup)
	if err != nil {
		err = data.GetDataStore().CreateUserGroup(&userGroup)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Create UserGroup",
				Message:   "UserGroup creation failed",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Create UserGroup",
				Message:   "UserGroup creation successful",
				Error:     nil})
	}

}

func updateUserGroup(ctx *iris.Context) {
	var userGroup data.UserGroup
	err := ctx.ReadJSON(&userGroup)
	if err != nil {
		err = data.GetDataStore().UpdateUserGroup(&userGroup)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Update UserGroup",
				Message:   "UserGroup update failed",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Update UserGroup",
				Message:   "UserGroup update successful",
				Error:     nil})
	}
}

func deleteUserGroup(ctx *iris.Context) {
	groupID := ctx.Param("groupID")
	err := data.GetDataStore().DeleteUserGroup(groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Delete UserGroup",
				Message:   "UserGroup deletion failed",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Delete UserGroup",
				Message:   "UserGroup deletion successful",
				Error:     nil})
	}
}

func addUserToGroup(ctx *iris.Context) {
	userName := ctx.Param("userName")
	groupID := ctx.Param("groupID")
	err := data.GetDataStore().AddUserToGroup(userName, groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Associate User to Group",
				Message:   "Failed to associate user to a group",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Associate User to Group",
				Message:   "User associated with a group successfully",
				Error:     nil})
	}
}

func removeUserFromGroup(ctx *iris.Context) {
	userName := ctx.Param("userName")
	groupID := ctx.Param("groupID")
	err := data.GetDataStore().RemoveUserFromGroup(userName, groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Disassociate User to Group",
				Message:   "Failed to disaassociate user to a group",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Associate User to Group",
				Message:   "User associated with a group successfully",
				Error:     nil})
	}
}

func getUsersInGroup(ctx *iris.Context) {
	groupID := ctx.Param("groupID")
	users, err := data.GetDataStore().GetUsersInGroup(groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List User in Group",
				Message:   "Failed to fetch users associated with a group",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, users))
	}
}

func getGroupsForUser(ctx *iris.Context) {
	userName := ctx.Param("userName")
	groups, err := data.GetDataStore().GetGroupsForUser(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "List Groups of User",
				Message:   "Failed to fetch groups to which a user is  associated",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, groups))
	}
}

func addVariableValue(ctx *iris.Context) {
	varValue := struct {
		VariableID string `json:"variableID"`
		Value      string `json:"value"`
	}{}
	err := ctx.ReadJSON(&varValue)
	if err != nil {
		err = data.GetDataStore().AddVariableValue(varValue.VariableID,
			varValue.Value)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Add Variable Value",
				Message:   "Failed add a value for a variable",
				Error:     err})
	} else {
		ctx.JSON(http.StatusOK,
			Result{
				Operation: "Add Variable Value",
				Message:   "Value for a variable added",
				Error:     err})
	}
}

func getValuesForVariable(ctx *iris.Context) {
	variableID := ctx.Param("variableID")
	values, err := data.GetDataStore().GetValuesForVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Fetch Variable Values",
				Message:   "Failed fetch values of a variable",
				Error:     err})
	} else {
		logIfError(ctx.JSON(http.StatusOK, values))
	}
}

func clearValuesForVariable(ctx *iris.Context) {
	variableID := ctx.Param("variableID")
	err := data.GetDataStore().ClearValuesForVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Clear Variable Values",
				Message:   "Failed clear the values of a variable",
				Error:     err})
	} else {
		ctx.JSON(http.StatusInternalServerError,
			Result{
				Operation: "Clear Variable Values",
				Message:   "Cleared values of a variable",
				Error:     err})
	}
}

func login(ctx *iris.Context) {

}

func defaultHandler(ctx *iris.Context) {
	ctx.Write([]byte("Hello from Orek!!"))
}

func fromCookie(ctx *iris.Context) (token string, err error) {
	token = ""
	return token, err
}

func errorHandler(ctx *iris.Context, errString string) {
	ctx.JSON(http.StatusUnauthorized, Result{
		Operation: "Authorization",
		Message:   fmt.Sprintf("JWT Authorization failed! %s", errString),
		Error:     nil,
	})
}

//Map - maps a route to handler function
func Map() {
	extractor := irisjwt.FromFirst(irisjwt.FromAuthHeader, fromCookie)
	jwtMiddleWare := irisjwt.New(irisjwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(ky), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		Extractor:     extractor,
		ErrorHandler:  errorHandler,
	})
	iris.Get("/", jwtMiddleWare.Serve, defaultHandler)
}
