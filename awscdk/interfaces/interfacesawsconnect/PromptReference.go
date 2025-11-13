package interfacesawsconnect


// A reference to a Prompt resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptReference := &PromptReference{
//   	PromptArn: jsii.String("promptArn"),
//   }
//
type PromptReference struct {
	// The PromptArn of the Prompt resource.
	PromptArn *string `field:"required" json:"promptArn" yaml:"promptArn"`
}

