package awsssm


// The `MaintenanceWindowRunCommandParameters` property type specifies the parameters for a `RUN_COMMAND` task type for a maintenance window task in AWS Systems Manager .
//
// This means that these parameters are the same as those for the `SendCommand` API call. For more information about `SendCommand` parameters, see [SendCommand](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_SendCommand.html) in the *AWS Systems Manager API Reference* .
//
// For information about available parameters in SSM Command documents, you can view the content of the document itself in the Systems Manager console. For information, see [Viewing SSM command document content](https://docs.aws.amazon.com/systems-manager/latest/userguide/viewing-ssm-document-content.html) in the *AWS Systems Manager User Guide* .
//
// `MaintenanceWindowRunCommandParameters` is a property of the [TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   maintenanceWindowRunCommandParametersProperty := &maintenanceWindowRunCommandParametersProperty{
//   	cloudWatchOutputConfig: &cloudWatchOutputConfigProperty{
//   		cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		cloudWatchOutputEnabled: jsii.Boolean(false),
//   	},
//   	comment: jsii.String("comment"),
//   	documentHash: jsii.String("documentHash"),
//   	documentHashType: jsii.String("documentHashType"),
//   	documentVersion: jsii.String("documentVersion"),
//   	notificationConfig: &notificationConfigProperty{
//   		notificationArn: jsii.String("notificationArn"),
//
//   		// the properties below are optional
//   		notificationEvents: []*string{
//   			jsii.String("notificationEvents"),
//   		},
//   		notificationType: jsii.String("notificationType"),
//   	},
//   	outputS3BucketName: jsii.String("outputS3BucketName"),
//   	outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   	parameters: parameters,
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//   	timeoutSeconds: jsii.Number(123),
//   }
//
type CfnMaintenanceWindowTask_MaintenanceWindowRunCommandParametersProperty struct {
	// Configuration options for sending command output to Amazon CloudWatch Logs.
	CloudWatchOutputConfig interface{} `field:"optional" json:"cloudWatchOutputConfig" yaml:"cloudWatchOutputConfig"`
	// Information about the command or commands to run.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The SHA-256 or SHA-1 hash created by the system when the document was created.
	//
	// SHA-1 hashes have been deprecated.
	DocumentHash *string `field:"optional" json:"documentHash" yaml:"documentHash"`
	// The SHA-256 or SHA-1 hash type.
	//
	// SHA-1 hashes are deprecated.
	DocumentHashType *string `field:"optional" json:"documentHashType" yaml:"documentHashType"`
	// The AWS Systems Manager document (SSM document) version to use in the request.
	//
	// You can specify `$DEFAULT` , `$LATEST` , or a specific version number. If you run commands by using the AWS CLI, then you must escape the first two options by using a backslash. If you specify a version number, then you don't need to use the backslash. For example:
	//
	// `--document-version "\$DEFAULT"`
	//
	// `--document-version "\$LATEST"`
	//
	// `--document-version "3"`.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// Configurations for sending notifications about command status changes on a per-managed node basis.
	NotificationConfig interface{} `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// The name of the Amazon Simple Storage Service (Amazon S3) bucket.
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
	// The S3 bucket subfolder.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// The parameters for the `RUN_COMMAND` task execution.
	//
	// The supported parameters are the same as those for the `SendCommand` API call. For more information, see [SendCommand](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_SendCommand.html) in the *AWS Systems Manager API Reference* .
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// If this time is reached and the command hasn't already started running, it doesn't run.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

