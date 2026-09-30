package interfacesawsappsync


// A reference to a ApiKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyReference := &ApiKeyReference{
//   	ApiId: jsii.String("apiId"),
//   	ApiKeyArn: jsii.String("apiKeyArn"),
//   	ApiKeyId: jsii.String("apiKeyId"),
//   }
//
type ApiKeyReference struct {
	// The ApiId of the ApiKey resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The ARN of the ApiKey resource.
	ApiKeyArn *string `field:"required" json:"apiKeyArn" yaml:"apiKeyArn"`
	// The ApiKeyId of the ApiKey resource.
	ApiKeyId *string `field:"required" json:"apiKeyId" yaml:"apiKeyId"`
}

