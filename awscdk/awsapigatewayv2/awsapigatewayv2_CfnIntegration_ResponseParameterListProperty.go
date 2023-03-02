package awsapigatewayv2


// Specifies a list of response parameters for an HTTP API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseParameterListProperty := &responseParameterListProperty{
//   	responseParameters: []interface{}{
//   		&responseParameterProperty{
//   			destination: jsii.String("destination"),
//   			source: jsii.String("source"),
//   		},
//   	},
//   }
//
type CfnIntegration_ResponseParameterListProperty struct {
	// Supported only for HTTP APIs.
	//
	// You use response parameters to transform the HTTP response from a backend integration before returning the response to clients. Specify a key-value map from a selection key to response parameters. The selection key must be a valid HTTP status code within the range of 200-599. Response parameters are a key-value map. The key must match the pattern `<action>:<header>.<location>` or `overwrite.statuscode` . The action can be `append` , `overwrite` or `remove` . The value can be a static value, or map to response data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) .
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

