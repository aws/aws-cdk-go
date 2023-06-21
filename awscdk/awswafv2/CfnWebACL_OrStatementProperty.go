package awswafv2


// A logical rule statement used to combine other rule statements with OR logic.
//
// You provide more than one `Statement` within the `OrStatement` .
//
// Example:
//
//
type CfnWebACL_OrStatementProperty struct {
	// The statements to combine with OR logic.
	//
	// You can use any statements that can be nested.
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

