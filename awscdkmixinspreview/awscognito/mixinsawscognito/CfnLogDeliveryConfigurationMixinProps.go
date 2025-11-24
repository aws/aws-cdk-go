package mixinsawscognito


// Properties for CfnLogDeliveryConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLogDeliveryConfigurationMixinProps := &CfnLogDeliveryConfigurationMixinProps{
//   	LogConfigurations: []interface{}{
//   		&LogConfigurationProperty{
//   			CloudWatchLogsConfiguration: &CloudWatchLogsConfigurationProperty{
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   			EventSource: jsii.String("eventSource"),
//   			FirehoseConfiguration: &FirehoseConfigurationProperty{
//   				StreamArn: jsii.String("streamArn"),
//   			},
//   			LogLevel: jsii.String("logLevel"),
//   			S3Configuration: &S3ConfigurationProperty{
//   				BucketArn: jsii.String("bucketArn"),
//   			},
//   		},
//   	},
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-logdeliveryconfiguration.html
//
type CfnLogDeliveryConfigurationMixinProps struct {
	// A logging destination of a user pool.
	//
	// User pools can have multiple logging destinations for message-delivery and user-activity logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-logdeliveryconfiguration.html#cfn-cognito-logdeliveryconfiguration-logconfigurations
	//
	LogConfigurations interface{} `field:"optional" json:"logConfigurations" yaml:"logConfigurations"`
	// The ID of the user pool where you configured logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-logdeliveryconfiguration.html#cfn-cognito-logdeliveryconfiguration-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

