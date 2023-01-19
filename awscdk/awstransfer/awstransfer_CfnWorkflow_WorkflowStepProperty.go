package awstransfer


// The basic building block of a workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var copyStepDetails interface{}
//   var customStepDetails interface{}
//   var deleteStepDetails interface{}
//   var tagStepDetails interface{}
//
//   workflowStepProperty := &workflowStepProperty{
//   	copyStepDetails: copyStepDetails,
//   	customStepDetails: customStepDetails,
//   	deleteStepDetails: deleteStepDetails,
//   	tagStepDetails: tagStepDetails,
//   	type: jsii.String("type"),
//   }
//
type CfnWorkflow_WorkflowStepProperty struct {
	// Details for a step that performs a file copy.
	//
	// Consists of the following values:
	//
	// - A description
	// - An Amazon S3 location for the destination of the file copy.
	// - A flag that indicates whether to overwrite an existing file of the same name. The default is `FALSE` .
	CopyStepDetails interface{} `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// Details for a step that invokes an AWS Lambda function.
	//
	// Consists of the Lambda function's name, target, and timeout (in seconds).
	CustomStepDetails interface{} `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// Details for a step that deletes the file.
	DeleteStepDetails interface{} `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// Details for a step that creates one or more tags.
	//
	// You specify one or more tags. Each tag contains a key-value pair.
	TagStepDetails interface{} `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
	// Currently, the following step types are supported.
	//
	// - *`COPY`* - Copy the file to another location.
	// - *`CUSTOM`* - Perform a custom step with an AWS Lambda function target.
	// - *`DECRYPT`* - Decrypt a file that was encrypted before it was uploaded.
	// - *`DELETE`* - Delete the file.
	// - *`TAG`* - Add a tag to the file.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

