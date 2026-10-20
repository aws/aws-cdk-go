package awsworkspaces


// Describes the storage connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   storageConnectorProperty := &StorageConnectorProperty{
//   	ConnectorType: jsii.String("connectorType"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-storageconnector.html
//
type CfnDirectoryPropsMixin_StorageConnectorProperty struct {
	// The type of connector used to save user files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-storageconnector.html#cfn-workspaces-directory-storageconnector-connectortype
	//
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// Indicates if the storage connector is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-storageconnector.html#cfn-workspaces-directory-storageconnector-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

