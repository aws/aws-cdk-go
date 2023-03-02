package awsappflow


// The connector-specific profile properties required when using Slack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackConnectorProfilePropertiesProperty := &slackConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_SlackConnectorProfilePropertiesProperty struct {
	// The location of the Slack resource.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

