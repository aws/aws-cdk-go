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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-storageconnector.html
//
type CfnStack_StorageConnectorProperty struct {
	// The type of storage connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-storageconnector.html#cfn-appstream-stack-storageconnector-connectortype
	//
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// The names of the domains for the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-storageconnector.html#cfn-appstream-stack-storageconnector-domains
	//
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// The ARN of the storage connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-storageconnector.html#cfn-appstream-stack-storageconnector-resourceidentifier
	//
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

