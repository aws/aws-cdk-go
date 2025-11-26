package previewawstransfermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWorkflowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var copyStepDetails interface{}
//   var customStepDetails interface{}
//   var deleteStepDetails interface{}
//   var tagStepDetails interface{}
//
//   cfnWorkflowMixinProps := &CfnWorkflowMixinProps{
//   	Description: jsii.String("description"),
//   	OnExceptionSteps: []interface{}{
//   		&WorkflowStepProperty{
//   			CopyStepDetails: copyStepDetails,
//   			CustomStepDetails: customStepDetails,
//   			DecryptStepDetails: &DecryptStepDetailsProperty{
//   				DestinationFileLocation: &InputFileLocationProperty{
//   					EfsFileLocation: &EfsInputFileLocationProperty{
//   						FileSystemId: jsii.String("fileSystemId"),
//   						Path: jsii.String("path"),
//   					},
//   					S3FileLocation: &S3InputFileLocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				OverwriteExisting: jsii.String("overwriteExisting"),
//   				SourceFileLocation: jsii.String("sourceFileLocation"),
//   				Type: jsii.String("type"),
//   			},
//   			DeleteStepDetails: deleteStepDetails,
//   			TagStepDetails: tagStepDetails,
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Steps: []interface{}{
//   		&WorkflowStepProperty{
//   			CopyStepDetails: copyStepDetails,
//   			CustomStepDetails: customStepDetails,
//   			DecryptStepDetails: &DecryptStepDetailsProperty{
//   				DestinationFileLocation: &InputFileLocationProperty{
//   					EfsFileLocation: &EfsInputFileLocationProperty{
//   						FileSystemId: jsii.String("fileSystemId"),
//   						Path: jsii.String("path"),
//   					},
//   					S3FileLocation: &S3InputFileLocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				OverwriteExisting: jsii.String("overwriteExisting"),
//   				SourceFileLocation: jsii.String("sourceFileLocation"),
//   				Type: jsii.String("type"),
//   			},
//   			DeleteStepDetails: deleteStepDetails,
//   			TagStepDetails: tagStepDetails,
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html
//
type CfnWorkflowMixinProps struct {
	// Specifies the text description for the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html#cfn-transfer-workflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the steps (actions) to take if errors are encountered during execution of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html#cfn-transfer-workflow-onexceptionsteps
	//
	OnExceptionSteps interface{} `field:"optional" json:"onExceptionSteps" yaml:"onExceptionSteps"`
	// Specifies the details for the steps that are in the specified workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html#cfn-transfer-workflow-steps
	//
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
	// Key-value pairs that can be used to group and search for workflows.
	//
	// Tags are metadata attached to workflows for any purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html#cfn-transfer-workflow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

