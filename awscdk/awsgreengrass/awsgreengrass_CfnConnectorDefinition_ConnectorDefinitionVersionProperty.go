package awsgreengrass


// A connector definition version contains a list of connectors.
//
// > After you create a connector definition version that contains the connectors you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `ConnectorDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::ConnectorDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   connectorDefinitionVersionProperty := &connectorDefinitionVersionProperty{
//   	connectors: []interface{}{
//   		&connectorProperty{
//   			connectorArn: jsii.String("connectorArn"),
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			parameters: parameters,
//   		},
//   	},
//   }
//
type CfnConnectorDefinition_ConnectorDefinitionVersionProperty struct {
	// The connectors in this version.
	//
	// Only one instance of a given connector can be added to a connector definition version at a time.
	Connectors interface{} `field:"required" json:"connectors" yaml:"connectors"`
}

