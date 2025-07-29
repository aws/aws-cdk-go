package awscdkclilibalpha


// In what scenarios should the CLI ask for approval.
// Deprecated.
type RequireApproval string

const (
	// Never ask for approval.
	// Deprecated.
	RequireApproval_NEVER RequireApproval = "NEVER"
	// Prompt for approval for any type  of change to the stack.
	// Deprecated.
	RequireApproval_ANYCHANGE RequireApproval = "ANYCHANGE"
	// Only prompt for approval if there are security related changes.
	// Deprecated.
	RequireApproval_BROADENING RequireApproval = "BROADENING"
)

