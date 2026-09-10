package interfacesawsappconfig


// A reference to a ExperimentRun resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentRunReference := &ExperimentRunReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	ExperimentDefinitionId: jsii.String("experimentDefinitionId"),
//   	Run: jsii.String("run"),
//   }
//
type ExperimentRunReference struct {
	// The ApplicationId of the ExperimentRun resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ExperimentDefinitionId of the ExperimentRun resource.
	ExperimentDefinitionId *string `field:"required" json:"experimentDefinitionId" yaml:"experimentDefinitionId"`
	// The Run of the ExperimentRun resource.
	Run *string `field:"required" json:"run" yaml:"run"`
}

