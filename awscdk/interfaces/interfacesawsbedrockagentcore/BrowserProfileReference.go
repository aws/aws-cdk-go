package interfacesawsbedrockagentcore


// A reference to a BrowserProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   browserProfileReference := &BrowserProfileReference{
//   	ProfileArn: jsii.String("profileArn"),
//   	ProfileId: jsii.String("profileId"),
//   }
//
type BrowserProfileReference struct {
	// The ARN of the BrowserProfile resource.
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
	// The ProfileId of the BrowserProfile resource.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
}

