package awsappflow


// The connector-specific profile properties required when using Zendesk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zendeskConnectorProfilePropertiesProperty := &zendeskConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_ZendeskConnectorProfilePropertiesProperty struct {
	// The location of the Zendesk resource.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

