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
//   }
//
type ExperimentTrialComponentReference struct {
	// The Arn of the ExperimentTrialComponent resource.
	ExperimentTrialComponentArn *string `field:"required" json:"experimentTrialComponentArn" yaml:"experimentTrialComponentArn"`
}

