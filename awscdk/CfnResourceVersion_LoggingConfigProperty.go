package awscdk


// Logging configuration information for a resource.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-resourceversion-loggingconfig.html
//
type CfnResourceVersion_LoggingConfigProperty struct {
	// The Amazon CloudWatch logs group to which CloudFormation sends error logging information when invoking the type's handlers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-resourceversion-loggingconfig.html#cfn-cloudformation-resourceversion-loggingconfig-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-resourceversion-loggingconfig.html#cfn-cloudformation-resourceversion-loggingconfig-logrolearn
	//
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}

