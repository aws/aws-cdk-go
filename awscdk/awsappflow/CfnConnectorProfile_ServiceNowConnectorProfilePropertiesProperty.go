package awsappflow


// The connector-specific profile properties required when using ServiceNow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowConnectorProfilePropertiesProperty := &ServiceNowConnectorProfilePropertiesProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-servicenowconnectorprofileproperties.html
//
type CfnConnectorProfile_ServiceNowConnectorProfilePropertiesProperty struct {
	// The location of the ServiceNow resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-servicenowconnectorprofileproperties.html#cfn-appflow-connectorprofile-servicenowconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

