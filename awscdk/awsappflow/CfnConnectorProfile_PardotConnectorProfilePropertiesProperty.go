package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pardotConnectorProfilePropertiesProperty := &PardotConnectorProfilePropertiesProperty{
//   	BusinessUnitId: jsii.String("businessUnitId"),
//
//   	// the properties below are optional
//   	InstanceUrl: jsii.String("instanceUrl"),
//   	IsSandboxEnvironment: jsii.Boolean(false),
//   }
//
type CfnConnectorProfile_PardotConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.PardotConnectorProfilePropertiesProperty.BusinessUnitId`.
	BusinessUnitId *string `field:"required" json:"businessUnitId" yaml:"businessUnitId"`
	// `CfnConnectorProfile.PardotConnectorProfilePropertiesProperty.InstanceUrl`.
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// `CfnConnectorProfile.PardotConnectorProfilePropertiesProperty.IsSandboxEnvironment`.
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
}

