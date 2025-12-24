package tense

import (
	"fmt"
	"net/http"

	"github.com/kasariks/a_nameless_project_yet/types"
	"github.com/kasariks/a_nameless_project_yet/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /gettense", h.GetTenseHandler)
}

func (h *Handler) GetTenseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Query().Get("tense") {
	case "simplepresent":
		utils.WriteJSON(w, http.StatusOK, &types.Tense{
			Name:      "Simple Present Tense",
			Structure: "Subject + verb(v1) + s/es",
			Example:   "She goes to the shop.",
		})
	case "presentcontinuos":
		utils.WriteJSON(w, http.StatusOK, &types.Tense{
			Name:      "Present Continuos Tense",
			Structure: "Subject + is/am/are + verb(ing)",
			Example:   "I am studing in a high school.",
		})
	case "presentperfect":
		utils.WriteJSON(w, http.StatusOK, &types.Tense{
			Name:      "Present Perfect Tense",
			Structure: "Subject + has/have + verb(v3)",
			Example:   "He has made his homework.",
		})
	case "presentperfectcontinuos":
		utils.WriteJSON(w, http.StatusOK, &types.Tense{
			Name:      "Present Perfect Continuos Tense",
			Structure: "Subject + has/have + been + verb(ing) + since/for",
			Example:   "It has been working since yesterday.",
		})
	default:
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing tense parameter"))
		return
	}
}
