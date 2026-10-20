package awsdevopsagent


// Optional bidirectional communication configuration.
//
// Supply this configuration and set Enabled to true so you can interact with the agent directly from Slack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   slackBidirectionalConfigurationProperty := &SlackBidirectionalConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackbidirectionalconfiguration.html
//
type CfnAssociationPropsMixin_SlackBidirectionalConfigurationProperty struct {
	// Whether bidirectional communication is enabled for this association.
	//
	// When set to true, you can mention the agent in a configured Slack channel and it responds in that channel. When omitted or set to false, the agent ignores mentions and only sends notifications
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackbidirectionalconfiguration.html#cfn-devopsagent-association-slackbidirectionalconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// IAM role ARN that AWS DevOps Agent assumes to exchange messages with your Slack workspace on behalf of this association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackbidirectionalconfiguration.html#cfn-devopsagent-association-slackbidirectionalconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

