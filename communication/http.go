package communication

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//userFacingError will return the occurred error as a custom html page to the caller.
func userFacingError(w http.ResponseWriter, errorMessage string) {
	err := errorPage.ExecuteTemplate(w, "error.html", errorMessage)
	//This should never happen, but if it does, something is very wrong.
	if err != nil {
		panic(err)
	}
}

func remoteAddressToSimpleIP(input string) string {
	address := input
	lastIndexOfDoubleColon := strings.LastIndex(address, ":")
	if lastIndexOfDoubleColon != -1 {
		address = address[:lastIndexOfDoubleColon]
	}

	return strings.TrimSuffix(strings.TrimPrefix(address, "["), "]")
}

func Serve(port int) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
