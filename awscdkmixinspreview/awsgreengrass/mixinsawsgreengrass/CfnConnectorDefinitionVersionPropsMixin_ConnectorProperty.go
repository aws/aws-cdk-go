package mixinsawsgreengrass


// Connectors are modules that provide built-in integration with local infrastructure, device protocols, AWS , and other cloud services.
//
// For more information, see [Integrate with Services and Protocols Using Greengrass Connectors](https://docs.aws.amazon.com/greengrass/v1/developerguide/connectors.html) in the *Developer Guide* .
//
// In an CloudFormation template, the `Connectors` property of the [`AWS::Greengrass::ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html) resource contains a list of `Connector` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//
//   connectorProperty := &ConnectorProperty{
//   	ConnectorArn: jsii.String("connectorArn"),
//   	Id: jsii.String("id"),
//   	Parameters: parameters,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinitionversion-connector.html
//
type CfnConnectorDefinitionVersionPropsMixin_ConnectorProperty struct {
	// The Amazon Resource Name (ARN) of the connector.
	//
	// For more information about connectors provided by AWS , see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/v1/developerguide/connectors-list.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinitionversion-connector.html#cfn-greengrass-connectordefinitionversion-connector-connectorarn
	//
	ConnectorArn *string `field:"optional" json:"connectorArn" yaml:"connectorArn"`
	// A descriptive or arbitrary ID for the connector.
	//
	// This value must be unique within the connector definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinitionversion-connector.html#cfn-greengrass-connectordefinitionversion-connector-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The parameters or configuration that the connector uses.
	//
	// For more information about connectors provided by AWS , see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/v1/developerguide/connectors-list.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinitionversion-connector.html#cfn-greengrass-connectordefinitionversion-connector-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

