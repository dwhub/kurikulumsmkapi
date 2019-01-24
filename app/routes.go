package app

import (
	"net/http"
	"os"

	c "github.com/dwhub/kurikulumsmkapi/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// GetRouter initiate system router
func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/v1/health", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.CheckHealth))).Methods("GET")

	router.Handle("/v1/contacts", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetContacts))).Methods("GET")

	router.Handle("/v1/Contacts/all", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetAllContacts))).Methods("GET")

	router.Handle("/v1/provinces", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetAllProvinces))).Methods("GET")

	router.Handle("/v1/districts", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetDistricts))).Methods("GET")

	router.Handle("/v1/districts/with/province/{provinceId:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetDistrictByProvinceID))).Methods("GET")

	router.Handle("/v1/schools", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetSchools))).Methods("GET")

	router.Handle("/v1/schools/all", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetAllSchools))).Methods("GET")

	router.Handle("/v1/expertiseFields", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetExpertiseFields))).Methods("GET")

	router.Handle("/v1/expertiseFields/curriculumStructure", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetCurriculumStructures))).Methods("GET")

	router.Handle("/v1/expertisePrograms", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetExpertisePrograms))).Methods("GET")

	router.Handle("/v1/expertisePrograms/with/expertiseField/{fieldId:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetExpertiseProgramsByFieldID))).Methods("GET")

	router.Handle("/v1/expertiseCompetencies", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetExpertiseCompetencies))).Methods("GET")

	router.Handle("/v1/expertiseCompetencies/with/expertiseProgram/{programId:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetExpertiseCompetenciesByProgramID))).Methods("GET")

	router.Handle("/v1/course/duration/with/competency/{competencyId:[0-9]+}/group/{groupId:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetCourseDurations))).Methods("GET")

	router.Handle("/v1/course/allocation/with/competency/{competencyId:[0-9]+}/group/{groupId:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetCourseAllocations))).Methods("GET")

	router.Handle("/v1/course/kikd/detail/with/competency/{competencyId:[0-9]+}/course/{courseId:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(c.GetKIKDDetails))).Methods("GET")

	return router
}
