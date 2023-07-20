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
//   treatmentObjectProperty := &TreatmentObjectProperty{
//   	Feature: jsii.String("feature"),
//   	TreatmentName: jsii.String("treatmentName"),
//   	Variation: jsii.String("variation"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html
//
type CfnExperiment_TreatmentObjectProperty struct {
	// The name of the feature for this experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html#cfn-evidently-experiment-treatmentobject-feature
	//
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// A name for this treatment.
	//
	// It can include up to 127 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html#cfn-evidently-experiment-treatmentobject-treatmentname
	//
	TreatmentName *string `field:"required" json:"treatmentName" yaml:"treatmentName"`
	// The name of the variation to use for this treatment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html#cfn-evidently-experiment-treatmentobject-variation
	//
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// The description of the treatment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html#cfn-evidently-experiment-treatmentobject-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

