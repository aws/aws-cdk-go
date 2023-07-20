package awsappflow


// The connector-specific profile properties required when using Zendesk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zendeskConnectorProfilePropertiesProperty := &ZendeskConnectorProfilePropertiesProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-zendeskconnectorprofileproperties.html
//
type CfnConnectorProfile_ZendeskConnectorProfilePropertiesProperty struct {
	// The location of the Zendesk resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-zendeskconnectorprofileproperties.html#cfn-appflow-connectorprofile-zendeskconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

