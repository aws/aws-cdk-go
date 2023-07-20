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
//   maintenanceWindowRunCommandParametersProperty := &MaintenanceWindowRunCommandParametersProperty{
//   	CloudWatchOutputConfig: &CloudWatchOutputConfigProperty{
//   		CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		CloudWatchOutputEnabled: jsii.Boolean(false),
//   	},
//   	Comment: jsii.String("comment"),
//   	DocumentHash: jsii.String("documentHash"),
//   	DocumentHashType: jsii.String("documentHashType"),
//   	DocumentVersion: jsii.String("documentVersion"),
//   	NotificationConfig: &NotificationConfigProperty{
//   		NotificationArn: jsii.String("notificationArn"),
//
//   		// the properties below are optional
//   		NotificationEvents: []*string{
//   			jsii.String("notificationEvents"),
//   		},
//   		NotificationType: jsii.String("notificationType"),
//   	},
//   	OutputS3BucketName: jsii.String("outputS3BucketName"),
//   	OutputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   	Parameters: parameters,
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//   	TimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html
//
type CfnMaintenanceWindowTask_MaintenanceWindowRunCommandParametersProperty struct {
	// Configuration options for sending command output to Amazon CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-cloudwatchoutputconfig
	//
	CloudWatchOutputConfig interface{} `field:"optional" json:"cloudWatchOutputConfig" yaml:"cloudWatchOutputConfig"`
	// Information about the command or commands to run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The SHA-256 or SHA-1 hash created by the system when the document was created.
	//
	// SHA-1 hashes have been deprecated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-documenthash
	//
	DocumentHash *string `field:"optional" json:"documentHash" yaml:"documentHash"`
	// The SHA-256 or SHA-1 hash type.
	//
	// SHA-1 hashes are deprecated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-documenthashtype
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-documentversion
	//
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// Configurations for sending notifications about command status changes on a per-managed node basis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-notificationconfig
	//
	NotificationConfig interface{} `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// The name of the Amazon Simple Storage Service (Amazon S3) bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-outputs3bucketname
	//
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
	// The S3 bucket subfolder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-outputs3keyprefix
	//
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// The parameters for the `RUN_COMMAND` task execution.
	//
	// The supported parameters are the same as those for the `SendCommand` API call. For more information, see [SendCommand](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_SendCommand.html) in the *AWS Systems Manager API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-servicerolearn
	//
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// If this time is reached and the command hasn't already started running, it doesn't run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowruncommandparameters-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

