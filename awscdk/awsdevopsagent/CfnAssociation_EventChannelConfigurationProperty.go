package awsdevopsagent


// EventChannelconfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventChannelConfigurationProperty := &EventChannelConfigurationProperty{
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-eventchannelconfiguration.html
//
type CfnAssociation_EventChannelConfigurationProperty struct {
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-eventchannelconfiguration.html#cfn-devopsagent-association-eventchannelconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
}

