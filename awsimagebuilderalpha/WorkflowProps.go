package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Workflow resource.
//
// Example:
//   workflow := imagebuilder.NewWorkflow(this, jsii.String("EncryptedWorkflow"), &WorkflowProps{
//   	WorkflowType: imagebuilder.WorkflowType_BUILD,
//   	KmsKey: kms.NewKey(this, jsii.String("WorkflowKey")),
//   	Data: imagebuilder.WorkflowData_FromJsonObject(map[string]interface{}{
//   		"schemaVersion": imagebuilder.WorkflowSchemaVersion_V1_0,
//   		"steps": []interface{}{
//   			map[string]interface{}{
//   				"name": jsii.String("LaunchBuildInstance"),
//   				"action": imagebuilder.WorkflowAction_LAUNCH_INSTANCE,
//   				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
//   				"inputs": map[string]*string{
//   					"waitFor": jsii.String("ssmAgent"),
//   				},
//   			},
//   			map[string]interface{}{
//   				"name": jsii.String("CreateImage"),
//   				"action": imagebuilder.WorkflowAction_CREATE_IMAGE,
//   				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
//   				"inputs": map[string]*string{
//   					"instanceId": jsii.String("i-123"),
//   				},
//   			},
//   			map[string]interface{}{
//   				"name": jsii.String("TerminateInstance"),
//   				"action": imagebuilder.WorkflowAction_TERMINATE_INSTANCE,
//   				"onFailure": imagebuilder.WorkflowOnFailure_CONTINUE,
//   				"inputs": map[string]*string{
//   					"instanceId": jsii.String("i-123"),
//   				},
//   			},
//   		},
//   		"outputs": []interface{}{
//   			map[string]*string{
//   				"name": jsii.String("ImageId"),
//   				"value": jsii.String("$.stepOutputs.CreateImage.imageId"),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type WorkflowProps struct {
	// The workflow document content that defines the image creation process.
	// Experimental.
	Data WorkflowData `field:"required" json:"data" yaml:"data"`
	// The phase in the image build process for which the workflow resource is responsible.
	// Experimental.
	WorkflowType WorkflowType `field:"required" json:"workflowType" yaml:"workflowType"`
	// The change description of the workflow.
	//
	// Describes what change has been made in this version of the workflow, or
	// what makes this version different from other versions.
	// Default: None.
	//
	// Experimental.
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// The description of the workflow.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key used to encrypt this workflow.
	// Default: - an Image Builder owned key will be used to encrypt the workflow.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The tags to apply to the workflow.
	// Default: None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workflow.
	// Default: - a name is generated.
	//
	// Experimental.
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
	// The version of the workflow.
	// Default: 1.0.0
	//
	// Experimental.
	WorkflowVersion *string `field:"optional" json:"workflowVersion" yaml:"workflowVersion"`
}

