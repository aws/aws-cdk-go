package awsappflow


// The connector-specific profile properties required when using Salesforce.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceConnectorProfilePropertiesProperty := &salesforceConnectorProfilePropertiesProperty{
//   	instanceUrl: jsii.String("instanceUrl"),
//   	isSandboxEnvironment: jsii.Boolean(false),
//   }
//
type CfnConnectorProfile_SalesforceConnectorProfilePropertiesProperty struct {
	// The location of the Salesforce resource.
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
	// Indicates whether the connector profile applies to a sandbox or production environment.
	IsSandboxEnvironment interface{} `field:"optional" json:"isSandboxEnvironment" yaml:"isSandboxEnvironment"`
}

