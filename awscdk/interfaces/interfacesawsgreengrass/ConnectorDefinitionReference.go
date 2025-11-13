package interfacesawsgreengrass


// A reference to a ConnectorDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorDefinitionReference := &ConnectorDefinitionReference{
//   	ConnectorDefinitionArn: jsii.String("connectorDefinitionArn"),
//   	ConnectorDefinitionId: jsii.String("connectorDefinitionId"),
//   }
//
type ConnectorDefinitionReference struct {
	// The ARN of the ConnectorDefinition resource.
	ConnectorDefinitionArn *string `field:"required" json:"connectorDefinitionArn" yaml:"connectorDefinitionArn"`
	// The Id of the ConnectorDefinition resource.
	ConnectorDefinitionId *string `field:"required" json:"connectorDefinitionId" yaml:"connectorDefinitionId"`
}

