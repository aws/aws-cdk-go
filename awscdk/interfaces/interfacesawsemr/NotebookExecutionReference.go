package interfacesawsemr


// A reference to a NotebookExecution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notebookExecutionReference := &NotebookExecutionReference{
//   	NotebookExecutionArn: jsii.String("notebookExecutionArn"),
//   }
//
type NotebookExecutionReference struct {
	// The Arn of the NotebookExecution resource.
	NotebookExecutionArn *string `field:"required" json:"notebookExecutionArn" yaml:"notebookExecutionArn"`
}

