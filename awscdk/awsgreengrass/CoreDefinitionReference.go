package awsgreengrass


// A reference to a CoreDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreDefinitionReference := &CoreDefinitionReference{
//   	CoreDefinitionArn: jsii.String("coreDefinitionArn"),
//   	CoreDefinitionId: jsii.String("coreDefinitionId"),
//   }
//
type CoreDefinitionReference struct {
	// The ARN of the CoreDefinition resource.
	CoreDefinitionArn *string `field:"required" json:"coreDefinitionArn" yaml:"coreDefinitionArn"`
	// The Id of the CoreDefinition resource.
	CoreDefinitionId *string `field:"required" json:"coreDefinitionId" yaml:"coreDefinitionId"`
}

