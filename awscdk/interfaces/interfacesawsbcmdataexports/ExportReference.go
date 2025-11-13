package interfacesawsbcmdataexports


// A reference to a Export resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exportReference := &ExportReference{
//   	ExportArn: jsii.String("exportArn"),
//   }
//
type ExportReference struct {
	// The ExportArn of the Export resource.
	ExportArn *string `field:"required" json:"exportArn" yaml:"exportArn"`
}

