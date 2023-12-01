package awscdkapigatewayv2alpha


// Attributes for importing an HttpApi into the CDK.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   httpApiAttributes := &HttpApiAttributes{
//   	HttpApiId: jsii.String("httpApiId"),
//
//   	// the properties below are optional
//   	ApiEndpoint: jsii.String("apiEndpoint"),
//   }
//
// Deprecated.
type HttpApiAttributes struct {
	// The identifier of the HttpApi.
	// Deprecated.
	HttpApiId *string `field:"required" json:"httpApiId" yaml:"httpApiId"`
	// The endpoint URL of the HttpApi.
	// Default: - throws an error if apiEndpoint is accessed.
	//
	// Deprecated.
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
}

