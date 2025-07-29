package awsgreengrass


// Properties for defining a `CfnConnectorDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnConnectorDefinitionVersionProps := &CfnConnectorDefinitionVersionProps{
//   	ConnectorDefinitionId: jsii.String("connectorDefinitionId"),
//   	Connectors: []interface{}{
//   		&ConnectorProperty{
//   			ConnectorArn: jsii.String("connectorArn"),
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			Parameters: parameters,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html
//
type CfnConnectorDefinitionVersionProps struct {
	// The ID of the connector definition associated with this version.
	//
	// This value is a GUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html#cfn-greengrass-connectordefinitionversion-connectordefinitionid
	//
	ConnectorDefinitionId *string `field:"required" json:"connectorDefinitionId" yaml:"connectorDefinitionId"`
	// The connectors in this version.
	//
	// Only one instance of a given connector can be added to the connector definition version at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html#cfn-greengrass-connectordefinitionversion-connectors
	//
	Connectors interface{} `field:"required" json:"connectors" yaml:"connectors"`
}

