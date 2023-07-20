package awsevidently


// This structure defines how much experiment traffic to allocate to one treatment used in the experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   treatmentToWeightProperty := &TreatmentToWeightProperty{
//   	SplitWeight: jsii.Number(123),
//   	Treatment: jsii.String("treatment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmenttoweight.html
//
type CfnExperiment_TreatmentToWeightProperty struct {
	// The portion of experiment traffic to allocate to this treatment.
	//
	// Specify the traffic portion in thousandths of a percent, so 20,000 allocated to a treatment would allocate 20% of the experiment traffic to that treatment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmenttoweight.html#cfn-evidently-experiment-treatmenttoweight-splitweight
	//
	SplitWeight *float64 `field:"required" json:"splitWeight" yaml:"splitWeight"`
	// The name of the treatment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmenttoweight.html#cfn-evidently-experiment-treatmenttoweight-treatment
	//
	Treatment *string `field:"required" json:"treatment" yaml:"treatment"`
}

