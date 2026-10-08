package interfacesawsdirectoryservice


// A reference to a MicrosoftAD resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   microsoftADReference := &MicrosoftADReference{
//   	DirectoryId: jsii.String("directoryId"),
//   }
//
type MicrosoftADReference struct {
	// The DirectoryId of the MicrosoftAD resource.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
}

