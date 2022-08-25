package awsapigateway


// `MethodResponse` is a property of the [AWS::ApiGateway::Method](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html) resource that defines the responses that can be sent to the client that calls a method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   methodResponseProperty := &methodResponseProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	responseModels: map[string]*string{
//   		"responseModelsKey": jsii.String("responseModels"),
//   	},
//   	responseParameters: map[string]interface{}{
//   		"responseParametersKey": jsii.Boolean(false),
//   	},
//   }
//
type CfnMethod_MethodResponseProperty struct {
	// The method response's status code, which you map to an [IntegrationResponse](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html) .
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The resources used for the response's content type.
	//
	// Specify response models as key-value pairs (string-to-string maps), with a content type as the key and a [Model](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html) resource name as the value.
	ResponseModels interface{} `field:"optional" json:"responseModels" yaml:"responseModels"`
	// Response parameters that API Gateway sends to the client that called a method.
	//
	// Specify response parameters as key-value pairs (string-to-Boolean maps), with a destination as the key and a Boolean as the value. Specify the destination using the following pattern: `method.response.header. *name*` , where *name* is a valid, unique header name. The Boolean specifies whether a parameter is required.
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

