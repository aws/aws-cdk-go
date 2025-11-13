package interfacesawsgreengrass


// A reference to a ConnectorDefinitionVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorDefinitionVersionReference := &ConnectorDefinitionVersionReference{
//   	ConnectorDefinitionVersionId: jsii.String("connectorDefinitionVersionId"),
//   }
//
type ConnectorDefinitionVersionReference struct {
	// The Id of the ConnectorDefinitionVersion resource.
	ConnectorDefinitionVersionId *string `field:"required" json:"connectorDefinitionVersionId" yaml:"connectorDefinitionVersionId"`
}

