package awswafv2


// A logical rule statement used to negate the results of another rule statement.
//
// You provide one `Statement` within the `NotStatement` .
//
// Example:
//
//
type CfnWebACL_NotStatementProperty struct {
	// The statement to negate.
	//
	// You can use any statement that can be nested.
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
}

