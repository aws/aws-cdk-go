package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var model model
//
//   methodResponse := &methodResponse{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	responseModels: map[string]iModel{
//   		"responseModelsKey": model,
//   	},
//   	responseParameters: map[string]*bool{
//   		"responseParametersKey": jsii.Boolean(false),
//   	},
//   }
//
type MethodResponse struct {
	// The method response's status code, which you map to an IntegrationResponse.
	//
	// Required.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The resources used for the response's content type.
	//
	// Specify response models as
	// key-value pairs (string-to-string maps), with a content type as the key and a Model
	// resource name as the value.
	ResponseModels *map[string]IModel `field:"optional" json:"responseModels" yaml:"responseModels"`
	// Response parameters that API Gateway sends to the client that called a method.
	//
	// Specify response parameters as key-value pairs (string-to-Boolean maps), with
	// a destination as the key and a Boolean as the value. Specify the destination
	// using the following pattern: method.response.header.name, where the name is a
	// valid, unique header name. The Boolean specifies whether a parameter is required.
	ResponseParameters *map[string]*bool `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

