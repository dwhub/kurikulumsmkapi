// Golang SwaggerUI for Kurikulum SMK Service
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

// swagger:operation GET /expertiseFields expertiseFields listExpertiseFields
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

// swagger:operation GET /expertiseFields/curriculumStructure expertiseFields listCurriculumStructure
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

// swagger:operation GET /expertisePrograms expertisePrograms listExpertisePrograms
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

// swagger:operation GET /expertisePrograms/with/expertiseField/{fieldId} expertisePrograms getExpertiseProgramsByFieldId
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

// swagger:operation GET /expertiseCompetencies expertiseCompetencies listExpertisePrograms
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

// swagger:operation GET /expertiseCompetencies/with/expertiseProgram/{programId} expertiseCompetencies getExpertiseProgramsByFieldId
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