package awspipes


// These are custom parameter to be used when the target is an API Gateway REST APIs or EventBridge ApiDestinations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetHttpParametersProperty := &PipeTargetHttpParametersProperty{
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
type CfnPipe_PipeTargetHttpParametersProperty struct {
	// The headers that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
	HeaderParameters interface{} `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// The path parameter values to be used to populate API Gateway REST API or EventBridge ApiDestination path wildcards ("*").
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// The query string keys/values that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
	QueryStringParameters interface{} `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

