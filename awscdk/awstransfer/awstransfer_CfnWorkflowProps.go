package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkflow`.
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
//   cfnWorkflowProps := &cfnWorkflowProps{
//   	steps: []interface{}{
//   		&workflowStepProperty{
//   			copyStepDetails: copyStepDetails,
//   			customStepDetails: customStepDetails,
//   			decryptStepDetails: &decryptStepDetailsProperty{
//   				destinationFileLocation: &inputFileLocationProperty{
//   					efsFileLocation: &efsInputFileLocationProperty{
//   						fileSystemId: jsii.String("fileSystemId"),
//   						path: jsii.String("path"),
//   					},
//   					s3FileLocation: &s3InputFileLocationProperty{
//   						bucket: jsii.String("bucket"),
//   						key: jsii.String("key"),
//   					},
//   				},
//   				name: jsii.String("name"),
//   				overwriteExisting: jsii.String("overwriteExisting"),
//   				sourceFileLocation: jsii.String("sourceFileLocation"),
//   				type: jsii.String("type"),
//   			},
//   			deleteStepDetails: deleteStepDetails,
//   			tagStepDetails: tagStepDetails,
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	onExceptionSteps: []interface{}{
//   		&workflowStepProperty{
//   			copyStepDetails: copyStepDetails,
//   			customStepDetails: customStepDetails,
//   			decryptStepDetails: &decryptStepDetailsProperty{
//   				destinationFileLocation: &inputFileLocationProperty{
//   					efsFileLocation: &efsInputFileLocationProperty{
//   						fileSystemId: jsii.String("fileSystemId"),
//   						path: jsii.String("path"),
//   					},
//   					s3FileLocation: &s3InputFileLocationProperty{
//   						bucket: jsii.String("bucket"),
//   						key: jsii.String("key"),
//   					},
//   				},
//   				name: jsii.String("name"),
//   				overwriteExisting: jsii.String("overwriteExisting"),
//   				sourceFileLocation: jsii.String("sourceFileLocation"),
//   				type: jsii.String("type"),
//   			},
//   			deleteStepDetails: deleteStepDetails,
//   			tagStepDetails: tagStepDetails,
//   			type: jsii.String("type"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWorkflowProps struct {
	// Specifies the details for the steps that are in the specified workflow.
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// Specifies the text description for the workflow.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the steps (actions) to take if errors are encountered during execution of the workflow.
	OnExceptionSteps interface{} `field:"optional" json:"onExceptionSteps" yaml:"onExceptionSteps"`
	// Key-value pairs that can be used to group and search for workflows.
	//
	// Tags are metadata attached to workflows for any purpose.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

