package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   decalSettingsProperty := &DecalSettingsProperty{
//   	DecalColor: jsii.String("decalColor"),
//   	DecalPatternType: jsii.String("decalPatternType"),
//   	DecalStyleType: jsii.String("decalStyleType"),
//   	DecalVisibility: jsii.String("decalVisibility"),
//   	ElementValue: jsii.String("elementValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decalsettings.html
//
type CfnAnalysisPropsMixin_DecalSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decalsettings.html#cfn-quicksight-analysis-decalsettings-decalcolor
	//
	DecalColor *string `field:"optional" json:"decalColor" yaml:"decalColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decalsettings.html#cfn-quicksight-analysis-decalsettings-decalpatterntype
	//
	DecalPatternType *string `field:"optional" json:"decalPatternType" yaml:"decalPatternType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decalsettings.html#cfn-quicksight-analysis-decalsettings-decalstyletype
	//
	DecalStyleType *string `field:"optional" json:"decalStyleType" yaml:"decalStyleType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decalsettings.html#cfn-quicksight-analysis-decalsettings-decalvisibility
	//
	DecalVisibility *string `field:"optional" json:"decalVisibility" yaml:"decalVisibility"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decalsettings.html#cfn-quicksight-analysis-decalsettings-elementvalue
	//
	ElementValue *string `field:"optional" json:"elementValue" yaml:"elementValue"`
}

