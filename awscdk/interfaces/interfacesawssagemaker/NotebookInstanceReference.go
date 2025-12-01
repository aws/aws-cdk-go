package interfacesawssagemaker


// A reference to a NotebookInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notebookInstanceReference := &NotebookInstanceReference{
//   	NotebookInstanceId: jsii.String("notebookInstanceId"),
//   	NotebookInstanceName: jsii.String("notebookInstanceName"),
//   }
//
type NotebookInstanceReference struct {
	// The Id of the NotebookInstance resource.
	NotebookInstanceId *string `field:"required" json:"notebookInstanceId" yaml:"notebookInstanceId"`
	// The NotebookInstanceName of the NotebookInstance resource.
	NotebookInstanceName *string `field:"required" json:"notebookInstanceName" yaml:"notebookInstanceName"`
}

