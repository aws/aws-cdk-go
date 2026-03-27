package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trainedModelExportsMaxSizeProperty := &TrainedModelExportsMaxSizeProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsmaxsize.html
//
type CfnConfiguredModelAlgorithmAssociationPropsMixin_TrainedModelExportsMaxSizeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsmaxsize.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsmaxsize-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsmaxsize.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsmaxsize-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

