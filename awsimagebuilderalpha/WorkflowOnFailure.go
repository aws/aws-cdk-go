package awsimagebuilderalpha


// The action to take if the workflow fails.
//
// Example:
//   workflow := imagebuilder.NewWorkflow(this, jsii.String("MyWorkflow"), &WorkflowProps{
//   	WorkflowType: imagebuilder.WorkflowType_BUILD,
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
//   				"name": jsii.String("ExecuteComponents"),
//   				"action": imagebuilder.WorkflowAction_EXECUTE_COMPONENTS,
//   				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
//   				"inputs": map[string]*string{
//   					"instanceId": jsii.String("i-123"),
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
type WorkflowOnFailure string

const (
	// Fails the image build if the workflow fails.
	// Experimental.
	WorkflowOnFailure_ABORT WorkflowOnFailure = "ABORT"
	// Continues with the image build if the workflow fails.
	// Experimental.
	WorkflowOnFailure_CONTINUE WorkflowOnFailure = "CONTINUE"
)

