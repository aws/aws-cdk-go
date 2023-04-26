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
// Experimental.
type HttpApiAttributes struct {
	// The identifier of the HttpApi.
	// Experimental.
	HttpApiId *string `field:"required" json:"httpApiId" yaml:"httpApiId"`
	// The endpoint URL of the HttpApi.
	// Experimental.
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
}

