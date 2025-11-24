package awsimagebuilderalpha


// The action for a step within the workflow document.
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
type WorkflowAction string

const (
	// Applies customizations and configurations to the input AMIs, such as publishing the AMI to SSM Parameter Store, or creating launch template versions with the AMI IDs provided in the input.
	// Experimental.
	WorkflowAction_APPLY_IMAGE_CONFIGURATIONS WorkflowAction = "APPLY_IMAGE_CONFIGURATIONS"
	// The BootstrapInstanceForContainer action runs a service script to bootstrap the instance with minimum requirements to run container workflows.
	// Experimental.
	WorkflowAction_BOOTSTRAP_INSTANCE_FOR_CONTAINER WorkflowAction = "BOOTSTRAP_INSTANCE_FOR_CONTAINER"
	// The CollectImageMetadata action collects additional information about the instance, such as the list of packages and their respective versions.
	// Experimental.
	WorkflowAction_COLLECT_IMAGE_METADATA WorkflowAction = "COLLECT_IMAGE_METADATA"
	// The CollectImageScanFindings action collects findings reported by Amazon Inspector for the provided instance.
	// Experimental.
	WorkflowAction_COLLECT_IMAGE_SCAN_FINDINGS WorkflowAction = "COLLECT_IMAGE_SCAN_FINDINGS"
	// The CreateImage action creates an AMI from a running instance with the ec2:CreateImage API.
	// Experimental.
	WorkflowAction_CREATE_IMAGE WorkflowAction = "CREATE_IMAGE"
	// The DistributeImage action copies an AMI using the image's distribution configuration, or using the distribution settings in the step input.
	// Experimental.
	WorkflowAction_DISTRIBUTE_IMAGE WorkflowAction = "DISTRIBUTE_IMAGE"
	// The ExecuteComponents action runs components that are specified in the recipe for the current image being built.
	// Experimental.
	WorkflowAction_EXECUTE_COMPONENTS WorkflowAction = "EXECUTE_COMPONENTS"
	// The ExecuteStateMachine action executes a the state machine provided and waits for completion as part of the workflow.
	// Experimental.
	WorkflowAction_EXECUTE_STATE_MACHINE WorkflowAction = "EXECUTE_STATE_MACHINE"
	// The LaunchInstance action launches an instance using the settings from your recipe and infrastructure configuration resources.
	// Experimental.
	WorkflowAction_LAUNCH_INSTANCE WorkflowAction = "LAUNCH_INSTANCE"
	// Applies attribute updates to the provided set of distributed images, such as launch permission updates.
	// Experimental.
	WorkflowAction_MODIFY_IMAGE_ATTRIBUTES WorkflowAction = "MODIFY_IMAGE_ATTRIBUTES"
	// The RunCommand action runs a command document against the provided instance.
	// Experimental.
	WorkflowAction_RUN_COMMAND WorkflowAction = "RUN_COMMAND"
	// The RegisterImage action creates an AMI from a set of snapshots with the ec2:RegisterImage API.
	// Experimental.
	WorkflowAction_REGISTER_IMAGE WorkflowAction = "REGISTER_IMAGE"
	// The RunSysprep action runs the Sysprep document on the provided Windows instance.
	// Experimental.
	WorkflowAction_RUN_SYS_PREP WorkflowAction = "RUN_SYS_PREP"
	// The SanitizeInstance action runs a recommended sanitization script on Linux instances.
	// Experimental.
	WorkflowAction_SANITIZE_INSTANCE WorkflowAction = "SANITIZE_INSTANCE"
	// The TerminateInstance action terminates the provided instance.
	// Experimental.
	WorkflowAction_TERMINATE_INSTANCE WorkflowAction = "TERMINATE_INSTANCE"
	// The WaitForAction action pauses the workflow and waits to receive an external signal from the imagebuilder:SendWorkflowStepAction API.
	// Experimental.
	WorkflowAction_WAIT_FOR_ACTION WorkflowAction = "WAIT_FOR_ACTION"
	// The WaitForSSMAgent action waits for the given instance to have connectivity with SSM before proceeding.
	// Experimental.
	WorkflowAction_WAIT_FOR_SSM_AGENT WorkflowAction = "WAIT_FOR_SSM_AGENT"
)

