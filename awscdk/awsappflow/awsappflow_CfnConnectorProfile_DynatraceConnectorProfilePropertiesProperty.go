package awsappflow


// The connector-specific profile properties required by Dynatrace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynatraceConnectorProfilePropertiesProperty := &dynatraceConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_DynatraceConnectorProfilePropertiesProperty struct {
	// The location of the Dynatrace resource.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

