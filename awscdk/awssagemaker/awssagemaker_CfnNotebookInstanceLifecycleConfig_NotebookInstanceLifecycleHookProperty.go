package awssagemaker


// Specifies the notebook instance lifecycle configuration script.
//
// Each lifecycle configuration script has a limit of 16384 characters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notebookInstanceLifecycleHookProperty := &notebookInstanceLifecycleHookProperty{
//   	content: jsii.String("content"),
//   }
//
type CfnNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHookProperty struct {
	// A base64-encoded string that contains a shell script for a notebook instance lifecycle configuration.
	Content *string `field:"optional" json:"content" yaml:"content"`
}

