package mixinsawsevidently


// A structure that defines one treatment in an experiment.
//
// A treatment is a variation of the feature that you are including in the experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   treatmentObjectProperty := &TreatmentObjectProperty{
//   	Description: jsii.String("description"),
//   	Feature: jsii.String("feature"),
//   	TreatmentName: jsii.String("treatmentName"),
//   	Variation: jsii.String("variation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html
//
type CfnExperimentPropsMixin_TreatmentObjectProperty struct {
	// The description of the treatment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html#cfn-evidently-experiment-treatmentobject-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the feature for this experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html#cfn-evidently-experiment-treatmentobject-feature
	//
	Feature *string `field:"optional" json:"feature" yaml:"feature"`
	// A name for this treatment.
	//
	// It can include up to 127 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html#cfn-evidently-experiment-treatmentobject-treatmentname
	//
	TreatmentName *string `field:"optional" json:"treatmentName" yaml:"treatmentName"`
	// The name of the variation to use for this treatment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-experiment-treatmentobject.html#cfn-evidently-experiment-treatmentobject-variation
	//
	Variation *string `field:"optional" json:"variation" yaml:"variation"`
}

