package awsappflow


// The connector-specific profile properties required by Infor Nexus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inforNexusConnectorProfilePropertiesProperty := &inforNexusConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   }
//
type CfnConnectorProfile_InforNexusConnectorProfilePropertiesProperty struct {
	// The location of the Infor Nexus resource.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

