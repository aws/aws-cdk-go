package awscdk


// Contains logging configuration information for an extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigProperty := &LoggingConfigProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogRoleArn: jsii.String("logRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-typeactivation-loggingconfig.html
//
type CfnTypeActivation_LoggingConfigProperty struct {
	// The Amazon CloudWatch Logs group to which CloudFormation sends error logging information when invoking the extension's handlers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-typeactivation-loggingconfig.html#cfn-cloudformation-typeactivation-loggingconfig-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The Amazon Resource Name (ARN) of the role that CloudFormation should assume when sending log entries to CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-typeactivation-loggingconfig.html#cfn-cloudformation-typeactivation-loggingconfig-logrolearn
	//
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}

