package awsappflow


// The connector-specific profile properties required when using Marketo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   marketoConnectorProfilePropertiesProperty := &marketoConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_MarketoConnectorProfilePropertiesProperty struct {
	// The location of the Marketo resource.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

