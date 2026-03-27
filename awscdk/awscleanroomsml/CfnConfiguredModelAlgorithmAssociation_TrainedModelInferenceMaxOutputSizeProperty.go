package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainedModelInferenceMaxOutputSizeProperty := &TrainedModelInferenceMaxOutputSizeProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencemaxoutputsize.html
//
type CfnConfiguredModelAlgorithmAssociation_TrainedModelInferenceMaxOutputSizeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencemaxoutputsize.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencemaxoutputsize-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencemaxoutputsize.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencemaxoutputsize-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

