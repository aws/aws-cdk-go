package previewawsquicksightmixins


// The page break configuration to apply for each repeating instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bodySectionRepeatPageBreakConfigurationProperty := &BodySectionRepeatPageBreakConfigurationProperty{
//   	After: &SectionAfterPageBreakProperty{
//   		Status: jsii.String("status"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionrepeatpagebreakconfiguration.html
//
type CfnAnalysisPropsMixin_BodySectionRepeatPageBreakConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionrepeatpagebreakconfiguration.html#cfn-quicksight-analysis-bodysectionrepeatpagebreakconfiguration-after
	//
	After interface{} `field:"optional" json:"after" yaml:"after"`
}

