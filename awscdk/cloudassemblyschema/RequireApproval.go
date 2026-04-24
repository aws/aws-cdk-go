package cloudassemblyschema


// Specify what changes require manual approval.
type RequireApproval string

const (
	// Approval is not required.
	RequireApproval_NEVER RequireApproval = "NEVER"
	// Manual approval required for any change to the stack.
	RequireApproval_ANYCHANGE RequireApproval = "ANYCHANGE"
	// Manual approval required if changes involve a broadening of permissions or security group rules.
	RequireApproval_BROADENING RequireApproval = "BROADENING"
)

