package awswafv2


// A logical rule statement used to combine other rule statements with AND logic.
//
// You provide more than one `Statement` within the `AndStatement` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-andstatement.html
//
type CfnWebACL_AndStatementProperty struct {
	// The statements to combine with AND logic.
	//
	// You can use any statements that can be nested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-andstatement.html#cfn-wafv2-webacl-andstatement-statements
	//
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

