package awscdkclilibalpha


// In what scenarios should the CLI ask for approval.
// Experimental.
type RequireApproval string

const (
	// Never ask for approval.
	// Experimental.
	RequireApproval_NEVER RequireApproval = "NEVER"
	// Prompt for approval for any type  of change to the stack.
	// Experimental.
	RequireApproval_ANYCHANGE RequireApproval = "ANYCHANGE"
	// Only prompt for approval if there are security related changes.
	// Experimental.
	RequireApproval_BROADENING RequireApproval = "BROADENING"
)

