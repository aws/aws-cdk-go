package interfacesawssagemaker


// A reference to a ExperimentTrialComponent resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTrialComponentReference := &ExperimentTrialComponentReference{
//   	ExperimentTrialComponentArn: jsii.String("experimentTrialComponentArn"),
//   	ExperimentTrialComponentId: jsii.String("experimentTrialComponentId"),
//   }
//
type ExperimentTrialComponentReference struct {
	// The ARN of the ExperimentTrialComponent resource.
	ExperimentTrialComponentArn *string `field:"required" json:"experimentTrialComponentArn" yaml:"experimentTrialComponentArn"`
	// The Id of the ExperimentTrialComponent resource.
	ExperimentTrialComponentId *string `field:"required" json:"experimentTrialComponentId" yaml:"experimentTrialComponentId"`
}

