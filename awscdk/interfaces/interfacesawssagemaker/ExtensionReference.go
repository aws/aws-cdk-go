package interfacesawssagemaker


// A reference to a Extension resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extensionReference := &ExtensionReference{
//   	ExtensionArn: jsii.String("extensionArn"),
//   }
//
type ExtensionReference struct {
	// The ExtensionArn of the Extension resource.
	ExtensionArn *string `field:"required" json:"extensionArn" yaml:"extensionArn"`
}

