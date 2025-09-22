package awsapigateway


// A reference to a ApiKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyReference := &ApiKeyReference{
//   	ApiKeyId: jsii.String("apiKeyId"),
//   }
//
type ApiKeyReference struct {
	// The APIKeyId of the ApiKey resource.
	ApiKeyId *string `field:"required" json:"apiKeyId" yaml:"apiKeyId"`
}

