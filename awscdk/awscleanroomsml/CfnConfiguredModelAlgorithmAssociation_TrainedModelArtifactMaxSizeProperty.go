package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainedModelArtifactMaxSizeProperty := &TrainedModelArtifactMaxSizeProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelartifactmaxsize.html
//
type CfnConfiguredModelAlgorithmAssociation_TrainedModelArtifactMaxSizeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelartifactmaxsize.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelartifactmaxsize-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelartifactmaxsize.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelartifactmaxsize-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

