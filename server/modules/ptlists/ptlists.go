package ptlists

import (
	"net/http"

	"plist/internal/app/ptlist"

	pfhttp "plist/pkg/http"

	"github.com/gorilla/mux"
)

// Module struct.
type Module struct {
	ptlistService *ptlist.Service
}

// Setup registers the Ptlists module to the router.
func Setup(router *mux.Router, ptlistService *ptlist.Service) {
	m := &Module{
		ptlistService: ptlistService,
	}

	router.HandleFunc("/ptlist", m.GetPtList).Methods("GET")
}

// GetPtList.
func (m *Module) GetPtList(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Get expected url query values.
	values := r.URL.Query()
	period := values.Get("period")
	tz := values.Get("tz")
	t1 := values.Get("t1")
	t2 := values.Get("t2")

	// Call ptlist service.
	ptlist, err := m.ptlistService.GetPtList(
		ctx,
		period,
		tz,
		t1,
		t2,
	)

	// Handle error.
	if err != nil {
		pfhttp.WriteJSON(http.StatusInternalServerError, err, w)
		return
	}

	pfhttp.WriteJSON(http.StatusOK, ptlist, w)
}
