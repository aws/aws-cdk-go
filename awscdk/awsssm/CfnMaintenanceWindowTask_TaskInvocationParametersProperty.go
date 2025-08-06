package awsssm


// The `TaskInvocationParameters` property type specifies the task execution parameters for a maintenance window task in AWS Systems Manager .
//
// `TaskInvocationParameters` is a property of the [AWS::SSM::MaintenanceWindowTask](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   taskInvocationParametersProperty := &TaskInvocationParametersProperty{
//   	MaintenanceWindowAutomationParameters: &MaintenanceWindowAutomationParametersProperty{
//   		DocumentVersion: jsii.String("documentVersion"),
//   		Parameters: parameters,
//   	},
//   	MaintenanceWindowLambdaParameters: &MaintenanceWindowLambdaParametersProperty{
//   		ClientContext: jsii.String("clientContext"),
//   		Payload: jsii.String("payload"),
//   		Qualifier: jsii.String("qualifier"),
//   	},
//   	MaintenanceWindowRunCommandParameters: &MaintenanceWindowRunCommandParametersProperty{
//   		CloudWatchOutputConfig: &CloudWatchOutputConfigProperty{
//   			CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   			CloudWatchOutputEnabled: jsii.Boolean(false),
//   		},
//   		Comment: jsii.String("comment"),
//   		DocumentHash: jsii.String("documentHash"),
//   		DocumentHashType: jsii.String("documentHashType"),
//   		DocumentVersion: jsii.String("documentVersion"),
//   		NotificationConfig: &NotificationConfigProperty{
//   			NotificationArn: jsii.String("notificationArn"),
//
//   			// the properties below are optional
//   			NotificationEvents: []*string{
//   				jsii.String("notificationEvents"),
//   			},
//   			NotificationType: jsii.String("notificationType"),
//   		},
//   		OutputS3BucketName: jsii.String("outputS3BucketName"),
//   		OutputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   		Parameters: parameters,
//   		ServiceRoleArn: jsii.String("serviceRoleArn"),
//   		TimeoutSeconds: jsii.Number(123),
//   	},
//   	MaintenanceWindowStepFunctionsParameters: &MaintenanceWindowStepFunctionsParametersProperty{
//   		Input: jsii.String("input"),
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html
//
type CfnMaintenanceWindowTask_TaskInvocationParametersProperty struct {
	// The parameters for an `AUTOMATION` task type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowautomationparameters
	//
	MaintenanceWindowAutomationParameters interface{} `field:"optional" json:"maintenanceWindowAutomationParameters" yaml:"maintenanceWindowAutomationParameters"`
	// The parameters for a `LAMBDA` task type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowlambdaparameters
	//
	MaintenanceWindowLambdaParameters interface{} `field:"optional" json:"maintenanceWindowLambdaParameters" yaml:"maintenanceWindowLambdaParameters"`
	// The parameters for a `RUN_COMMAND` task type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters
	//
	MaintenanceWindowRunCommandParameters interface{} `field:"optional" json:"maintenanceWindowRunCommandParameters" yaml:"maintenanceWindowRunCommandParameters"`
	// The parameters for a `STEP_FUNCTIONS` task type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowstepfunctionsparameters
	//
	MaintenanceWindowStepFunctionsParameters interface{} `field:"optional" json:"maintenanceWindowStepFunctionsParameters" yaml:"maintenanceWindowStepFunctionsParameters"`
}

