package previewawsgreengrassmixins


// Properties for CfnConnectorDefinitionVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//
//   cfnConnectorDefinitionVersionMixinProps := &CfnConnectorDefinitionVersionMixinProps{
//   	ConnectorDefinitionId: jsii.String("connectorDefinitionId"),
//   	Connectors: []interface{}{
//   		&ConnectorProperty{
//   			ConnectorArn: jsii.String("connectorArn"),
//   			Id: jsii.String("id"),
//   			Parameters: parameters,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html
//
type CfnConnectorDefinitionVersionMixinProps struct {
	// The ID of the connector definition associated with this version.
	//
	// This value is a GUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html#cfn-greengrass-connectordefinitionversion-connectordefinitionid
	//
	ConnectorDefinitionId *string `field:"optional" json:"connectorDefinitionId" yaml:"connectorDefinitionId"`
	// The connectors in this version.
	//
	// Only one instance of a given connector can be added to the connector definition version at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html#cfn-greengrass-connectordefinitionversion-connectors
	//
	Connectors interface{} `field:"optional" json:"connectors" yaml:"connectors"`
}

