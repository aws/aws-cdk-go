package interfacesawssagemaker


// A reference to a ModelCardExportJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelCardExportJobReference := &ModelCardExportJobReference{
//   	ModelCardExportJobArn: jsii.String("modelCardExportJobArn"),
//   }
//
type ModelCardExportJobReference struct {
	// The ModelCardExportJobArn of the ModelCardExportJob resource.
	ModelCardExportJobArn *string `field:"required" json:"modelCardExportJobArn" yaml:"modelCardExportJobArn"`
}

