package previewawswafv2mixins


// A logical rule statement used to negate the results of another rule statement.
//
// You provide one `Statement` within the `NotStatement` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-notstatement.html
//
type CfnWebACLPropsMixin_NotStatementProperty struct {
	// The statement to negate.
	//
	// You can use any statement that can be nested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-notstatement.html#cfn-wafv2-webacl-notstatement-statement
	//
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}

