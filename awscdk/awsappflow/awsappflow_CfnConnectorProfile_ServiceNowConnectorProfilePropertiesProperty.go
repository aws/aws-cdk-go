package awsappflow


// The connector-specific profile properties required when using ServiceNow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowConnectorProfilePropertiesProperty := &serviceNowConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_ServiceNowConnectorProfilePropertiesProperty struct {
	// The location of the ServiceNow resource.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

