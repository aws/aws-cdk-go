package interfacesawssagemaker


// A reference to a TrialComponent resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trialComponentReference := &TrialComponentReference{
//   	TrialComponentArn: jsii.String("trialComponentArn"),
//   }
//
type TrialComponentReference struct {
	// The TrialComponentArn of the TrialComponent resource.
	TrialComponentArn *string `field:"required" json:"trialComponentArn" yaml:"trialComponentArn"`
}

