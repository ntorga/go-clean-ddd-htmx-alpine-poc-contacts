package apiController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/useCase"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra"
	apiHelper "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api/helper"
)

// GetContacts	 godoc
// @Summary      GetContacts
// @Description  List contacts.
// @Tags         contact
// @Accept       json
// @Produce      json
// @Success      200 {array} entity.Contact
// @Router       /v1/contact/ [get]
func GetContactsController(c echo.Context) error {
	contactsQueryRepo := infra.NewContactQueryRepo()
	contactsList, err := useCase.GetContacts(contactsQueryRepo)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, contactsList)
}
