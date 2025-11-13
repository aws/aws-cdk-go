package interfacesawslogs


// A reference to a QueryDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryDefinitionReference := &QueryDefinitionReference{
//   	QueryDefinitionId: jsii.String("queryDefinitionId"),
//   }
//
type QueryDefinitionReference struct {
	// The QueryDefinitionId of the QueryDefinition resource.
	QueryDefinitionId *string `field:"required" json:"queryDefinitionId" yaml:"queryDefinitionId"`
}

