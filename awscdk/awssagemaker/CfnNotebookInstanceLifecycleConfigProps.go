package awssagemaker


// Properties for defining a `CfnNotebookInstanceLifecycleConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotebookInstanceLifecycleConfigProps := &CfnNotebookInstanceLifecycleConfigProps{
//   	NotebookInstanceLifecycleConfigName: jsii.String("notebookInstanceLifecycleConfigName"),
//   	OnCreate: []interface{}{
//   		&NotebookInstanceLifecycleHookProperty{
//   			Content: jsii.String("content"),
//   		},
//   	},
//   	OnStart: []interface{}{
//   		&NotebookInstanceLifecycleHookProperty{
//   			Content: jsii.String("content"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html
//
type CfnNotebookInstanceLifecycleConfigProps struct {
	// The name of the lifecycle configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecycleconfigname
	//
	NotebookInstanceLifecycleConfigName *string `field:"optional" json:"notebookInstanceLifecycleConfigName" yaml:"notebookInstanceLifecycleConfigName"`
	// A shell script that runs only once, when you create a notebook instance.
	//
	// The shell script must be a base64-encoded string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-oncreate
	//
	OnCreate interface{} `field:"optional" json:"onCreate" yaml:"onCreate"`
	// A shell script that runs every time you start a notebook instance, including when you create the notebook instance.
	//
	// The shell script must be a base64-encoded string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-onstart
	//
	OnStart interface{} `field:"optional" json:"onStart" yaml:"onStart"`
}

