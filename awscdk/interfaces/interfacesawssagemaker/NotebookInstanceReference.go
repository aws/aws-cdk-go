package interfacesawssagemaker


// A reference to a NotebookInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notebookInstanceReference := &NotebookInstanceReference{
//   	NotebookInstanceArn: jsii.String("notebookInstanceArn"),
//   }
//
type NotebookInstanceReference struct {
	// The NotebookInstanceArn of the NotebookInstance resource.
	NotebookInstanceArn *string `field:"required" json:"notebookInstanceArn" yaml:"notebookInstanceArn"`
}

