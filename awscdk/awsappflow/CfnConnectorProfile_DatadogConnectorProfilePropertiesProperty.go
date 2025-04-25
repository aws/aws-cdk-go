package awsappflow


// The connector-specific profile properties required by Datadog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datadogConnectorProfilePropertiesProperty := &DatadogConnectorProfilePropertiesProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-datadogconnectorprofileproperties.html
//
type CfnConnectorProfile_DatadogConnectorProfilePropertiesProperty struct {
	// The location of the Datadog resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-datadogconnectorprofileproperties.html#cfn-appflow-connectorprofile-datadogconnectorprofileproperties-instanceurl
	//
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

