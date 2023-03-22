package awsevidently


// A structure that defines one treatment in an experiment.
//
// A treatment is a variation of the feature that you are including in the experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   treatmentObjectProperty := &treatmentObjectProperty{
//   	feature: jsii.String("feature"),
//   	treatmentName: jsii.String("treatmentName"),
//   	variation: jsii.String("variation"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnExperiment_TreatmentObjectProperty struct {
	// The name of the feature for this experiment.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// A name for this treatment.
	//
	// It can include up to 127 characters.
	TreatmentName *string `field:"required" json:"treatmentName" yaml:"treatmentName"`
	// The name of the variation to use for this treatment.
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// The description of the treatment.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

