package awsapigatewayv2


// Attributes for importing an HttpApi into the CDK.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApiAttributes := &HttpApiAttributes{
//   	HttpApiId: jsii.String("httpApiId"),
//
//   	// the properties below are optional
//   	ApiEndpoint: jsii.String("apiEndpoint"),
//   }
//
type HttpApiAttributes struct {
	// The identifier of the HttpApi.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#aws-resource-apigatewayv2-api-return-values
	//
	HttpApiId *string `field:"required" json:"httpApiId" yaml:"httpApiId"`
	// The endpoint URL of the HttpApi.
	// Default: - throws an error if apiEndpoint is accessed.
	//
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
}

