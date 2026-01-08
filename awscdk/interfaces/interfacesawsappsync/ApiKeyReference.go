package interfacesawsappsync


// A reference to a ApiKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyReference := &ApiKeyReference{
//   	ApiKeyArn: jsii.String("apiKeyArn"),
//   }
//
type ApiKeyReference struct {
	// The Arn of the ApiKey resource.
	ApiKeyArn *string `field:"required" json:"apiKeyArn" yaml:"apiKeyArn"`
}

