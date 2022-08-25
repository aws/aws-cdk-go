package awsgreengrass


// Connectors are modules that provide built-in integration with local infrastructure, device protocols, AWS , and other cloud services.
//
// For more information, see [Integrate with Services and Protocols Using Greengrass Connectors](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Connectors` property of the [`ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinition-connectordefinitionversion.html) property type contains a list of `Connector` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   connectorProperty := &connectorProperty{
//   	connectorArn: jsii.String("connectorArn"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	parameters: parameters,
//   }
//
type CfnConnectorDefinition_ConnectorProperty struct {
	// The Amazon Resource Name (ARN) of the connector.
	//
	// For more information about connectors provided by AWS , see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors-list.html) .
	ConnectorArn *string `field:"required" json:"connectorArn" yaml:"connectorArn"`
	// A descriptive or arbitrary ID for the connector.
	//
	// This value must be unique within the connector definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The parameters or configuration used by the connector.
	//
	// For more information about connectors provided by AWS , see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/latest/developerguide/connectors-list.html) .
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

