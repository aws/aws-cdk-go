package awsappstream


// A connector that enables persistent storage for users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageConnectorProperty := &storageConnectorProperty{
//   	connectorType: jsii.String("connectorType"),
//
//   	// the properties below are optional
//   	domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
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

