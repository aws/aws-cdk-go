package awslookoutmetrics


// A configuration that specifies the action to perform when anomalies are detected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	LambdaConfiguration: &LambdaConfigurationProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SnsConfiguration: &SNSConfigurationProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SnsTopicArn: jsii.String("snsTopicArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-alert-action.html
//
type CfnAlert_ActionProperty struct {
	// A configuration for an AWS Lambda channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-alert-action.html#cfn-lookoutmetrics-alert-action-lambdaconfiguration
	//
	LambdaConfiguration interface{} `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// A configuration for an Amazon SNS channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-alert-action.html#cfn-lookoutmetrics-alert-action-snsconfiguration
	//
	SnsConfiguration interface{} `field:"optional" json:"snsConfiguration" yaml:"snsConfiguration"`
}

