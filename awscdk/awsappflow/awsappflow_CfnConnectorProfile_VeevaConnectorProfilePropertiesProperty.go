package awsappflow


// The connector-specific profile properties required when using Veeva.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   veevaConnectorProfilePropertiesProperty := &veevaConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_VeevaConnectorProfilePropertiesProperty struct {
	// The location of the Veeva resource.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

