package awselasticloadbalancingv2


// Information about an HTTP header condition.
//
// There is a set of standard HTTP header fields. You can also define custom HTTP header fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpHeaderConfigProperty := &HttpHeaderConfigProperty{
//   	HttpHeaderName: jsii.String("httpHeaderName"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-httpheaderconfig.html
//
type CfnListenerRule_HttpHeaderConfigProperty struct {
	// The name of the HTTP header field.
	//
	// The maximum size is 40 characters. The header name is case insensitive. The allowed characters are specified by RFC 7230. Wildcards are not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-httpheaderconfig.html#cfn-elasticloadbalancingv2-listenerrule-httpheaderconfig-httpheadername
	//
	HttpHeaderName *string `field:"optional" json:"httpHeaderName" yaml:"httpHeaderName"`
	// The strings to compare against the value of the HTTP header.
	//
	// The maximum size of each string is 128 characters. The comparison strings are case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If the same header appears multiple times in the request, we search them in order until a match is found.
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the value of the HTTP header. To require that all of the strings are a match, create one condition per string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-httpheaderconfig.html#cfn-elasticloadbalancingv2-listenerrule-httpheaderconfig-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

