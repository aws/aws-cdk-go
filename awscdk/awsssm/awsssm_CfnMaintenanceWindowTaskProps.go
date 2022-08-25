package awsssm


// Properties for defining a `CfnMaintenanceWindowTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var taskParameters interface{}
//
//   cfnMaintenanceWindowTaskProps := &cfnMaintenanceWindowTaskProps{
//   	priority: jsii.Number(123),
//   	taskArn: jsii.String("taskArn"),
//   	taskType: jsii.String("taskType"),
//   	windowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	cutoffBehavior: jsii.String("cutoffBehavior"),
//   	description: jsii.String("description"),
//   	loggingInfo: &loggingInfoProperty{
//   		region: jsii.String("region"),
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		s3Prefix: jsii.String("s3Prefix"),
//   	},
//   	maxConcurrency: jsii.String("maxConcurrency"),
//   	maxErrors: jsii.String("maxErrors"),
//   	name: jsii.String("name"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	taskInvocationParameters: &taskInvocationParametersProperty{
//   		maintenanceWindowAutomationParameters: &maintenanceWindowAutomationParametersProperty{
//   			documentVersion: jsii.String("documentVersion"),
//   			parameters: parameters,
//   		},
//   		maintenanceWindowLambdaParameters: &maintenanceWindowLambdaParametersProperty{
//   			clientContext: jsii.String("clientContext"),
//   			payload: jsii.String("payload"),
//   			qualifier: jsii.String("qualifier"),
//   		},
//   		maintenanceWindowRunCommandParameters: &maintenanceWindowRunCommandParametersProperty{
//   			cloudWatchOutputConfig: &cloudWatchOutputConfigProperty{
//   				cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				cloudWatchOutputEnabled: jsii.Boolean(false),
//   			},
//   			comment: jsii.String("comment"),
//   			documentHash: jsii.String("documentHash"),
//   			documentHashType: jsii.String("documentHashType"),
//   			documentVersion: jsii.String("documentVersion"),
//   			notificationConfig: &notificationConfigProperty{
//   				notificationArn: jsii.String("notificationArn"),
//
//   				// the properties below are optional
//   				notificationEvents: []*string{
//   					jsii.String("notificationEvents"),
//   				},
//   				notificationType: jsii.String("notificationType"),
//   			},
//   			outputS3BucketName: jsii.String("outputS3BucketName"),
//   			outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			parameters: parameters,
//   			serviceRoleArn: jsii.String("serviceRoleArn"),
//   			timeoutSeconds: jsii.Number(123),
//   		},
//   		maintenanceWindowStepFunctionsParameters: &maintenanceWindowStepFunctionsParametersProperty{
//   			input: jsii.String("input"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	taskParameters: taskParameters,
//   }
//
type CfnMaintenanceWindowTaskProps struct {
	// The priority of the task in the maintenance window.
	//
	// The lower the number, the higher the priority. Tasks that have the same priority are scheduled in parallel.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The resource that the task uses during execution.
	//
	// For `RUN_COMMAND` and `AUTOMATION` task types, `TaskArn` is the SSM document name or Amazon Resource Name (ARN).
	//
	// For `LAMBDA` tasks, `TaskArn` is the function name or ARN.
	//
	// For `STEP_FUNCTIONS` tasks, `TaskArn` is the state machine ARN.
	TaskArn *string `field:"required" json:"taskArn" yaml:"taskArn"`
	// The type of task.
	//
	// Valid values: `RUN_COMMAND` , `AUTOMATION` , `LAMBDA` , `STEP_FUNCTIONS` .
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// The ID of the maintenance window where the task is registered.
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// The specification for whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.
	CutoffBehavior *string `field:"optional" json:"cutoffBehavior" yaml:"cutoffBehavior"`
	// A description of the task.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about an Amazon S3 bucket to write task-level logs to.
	//
	// > `LoggingInfo` has been deprecated. To specify an Amazon S3 bucket to contain logs, instead use the `OutputS3BucketName` and `OutputS3KeyPrefix` options in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [AWS Systems Manager MaintenanceWindowTask TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) .
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The maximum number of targets this task can be run for, in parallel.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	MaxConcurrency *string `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	MaxErrors *string `field:"optional" json:"maxErrors" yaml:"maxErrors"`
	// The task name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// The targets, either instances or window target IDs.
	//
	// - Specify instances using `Key=InstanceIds,Values= *instanceid1* , *instanceid2*` .
	// - Specify window target IDs using `Key=WindowTargetIds,Values= *window-target-id-1* , *window-target-id-2*` .
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The parameters to pass to the task when it runs.
	//
	// Populate only the fields that match the task type. All other fields should be empty.
	//
	// > When you update a maintenance window task that has options specified in `TaskInvocationParameters` , you must provide again all the `TaskInvocationParameters` values that you want to retain. The values you do not specify again are removed. For example, suppose that when you registered a Run Command task, you specified `TaskInvocationParameters` values for `Comment` , `NotificationConfig` , and `OutputS3BucketName` . If you update the maintenance window task and specify only a different `OutputS3BucketName` value, the values for `Comment` and `NotificationConfig` are removed.
	TaskInvocationParameters interface{} `field:"optional" json:"taskInvocationParameters" yaml:"taskInvocationParameters"`
	// The parameters to pass to the task when it runs.
	//
	// > `TaskParameters` has been deprecated. To specify parameters to pass to a task when it runs, instead use the `Parameters` option in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [MaintenanceWindowTaskInvocationParameters](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_MaintenanceWindowTaskInvocationParameters.html) .
	TaskParameters interface{} `field:"optional" json:"taskParameters" yaml:"taskParameters"`
}

