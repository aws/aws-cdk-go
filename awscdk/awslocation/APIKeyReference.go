package awslocation


// A reference to a APIKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPIKeyReference := &APIKeyReference{
//   	ApiKeyArn: jsii.String("apiKeyArn"),
//   	KeyName: jsii.String("keyName"),
//   }
//
type APIKeyReference struct {
	// The ARN of the APIKey resource.
	ApiKeyArn *string `field:"required" json:"apiKeyArn" yaml:"apiKeyArn"`
	// The KeyName of the APIKey resource.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
}

