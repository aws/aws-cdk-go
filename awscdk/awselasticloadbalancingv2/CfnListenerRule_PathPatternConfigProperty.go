package awselasticloadbalancingv2


// Information about a path pattern condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathPatternConfigProperty := &PathPatternConfigProperty{
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-pathpatternconfig.html
//
type CfnListenerRule_PathPatternConfigProperty struct {
	// The path patterns to compare against the request URL.
	//
	// The maximum size of each string is 128 characters. The comparison is case sensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If you specify multiple strings, the condition is satisfied if one of them matches the request URL. The path pattern is compared only to the path of the URL, not to its query string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-pathpatternconfig.html#cfn-elasticloadbalancingv2-listenerrule-pathpatternconfig-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

