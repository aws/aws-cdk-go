package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainedModelExportsConfigurationPolicyProperty := &TrainedModelExportsConfigurationPolicyProperty{
//   	FilesToExport: []*string{
//   		jsii.String("filesToExport"),
//   	},
//   	MaxSize: &TrainedModelExportsMaxSizeProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsconfigurationpolicy.html
//
type CfnConfiguredModelAlgorithmAssociation_TrainedModelExportsConfigurationPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsconfigurationpolicy-filestoexport
	//
	FilesToExport *[]*string `field:"required" json:"filesToExport" yaml:"filesToExport"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsconfigurationpolicy-maxsize
	//
	MaxSize interface{} `field:"required" json:"maxSize" yaml:"maxSize"`
}

