package awstransfer


// The basic building block of a workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowStepProperty := &workflowStepProperty{
//   	copyStepDetails: &copyStepDetailsProperty{
//   		destinationFileLocation: &inputFileLocationProperty{
//   			s3FileLocation: &s3InputFileLocationProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//   			},
//   		},
//   		name: jsii.String("name"),
//   		overwriteExisting: jsii.String("overwriteExisting"),
//   		sourceFileLocation: jsii.String("sourceFileLocation"),
//   	},
//   	customStepDetails: &customStepDetailsProperty{
//   		name: jsii.String("name"),
//   		sourceFileLocation: jsii.String("sourceFileLocation"),
//   		target: jsii.String("target"),
//   		timeoutSeconds: jsii.Number(123),
//   	},
//   	deleteStepDetails: &deleteStepDetailsProperty{
//   		name: jsii.String("name"),
//   		sourceFileLocation: jsii.String("sourceFileLocation"),
//   	},
//   	tagStepDetails: &tagStepDetailsProperty{
//   		name: jsii.String("name"),
//   		sourceFileLocation: jsii.String("sourceFileLocation"),
//   		tags: []s3TagProperty{
//   			&s3TagProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnWorkflow_WorkflowStepProperty struct {
	// Details for a step that performs a file copy.
	//
	// Consists of the following values:
	//
	// - A description
	// - An S3 location for the destination of the file copy.
	// - A flag that indicates whether or not to overwrite an existing file of the same name. The default is `FALSE` .
	CopyStepDetails interface{} `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// Details for a step that invokes a lambda function.
	//
	// Consists of the lambda function name, target, and timeout (in seconds).
	CustomStepDetails interface{} `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// Details for a step that deletes the file.
	DeleteStepDetails interface{} `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// Details for a step that creates one or more tags.
	//
	// You specify one or more tags: each tag contains a key/value pair.
	TagStepDetails interface{} `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
	// Currently, the following step types are supported.
	//
	// - *COPY* : copy the file to another location
	// - *CUSTOM* : custom step with a lambda target
	// - *DELETE* : delete the file
	// - *TAG* : add a tag to the file.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

