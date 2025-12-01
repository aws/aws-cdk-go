package interfacesawsbedrockagentcore


// A reference to a BrowserCustom resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   browserCustomReference := &BrowserCustomReference{
//   	BrowserArn: jsii.String("browserArn"),
//   	BrowserId: jsii.String("browserId"),
//   }
//
type BrowserCustomReference struct {
	// The ARN of the BrowserCustom resource.
	BrowserArn *string `field:"required" json:"browserArn" yaml:"browserArn"`
	// The BrowserId of the BrowserCustom resource.
	BrowserId *string `field:"required" json:"browserId" yaml:"browserId"`
}

