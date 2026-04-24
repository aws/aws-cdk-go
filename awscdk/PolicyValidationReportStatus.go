package awscdk


// The final status of the validation report.
type PolicyValidationReportStatus string

const (
	// No violations were found.
	PolicyValidationReportStatus_SUCCESS PolicyValidationReportStatus = "SUCCESS"
	// At least one violation was found.
	PolicyValidationReportStatus_FAILURE PolicyValidationReportStatus = "FAILURE"
)

