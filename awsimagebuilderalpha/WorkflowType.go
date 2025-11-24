package awsimagebuilderalpha


// The type of the workflow.
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
type WorkflowType string

const (
	// Indicates the workflow is for building images.
	// Experimental.
	WorkflowType_BUILD WorkflowType = "BUILD"
	// Indicates the workflow is for testing images.
	// Experimental.
	WorkflowType_TEST WorkflowType = "TEST"
	// Indicates the workflow is for distributing images.
	// Experimental.
	WorkflowType_DISTRIBUTION WorkflowType = "DISTRIBUTION"
)

