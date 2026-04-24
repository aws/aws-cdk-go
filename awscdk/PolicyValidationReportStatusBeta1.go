package awscdk


// The final status of the validation report.
// Deprecated: Use `PolicyValidationReportStatus` instead.
type PolicyValidationReportStatusBeta1 string

const (
	// No violations were found.
	// Deprecated: Use `PolicyValidationReportStatus` instead.
	PolicyValidationReportStatusBeta1_SUCCESS PolicyValidationReportStatusBeta1 = "SUCCESS"
	// At least one violation was found.
	// Deprecated: Use `PolicyValidationReportStatus` instead.
	PolicyValidationReportStatusBeta1_FAILURE PolicyValidationReportStatusBeta1 = "FAILURE"
)

