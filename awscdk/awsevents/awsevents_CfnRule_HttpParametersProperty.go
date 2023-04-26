package awsevents


// These are custom parameter to be used when the target is an API Gateway APIs or EventBridge ApiDestinations.
//
// In the latter case, these are merged with any InvocationParameters specified on the Connection, with any values from the Connection taking precedence.
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
//   	PathParameterValues: []*string{
//   		jsii.String("pathParameterValues"),
//   	},
//   	QueryStringParameters: map[string]*string{
//   		"queryStringParametersKey": jsii.String("queryStringParameters"),
//   	},
//   }
//
type CfnRule_HttpParametersProperty struct {
	// The headers that need to be sent as part of request invoking the API Gateway API or EventBridge ApiDestination.
	HeaderParameters interface{} `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// The path parameter values to be used to populate API Gateway API or EventBridge ApiDestination path wildcards ("*").
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// The query string keys/values that need to be sent as part of request invoking the API Gateway API or EventBridge ApiDestination.
	QueryStringParameters interface{} `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

