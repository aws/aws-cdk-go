package mixinsawsaps


// Contains information about the rules and alerting logging configuration for the workspace.
//
// > These logging configurations are only for rules and alerting logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loggingConfigurationProperty := &LoggingConfigurationProperty{
//   	LogGroupArn: jsii.String("logGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingconfiguration.html
//
type CfnWorkspacePropsMixin_LoggingConfigurationProperty struct {
	// The ARN of the CloudWatch log group to which the vended log data will be published.
	//
	// This log group must exist prior to calling this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingconfiguration.html#cfn-aps-workspace-loggingconfiguration-loggrouparn
	//
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}

