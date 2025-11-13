package interfacesawsgreengrass


// A reference to a CoreDefinitionVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreDefinitionVersionReference := &CoreDefinitionVersionReference{
//   	CoreDefinitionVersionId: jsii.String("coreDefinitionVersionId"),
//   }
//
type CoreDefinitionVersionReference struct {
	// The Id of the CoreDefinitionVersion resource.
	CoreDefinitionVersionId *string `field:"required" json:"coreDefinitionVersionId" yaml:"coreDefinitionVersionId"`
}

