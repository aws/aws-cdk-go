package interfacesawsbedrockagentcore


// A reference to a BrowserCustom resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   browserCustomReference := &BrowserCustomReference{
//   	BrowserId: jsii.String("browserId"),
//   }
//
type BrowserCustomReference struct {
	// The BrowserId of the BrowserCustom resource.
	BrowserId *string `field:"required" json:"browserId" yaml:"browserId"`
}

