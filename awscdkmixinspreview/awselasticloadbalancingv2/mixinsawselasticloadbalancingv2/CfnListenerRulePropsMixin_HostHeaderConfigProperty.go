package mixinsawselasticloadbalancingv2


// Information about a host header condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hostHeaderConfigProperty := &HostHeaderConfigProperty{
//   	RegexValues: []*string{
//   		jsii.String("regexValues"),
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-hostheaderconfig.html
//
type CfnListenerRulePropsMixin_HostHeaderConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-hostheaderconfig.html#cfn-elasticloadbalancingv2-listenerrule-hostheaderconfig-regexvalues
	//
	RegexValues *[]*string `field:"optional" json:"regexValues" yaml:"regexValues"`
	// The host names.
	//
	// The maximum length of each string is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character). You must include at least one "." character. You can include only alphabetical characters after the final "." character.
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the host name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-hostheaderconfig.html#cfn-elasticloadbalancingv2-listenerrule-hostheaderconfig-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

