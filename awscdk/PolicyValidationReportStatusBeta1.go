package awscdk


// The final status of the validation report.
type PolicyValidationReportStatusBeta1 string

const (
	// No violations were found.
	PolicyValidationReportStatusBeta1_SUCCESS PolicyValidationReportStatusBeta1 = "SUCCESS"
	// At least one violation was found.
	PolicyValidationReportStatusBeta1_FAILURE PolicyValidationReportStatusBeta1 = "FAILURE"
)

