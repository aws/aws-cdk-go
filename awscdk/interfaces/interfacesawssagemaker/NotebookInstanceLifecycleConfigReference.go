package interfacesawssagemaker


// A reference to a NotebookInstanceLifecycleConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notebookInstanceLifecycleConfigReference := &NotebookInstanceLifecycleConfigReference{
//   	NotebookInstanceLifecycleConfigId: jsii.String("notebookInstanceLifecycleConfigId"),
//   	NotebookInstanceLifecycleConfigName: jsii.String("notebookInstanceLifecycleConfigName"),
//   }
//
type NotebookInstanceLifecycleConfigReference struct {
	// The Id of the NotebookInstanceLifecycleConfig resource.
	NotebookInstanceLifecycleConfigId *string `field:"required" json:"notebookInstanceLifecycleConfigId" yaml:"notebookInstanceLifecycleConfigId"`
	// The NotebookInstanceLifecycleConfigName of the NotebookInstanceLifecycleConfig resource.
	NotebookInstanceLifecycleConfigName *string `field:"required" json:"notebookInstanceLifecycleConfigName" yaml:"notebookInstanceLifecycleConfigName"`
}

