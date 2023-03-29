package awsappflow


// The connector-specific profile properties required when using Salesforce Pardot.
//
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
	// The business unit id of Salesforce Pardot instance.
	BusinessUnitId *string `field:"required" json:"businessUnitId" yaml:"businessUnitId"`
	// The location of the Salesforce Pardot resource.
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// Indicates whether the connector profile applies to a sandbox or production environment.
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
}

