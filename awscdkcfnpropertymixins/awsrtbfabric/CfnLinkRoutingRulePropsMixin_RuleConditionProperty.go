package awsrtbfabric


// Conditions for a routing rule.
//
// All non-null fields must match (AND logic). At least one field must be set. HostHeader and HostHeaderWildcard are mutually exclusive. PathPrefix and PathExact are mutually exclusive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ruleConditionProperty := &RuleConditionProperty{
//   	HostHeader: jsii.String("hostHeader"),
//   	HostHeaderWildcard: jsii.String("hostHeaderWildcard"),
//   	PathExact: jsii.String("pathExact"),
//   	PathPrefix: jsii.String("pathPrefix"),
//   	QueryStringEquals: &QueryStringKeyValuePairProperty{
//   		Key: jsii.String("key"),
//   		Value: jsii.String("value"),
//   	},
//   	QueryStringExists: jsii.String("queryStringExists"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-rulecondition.html
//
type CfnLinkRoutingRulePropsMixin_RuleConditionProperty struct {
	// Exact host match — RFC 3986 unreserved characters.
	//
	// Mutually exclusive with HostHeaderWildcard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-rulecondition.html#cfn-rtbfabric-linkroutingrule-rulecondition-hostheader
	//
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Wildcard host pattern (e.g., *.example.com) — RFC 3986 unreserved characters plus *. Mutually exclusive with HostHeader.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-rulecondition.html#cfn-rtbfabric-linkroutingrule-rulecondition-hostheaderwildcard
	//
	HostHeaderWildcard *string `field:"optional" json:"hostHeaderWildcard" yaml:"hostHeaderWildcard"`
	// Exact path match — must start with /.
	//
	// Mutually exclusive with PathPrefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-rulecondition.html#cfn-rtbfabric-linkroutingrule-rulecondition-pathexact
	//
	PathExact *string `field:"optional" json:"pathExact" yaml:"pathExact"`
	// Path prefix matching — strict starts-with, must start with /.
	//
	// Mutually exclusive with PathExact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-rulecondition.html#cfn-rtbfabric-linkroutingrule-rulecondition-pathprefix
	//
	PathPrefix *string `field:"optional" json:"pathPrefix" yaml:"pathPrefix"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-rulecondition.html#cfn-rtbfabric-linkroutingrule-rulecondition-querystringequals
	//
	QueryStringEquals interface{} `field:"optional" json:"queryStringEquals" yaml:"queryStringEquals"`
	// Query string key presence check (any value accepted).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-rulecondition.html#cfn-rtbfabric-linkroutingrule-rulecondition-querystringexists
	//
	QueryStringExists *string `field:"optional" json:"queryStringExists" yaml:"queryStringExists"`
}

