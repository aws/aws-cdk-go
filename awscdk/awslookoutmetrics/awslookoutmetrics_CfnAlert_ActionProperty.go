package awslookoutmetrics


// A configuration that specifies the action to perform when anomalies are detected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	lambdaConfiguration: &lambdaConfigurationProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	snsConfiguration: &sNSConfigurationProperty{
//   		roleArn: jsii.String("roleArn"),
//   		snsTopicArn: jsii.String("snsTopicArn"),
//   	},
//   }
//
type CfnAlert_ActionProperty struct {
	// A configuration for an AWS Lambda channel.
	LambdaConfiguration interface{} `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// A configuration for an Amazon SNS channel.
	SnsConfiguration interface{} `field:"optional" json:"snsConfiguration" yaml:"snsConfiguration"`
}

