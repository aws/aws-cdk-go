package awssagemaker


// Properties for defining a `CfnNotebookInstanceLifecycleConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotebookInstanceLifecycleConfigProps := &cfnNotebookInstanceLifecycleConfigProps{
//   	notebookInstanceLifecycleConfigName: jsii.String("notebookInstanceLifecycleConfigName"),
//   	onCreate: []interface{}{
//   		&notebookInstanceLifecycleHookProperty{
//   			content: jsii.String("content"),
//   		},
//   	},
//   	onStart: []interface{}{
//   		&notebookInstanceLifecycleHookProperty{
//   			content: jsii.String("content"),
//   		},
//   	},
//   }
//
type CfnNotebookInstanceLifecycleConfigProps struct {
	// The name of the lifecycle configuration.
	NotebookInstanceLifecycleConfigName *string `field:"optional" json:"notebookInstanceLifecycleConfigName" yaml:"notebookInstanceLifecycleConfigName"`
	// A shell script that runs only once, when you create a notebook instance.
	//
	// The shell script must be a base64-encoded string.
	OnCreate interface{} `field:"optional" json:"onCreate" yaml:"onCreate"`
	// A shell script that runs every time you start a notebook instance, including when you create the notebook instance.
	//
	// The shell script must be a base64-encoded string.
	OnStart interface{} `field:"optional" json:"onStart" yaml:"onStart"`
}

