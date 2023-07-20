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
//   cfnMaintenanceWindowTaskProps := &CfnMaintenanceWindowTaskProps{
//   	Priority: jsii.Number(123),
//   	TaskArn: jsii.String("taskArn"),
//   	TaskType: jsii.String("taskType"),
//   	WindowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	CutoffBehavior: jsii.String("cutoffBehavior"),
//   	Description: jsii.String("description"),
//   	LoggingInfo: &LoggingInfoProperty{
//   		Region: jsii.String("region"),
//   		S3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		S3Prefix: jsii.String("s3Prefix"),
//   	},
//   	MaxConcurrency: jsii.String("maxConcurrency"),
//   	MaxErrors: jsii.String("maxErrors"),
//   	Name: jsii.String("name"),
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	TaskInvocationParameters: &TaskInvocationParametersProperty{
//   		MaintenanceWindowAutomationParameters: &MaintenanceWindowAutomationParametersProperty{
//   			DocumentVersion: jsii.String("documentVersion"),
//   			Parameters: parameters,
//   		},
//   		MaintenanceWindowLambdaParameters: &MaintenanceWindowLambdaParametersProperty{
//   			ClientContext: jsii.String("clientContext"),
//   			Payload: jsii.String("payload"),
//   			Qualifier: jsii.String("qualifier"),
//   		},
//   		MaintenanceWindowRunCommandParameters: &MaintenanceWindowRunCommandParametersProperty{
//   			CloudWatchOutputConfig: &CloudWatchOutputConfigProperty{
//   				CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				CloudWatchOutputEnabled: jsii.Boolean(false),
//   			},
//   			Comment: jsii.String("comment"),
//   			DocumentHash: jsii.String("documentHash"),
//   			DocumentHashType: jsii.String("documentHashType"),
//   			DocumentVersion: jsii.String("documentVersion"),
//   			NotificationConfig: &NotificationConfigProperty{
//   				NotificationArn: jsii.String("notificationArn"),
//
//   				// the properties below are optional
//   				NotificationEvents: []*string{
//   					jsii.String("notificationEvents"),
//   				},
//   				NotificationType: jsii.String("notificationType"),
//   			},
//   			OutputS3BucketName: jsii.String("outputS3BucketName"),
//   			OutputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			Parameters: parameters,
//   			ServiceRoleArn: jsii.String("serviceRoleArn"),
//   			TimeoutSeconds: jsii.Number(123),
//   		},
//   		MaintenanceWindowStepFunctionsParameters: &MaintenanceWindowStepFunctionsParametersProperty{
//   			Input: jsii.String("input"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TaskParameters: taskParameters,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html
//
type CfnMaintenanceWindowTaskProps struct {
	// The priority of the task in the maintenance window.
	//
	// The lower the number, the higher the priority. Tasks that have the same priority are scheduled in parallel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The resource that the task uses during execution.
	//
	// For `RUN_COMMAND` and `AUTOMATION` task types, `TaskArn` is the SSM document name or Amazon Resource Name (ARN).
	//
	// For `LAMBDA` tasks, `TaskArn` is the function name or ARN.
	//
	// For `STEP_FUNCTIONS` tasks, `TaskArn` is the state machine ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskarn
	//
	TaskArn *string `field:"required" json:"taskArn" yaml:"taskArn"`
	// The type of task.
	//
	// Valid values: `RUN_COMMAND` , `AUTOMATION` , `LAMBDA` , `STEP_FUNCTIONS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-tasktype
	//
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// The ID of the maintenance window where the task is registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-windowid
	//
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// The specification for whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-cutoffbehavior
	//
	CutoffBehavior *string `field:"optional" json:"cutoffBehavior" yaml:"cutoffBehavior"`
	// A description of the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about an Amazon S3 bucket to write Run Command task-level logs to.
	//
	// > `LoggingInfo` has been deprecated. To specify an Amazon S3 bucket to contain logs for Run Command tasks, instead use the `OutputS3BucketName` and `OutputS3KeyPrefix` options in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [AWS ::SSM::MaintenanceWindowTask MaintenanceWindowRunCommandParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-logginginfo
	//
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The maximum number of targets this task can be run for, in parallel.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-maxconcurrency
	//
	MaxConcurrency *string `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-maxerrors
	//
	MaxErrors *string `field:"optional" json:"maxErrors" yaml:"maxErrors"`
	// The task name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-servicerolearn
	//
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// The targets, either instances or window target IDs.
	//
	// - Specify instances using `Key=InstanceIds,Values= *instanceid1* , *instanceid2*` .
	// - Specify window target IDs using `Key=WindowTargetIds,Values= *window-target-id-1* , *window-target-id-2*` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The parameters to pass to the task when it runs.
	//
	// Populate only the fields that match the task type. All other fields should be empty.
	//
	// > When you update a maintenance window task that has options specified in `TaskInvocationParameters` , you must provide again all the `TaskInvocationParameters` values that you want to retain. The values you do not specify again are removed. For example, suppose that when you registered a Run Command task, you specified `TaskInvocationParameters` values for `Comment` , `NotificationConfig` , and `OutputS3BucketName` . If you update the maintenance window task and specify only a different `OutputS3BucketName` value, the values for `Comment` and `NotificationConfig` are removed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters
	//
	TaskInvocationParameters interface{} `field:"optional" json:"taskInvocationParameters" yaml:"taskInvocationParameters"`
	// The parameters to pass to the task when it runs.
	//
	// > `TaskParameters` has been deprecated. To specify parameters to pass to a task when it runs, instead use the `Parameters` option in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [MaintenanceWindowTaskInvocationParameters](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_MaintenanceWindowTaskInvocationParameters.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskparameters
	//
	TaskParameters interface{} `field:"optional" json:"taskParameters" yaml:"taskParameters"`
}

