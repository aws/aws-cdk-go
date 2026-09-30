package interfacesawssagemaker


// A reference to a NotebookInstanceLifecycleConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notebookInstanceLifecycleConfigReference := &NotebookInstanceLifecycleConfigReference{
//   	NotebookInstanceLifecycleConfigArn: jsii.String("notebookInstanceLifecycleConfigArn"),
//   }
//
type NotebookInstanceLifecycleConfigReference struct {
	// The NotebookInstanceLifecycleConfigArn of the NotebookInstanceLifecycleConfig resource.
	NotebookInstanceLifecycleConfigArn *string `field:"required" json:"notebookInstanceLifecycleConfigArn" yaml:"notebookInstanceLifecycleConfigArn"`
}

