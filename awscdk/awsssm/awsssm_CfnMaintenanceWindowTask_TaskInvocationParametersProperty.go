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
//   taskInvocationParametersProperty := &taskInvocationParametersProperty{
//   	maintenanceWindowAutomationParameters: &maintenanceWindowAutomationParametersProperty{
//   		documentVersion: jsii.String("documentVersion"),
//   		parameters: parameters,
//   	},
//   	maintenanceWindowLambdaParameters: &maintenanceWindowLambdaParametersProperty{
//   		clientContext: jsii.String("clientContext"),
//   		payload: jsii.String("payload"),
//   		qualifier: jsii.String("qualifier"),
//   	},
//   	maintenanceWindowRunCommandParameters: &maintenanceWindowRunCommandParametersProperty{
//   		cloudWatchOutputConfig: &cloudWatchOutputConfigProperty{
//   			cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   			cloudWatchOutputEnabled: jsii.Boolean(false),
//   		},
//   		comment: jsii.String("comment"),
//   		documentHash: jsii.String("documentHash"),
//   		documentHashType: jsii.String("documentHashType"),
//   		documentVersion: jsii.String("documentVersion"),
//   		notificationConfig: &notificationConfigProperty{
//   			notificationArn: jsii.String("notificationArn"),
//
//   			// the properties below are optional
//   			notificationEvents: []*string{
//   				jsii.String("notificationEvents"),
//   			},
//   			notificationType: jsii.String("notificationType"),
//   		},
//   		outputS3BucketName: jsii.String("outputS3BucketName"),
//   		outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   		parameters: parameters,
//   		serviceRoleArn: jsii.String("serviceRoleArn"),
//   		timeoutSeconds: jsii.Number(123),
//   	},
//   	maintenanceWindowStepFunctionsParameters: &maintenanceWindowStepFunctionsParametersProperty{
//   		input: jsii.String("input"),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnMaintenanceWindowTask_TaskInvocationParametersProperty struct {
	// The parameters for an `AUTOMATION` task type.
	MaintenanceWindowAutomationParameters interface{} `field:"optional" json:"maintenanceWindowAutomationParameters" yaml:"maintenanceWindowAutomationParameters"`
	// The parameters for a `LAMBDA` task type.
	MaintenanceWindowLambdaParameters interface{} `field:"optional" json:"maintenanceWindowLambdaParameters" yaml:"maintenanceWindowLambdaParameters"`
	// The parameters for a `RUN_COMMAND` task type.
	MaintenanceWindowRunCommandParameters interface{} `field:"optional" json:"maintenanceWindowRunCommandParameters" yaml:"maintenanceWindowRunCommandParameters"`
	// The parameters for a `STEP_FUNCTIONS` task type.
	MaintenanceWindowStepFunctionsParameters interface{} `field:"optional" json:"maintenanceWindowStepFunctionsParameters" yaml:"maintenanceWindowStepFunctionsParameters"`
}

