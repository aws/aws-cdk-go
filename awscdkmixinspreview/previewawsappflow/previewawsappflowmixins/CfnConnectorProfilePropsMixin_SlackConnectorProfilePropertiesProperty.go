package previewawsappflowmixins


// The connector-specific profile properties required when using Slack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slackConnectorProfilePropertiesProperty := &SlackConnectorProfilePropertiesProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-slackconnectorprofileproperties.html
//
type CfnConnectorProfilePropsMixin_SlackConnectorProfilePropertiesProperty struct {
	// The location of the Slack resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-slackconnectorprofileproperties.html#cfn-appflow-connectorprofile-slackconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
}

