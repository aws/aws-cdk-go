package awswafv2


// A logical rule statement used to combine other rule statements with AND logic.
//
// You provide more than one `Statement` within the `AndStatement` .
//
// Example:
//
//
type CfnWebACL_AndStatementProperty struct {
	// The statements to combine with AND logic.
	//
	// You can use any statements that can be nested.
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

