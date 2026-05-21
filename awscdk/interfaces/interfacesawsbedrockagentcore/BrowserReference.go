package interfacesawsbedrockagentcore


// A reference to a Browser resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   browserReference := &BrowserReference{
//   	BrowserArn: jsii.String("browserArn"),
//   }
//
type BrowserReference struct {
	// The BrowserArn of the Browser resource.
	BrowserArn *string `field:"required" json:"browserArn" yaml:"browserArn"`
}

