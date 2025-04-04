package awswafv2


// A logical rule statement used to combine other rule statements with OR logic.
//
// You provide more than one `Statement` within the `OrStatement` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-orstatement.html
//
type CfnWebACL_OrStatementProperty struct {
	// The statements to combine with OR logic.
	//
	// You can use any statements that can be nested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-orstatement.html#cfn-wafv2-webacl-orstatement-statements
	//
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

