package awsiot


// Properties for defining a `CfnTopicRuleDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTopicRuleDestinationProps := &CfnTopicRuleDestinationProps{
//   	HttpUrlProperties: &HttpUrlDestinationSummaryProperty{
//   		ConfirmationUrl: jsii.String("confirmationUrl"),
//   	},
//   	InfluxDbProperties: &InfluxDBDestinationPropertiesProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		InfluxDbVersion: jsii.String("influxDbVersion"),
//   		SecretId: jsii.String("secretId"),
//
//   		// the properties below are optional
//   		SecretKey: jsii.String("secretKey"),
//   		SecretType: jsii.String("secretType"),
//   	},
//   	Status: jsii.String("status"),
//   	VpcProperties: &VpcDestinationPropertiesProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicruledestination.html
//
type CfnTopicRuleDestinationProps struct {
	// Properties of the HTTP URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicruledestination.html#cfn-iot-topicruledestination-httpurlproperties
	//
	HttpUrlProperties interface{} `field:"optional" json:"httpUrlProperties" yaml:"httpUrlProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicruledestination.html#cfn-iot-topicruledestination-influxdbproperties
	//
	InfluxDbProperties interface{} `field:"optional" json:"influxDbProperties" yaml:"influxDbProperties"`
	// - **IN_PROGRESS** - A topic rule destination was created but has not been confirmed.
	//
	// You can set status to `IN_PROGRESS` by calling `UpdateTopicRuleDestination` . Calling `UpdateTopicRuleDestination` causes a new confirmation challenge to be sent to your confirmation endpoint.
	// - **ENABLED** - Confirmation was completed, and traffic to this destination is allowed. You can set status to `DISABLED` by calling `UpdateTopicRuleDestination` .
	// - **DISABLED** - Confirmation was completed, and traffic to this destination is not allowed. You can set status to `ENABLED` by calling `UpdateTopicRuleDestination` .
	// - **ERROR** - Confirmation could not be completed; for example, if the confirmation timed out. You can call `GetTopicRuleDestination` for details about the error. You can set status to `IN_PROGRESS` by calling `UpdateTopicRuleDestination` . Calling `UpdateTopicRuleDestination` causes a new confirmation challenge to be sent to your confirmation endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicruledestination.html#cfn-iot-topicruledestination-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Properties of the virtual private cloud (VPC) connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicruledestination.html#cfn-iot-topicruledestination-vpcproperties
	//
	VpcProperties interface{} `field:"optional" json:"vpcProperties" yaml:"vpcProperties"`
}

