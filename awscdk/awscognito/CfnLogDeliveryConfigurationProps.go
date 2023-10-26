package awscognito


// Properties for defining a `CfnLogDeliveryConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogDeliveryConfigurationProps := &CfnLogDeliveryConfigurationProps{
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	LogConfigurations: []interface{}{
//   		&LogConfigurationProperty{
//   			CloudWatchLogsConfiguration: &CloudWatchLogsConfigurationProperty{
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   			EventSource: jsii.String("eventSource"),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-logdeliveryconfiguration.html
//
type CfnLogDeliveryConfigurationProps struct {
	// The ID of the user pool where you configured detailed activity logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-logdeliveryconfiguration.html#cfn-cognito-logdeliveryconfiguration-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The detailed activity logging destination of a user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-logdeliveryconfiguration.html#cfn-cognito-logdeliveryconfiguration-logconfigurations
	//
	LogConfigurations interface{} `field:"optional" json:"logConfigurations" yaml:"logConfigurations"`
}

