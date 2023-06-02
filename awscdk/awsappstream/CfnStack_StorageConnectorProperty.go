package awsappstream


// A connector that enables persistent storage for users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageConnectorProperty := &StorageConnectorProperty{
//   	ConnectorType: jsii.String("connectorType"),
//
//   	// the properties below are optional
//   	Domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   }
//
type CfnStack_StorageConnectorProperty struct {
	// The type of storage connector.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// The names of the domains for the account.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// The ARN of the storage connector.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

