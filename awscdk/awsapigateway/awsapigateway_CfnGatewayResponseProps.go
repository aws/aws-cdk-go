package awsapigateway


// Properties for defining a `CfnGatewayResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayResponseProps := &cfnGatewayResponseProps{
//   	responseType: jsii.String("responseType"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	responseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	responseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	statusCode: jsii.String("statusCode"),
//   }
//
type CfnGatewayResponseProps struct {
	// The response type.
	//
	// For valid values, see [GatewayResponse](https://docs.aws.amazon.com/apigateway/api-reference/resource/gateway-response/) in the *API Gateway API Reference* .
	ResponseType *string `field:"required" json:"responseType" yaml:"responseType"`
	// The identifier of the API.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The response parameters (paths, query strings, and headers) for the response.
	//
	// Duplicates not allowed.
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The response templates for the response.
	//
	// Duplicates not allowed.
	ResponseTemplates interface{} `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// The HTTP status code for the response.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

