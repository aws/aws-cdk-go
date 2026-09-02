package interfacesawsappconfig


// A reference to a ExperimentDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentDefinitionReference := &ExperimentDefinitionReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	ExperimentDefinitionId: jsii.String("experimentDefinitionId"),
//   }
//
type ExperimentDefinitionReference struct {
	// The ApplicationId of the ExperimentDefinition resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Id of the ExperimentDefinition resource.
	ExperimentDefinitionId *string `field:"required" json:"experimentDefinitionId" yaml:"experimentDefinitionId"`
}

