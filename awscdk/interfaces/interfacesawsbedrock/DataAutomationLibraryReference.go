package interfacesawsbedrock


// A reference to a DataAutomationLibrary resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataAutomationLibraryReference := &DataAutomationLibraryReference{
//   	LibraryArn: jsii.String("libraryArn"),
//   }
//
type DataAutomationLibraryReference struct {
	// The LibraryArn of the DataAutomationLibrary resource.
	LibraryArn *string `field:"required" json:"libraryArn" yaml:"libraryArn"`
}

