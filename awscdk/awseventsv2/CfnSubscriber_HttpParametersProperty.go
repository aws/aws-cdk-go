package awseventsv2


// Parameters for invoking an HTTP endpoint target, such as an Amazon API Gateway endpoint or an EventBridge API destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpParametersProperty := &HttpParametersProperty{
//   	HeaderParameters: map[string]*string{
//   		"headerParametersKey": jsii.String("headerParameters"),
//   	},
//   	InvocationTimeoutSeconds: jsii.String("invocationTimeoutSeconds"),
//   	PathParameterValues: []*string{
//   		jsii.String("pathParameterValues"),
//   	},
//   	QueryStringParameters: map[string]*string{
//   		"queryStringParametersKey": jsii.String("queryStringParameters"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-httpparameters.html
//
type CfnSubscriber_HttpParametersProperty struct {
	// A map of HTTP header names to values.
	//
	// A header name or value may also be a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-httpparameters.html#cfn-eventsv2-subscriber-httpparameters-headerparameters
	//
	HeaderParameters interface{} `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// The timeout in seconds for each invocation of the target, written as a string.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-httpparameters.html#cfn-eventsv2-subscriber-httpparameters-invocationtimeoutseconds
	//
	InvocationTimeoutSeconds *string `field:"optional" json:"invocationTimeoutSeconds" yaml:"invocationTimeoutSeconds"`
	// Values for the path parameters (wildcards) in the target URL, in order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-httpparameters.html#cfn-eventsv2-subscriber-httpparameters-pathparametervalues
	//
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// A map of query string parameter names to values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-httpparameters.html#cfn-eventsv2-subscriber-httpparameters-querystringparameters
	//
	QueryStringParameters interface{} `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

