package cloudassemblyschema


// In what scenarios should the CLI ask for approval.
type RequireApproval string

const (
	// Never ask for approval.
	RequireApproval_NEVER RequireApproval = "NEVER"
	// Prompt for approval for any type  of change to the stack.
	RequireApproval_ANYCHANGE RequireApproval = "ANYCHANGE"
	// Only prompt for approval if there are security related changes.
	RequireApproval_BROADENING RequireApproval = "BROADENING"
)

