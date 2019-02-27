// Package swagger Kurikulum SMK API.
//
//     Schemes: http
//     BasePath: /v1
//     Version: 1.0.0
//     Contact: Andi Yudistira<andi.yudistira81@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package swagger

// swagger:operation GET /health systems systemHealth
// ---
// summary: Return the status of the service health.
// description: Get the status of the service health.
// responses:
//   "200":
//     "$ref": "#/responses/healthResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/healthResp"

// swagger:operation GET /contacts contacts listContacts
// ---
// summary: Return list of contacts with paging.
// description: Get the contact entities with paging.
// responses:
//   "200":
//     "$ref": "#/responses/contactPagingResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/contactPagingResp"

// swagger:operation GET /contacts/all contacts listAllContacts
// ---
// summary: Return list of contacts without paging.
// description: Get the contact entities without paging.
// responses:
//   "200":
//     "$ref": "#/responses/contactResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/contactResp"

// swagger:operation GET /provinces provinces listProvinces
// ---
// summary: Return list of provinces.
// description: Get the province entities.
// responses:
//   "200":
//     "$ref": "#/responses/provinceResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/provinceResp"

// swagger:operation GET /districts districts listDistricts
// ---
// summary: Return list of districts.
// description: Get the district entities.
// responses:
//   "200":
//     "$ref": "#/responses/districtsResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/districtsResp"

// swagger:operation GET /districts/with/province/{provinceId} districts getDistrictByProvinceID
// ---
// summary: Return districts by Province id
// description: Get the district entity by Province id.
// parameters:
// - name: provinceId
//   in: path
//   description: ID of the province
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/districtsResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/districtsResp"

// swagger:operation GET /schools schools listSchools
// ---
// summary: Return list of schools with paging.
// description: Get the school entities with paging.
// responses:
//   "200":
//     "$ref": "#/responses/schoolPagingResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/schoolPagingResp"

// swagger:operation GET /schools/all schools listAllSchools
// ---
// summary: Return list of schools without paging.
// description: Get the school entities without paging.
// responses:
//   "200":
//     "$ref": "#/responses/schoolResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/schoolResp"

// swagger:operation GET /expertiseFields expertises listExpertiseFields
// ---
// summary: Return list of expertise fields.
// description: Get the expertise field entities.
// responses:
//   "200":
//     "$ref": "#/responses/expertiseFieldResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/expertiseFieldResp"

// swagger:operation GET /expertiseFields/curriculumStructure expertises listCurriculumStructure
// ---
// summary: Return list of curriculum structure.
// description: Get the curriculum structure field entities.
// responses:
//   "200":
//     "$ref": "#/responses/curriculumStructureResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/curriculumStructureResp"

// swagger:operation GET /expertisePrograms expertises listExpertisePrograms
// ---
// summary: Return list of expertise programs.
// description: Get the expertise program entities.
// responses:
//   "200":
//     "$ref": "#/responses/expertiseProgramsResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/expertiseProgramsResp"

// swagger:operation GET /expertisePrograms/with/expertiseField/{fieldId} expertises getExpertiseProgramsByFieldId
// ---
// summary: Return expertise program by field id
// description: Get the expertise program entity by field id.
// parameters:
// - name: fieldId
//   in: path
//   description: ID of the field
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/expertiseProgramsResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/expertiseProgramsResp"

// swagger:operation GET /expertiseCompetencies expertises listExpertisePrograms
// ---
// summary: Return list of expertise competencies.
// description: Get the expertise competency entities.
// responses:
//   "200":
//     "$ref": "#/responses/expertiseCompetenciesResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/expertiseCompetenciesResp"

// swagger:operation GET /expertiseCompetencies/with/expertiseProgram/{programId} expertises getExpertiseProgramsByFieldId
// ---
// summary: Return expertise competencies by program id
// description: Get the expertise competencies entity by program id.
// parameters:
// - name: programId
//   in: path
//   description: ID of the program
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/expertiseCompetenciesResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/expertiseCompetenciesResp"

// swagger:operation GET /course/duration/with/competency/{competencyId}/group/{groupId} courses getCourseDurations
// ---
// summary: Return course duration by Competency and Group id
// description: Get the course duration entity by Competency and Group id.
// parameters:
// - name: competencyId
//   in: path
//   description: ID of the competency
//   type: integer
//   required: true
// - name: groupId
//   in: path
//   description: ID of the group
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/courseDurationResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/courseDurationResp"

// swagger:operation GET /course/allocation/with/competency/{competencyId}/group/{groupId} courses getCourseAllocations
// ---
// summary: Return course allocation by Competency and Group id
// description: Get the course allocation entity by Competency and Group id.
// parameters:
// - name: competencyId
//   in: path
//   description: ID of the competency
//   type: integer
//   required: true
// - name: groupId
//   in: path
//   description: ID of the group
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/courseAllocationResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/courseAllocationResp"

// swagger:operation GET /course/kikd/with/competency/{competencyId}/group/{groupId} courses getCourseKIKD
// ---
// summary: Return course KI and KD by Competency and Group id
// description: Get the course KI and KD entity by Competency and Group id.
// parameters:
// - name: competencyId
//   in: path
//   description: ID of the competency
//   type: integer
//   required: true
// - name: groupId
//   in: path
//   description: ID of the group
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/courseKIKDResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/courseKIKDResp"

// swagger:operation GET /course/kikd/detail/with/competency/{competencyId}/course/{courseId} courses getKIKDDetail
// ---
// summary: Return KI and KD detail by Competency and Course id
// description: Get the KI and KD detail entity by Competency and Course id.
// parameters:
// - name: competencyId
//   in: path
//   description: ID of the competency
//   type: integer
//   required: true
// - name: courseId
//   in: path
//   description: ID of the course
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/KIKDDetailResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/KIKDDetailResp"

// swagger:operation GET /course/book/with/competency/{competencyId}/group/{groupId} courses getCourseBooks
// ---
// summary: Return course book by Competency and Group id
// description: Get the course book entity by Competency and Group id.
// parameters:
// - name: competencyId
//   in: path
//   description: ID of the competency
//   type: integer
//   required: true
// - name: groupId
//   in: path
//   description: ID of the group
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/courseBookResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/courseBookResp"

// swagger:operation GET /course/with/competency/{competencyId} courses getCourses
// ---
// summary: Return course by Competency id
// description: Get the course entity by Competency id.
// parameters:
// - name: competencyId
//   in: path
//   description: ID of the competency
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/courseResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/courseResp"

// swagger:operation GET /course/with/competency/{competencyId}/group/{groupId} courses getCoursesWithGroup
// ---
// summary: Return course by Competency and group id
// description: Get the course entity by Competency id and group id.
// parameters:
// - name: competencyId
//   in: path
//   description: ID of the competency
//   type: integer
//   required: true
// - name: groupId
//   in: path
//   description: ID of the group
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/courseResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/courseResp"

// swagger:operation GET /nationalExams nationalExams listNationalExams
// ---
// summary: Return list of national exams data.
// description: Get the national exams field entities.
// responses:
//   "200":
//     "$ref": "#/responses/nationalExamsResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/nationalExamsResp"

// swagger:operation GET /nes nationalEducationStandards listNationalEducationStandards
// ---
// summary: Return list of national education standards data.
// description: Get the national education standards field entities.
// responses:
//   "200":
//     "$ref": "#/responses/nesResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/nesResp"
