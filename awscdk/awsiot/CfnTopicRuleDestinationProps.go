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
type CfnTopicRuleDestinationProps struct {
	// Properties of the HTTP URL.
	HttpUrlProperties interface{} `field:"optional" json:"httpUrlProperties" yaml:"httpUrlProperties"`
	// - **IN_PROGRESS** - A topic rule destination was created but has not been confirmed.
	//
	// You can set status to `IN_PROGRESS` by calling `UpdateTopicRuleDestination` . Calling `UpdateTopicRuleDestination` causes a new confirmation challenge to be sent to your confirmation endpoint.
	// - **ENABLED** - Confirmation was completed, and traffic to this destination is allowed. You can set status to `DISABLED` by calling `UpdateTopicRuleDestination` .
	// - **DISABLED** - Confirmation was completed, and traffic to this destination is not allowed. You can set status to `ENABLED` by calling `UpdateTopicRuleDestination` .
	// - **ERROR** - Confirmation could not be completed; for example, if the confirmation timed out. You can call `GetTopicRuleDestination` for details about the error. You can set status to `IN_PROGRESS` by calling `UpdateTopicRuleDestination` . Calling `UpdateTopicRuleDestination` causes a new confirmation challenge to be sent to your confirmation endpoint.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Properties of the virtual private cloud (VPC) connection.
	VpcProperties interface{} `field:"optional" json:"vpcProperties" yaml:"vpcProperties"`
}

