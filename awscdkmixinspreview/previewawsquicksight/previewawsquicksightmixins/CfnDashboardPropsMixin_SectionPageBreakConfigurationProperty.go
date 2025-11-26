package previewawsquicksightmixins


// The configuration of a page break for a section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sectionPageBreakConfigurationProperty := &SectionPageBreakConfigurationProperty{
//   	After: &SectionAfterPageBreakProperty{
//   		Status: jsii.String("status"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionpagebreakconfiguration.html
//
type CfnDashboardPropsMixin_SectionPageBreakConfigurationProperty struct {
	// The configuration of a page break after a section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionpagebreakconfiguration.html#cfn-quicksight-dashboard-sectionpagebreakconfiguration-after
	//
	After interface{} `field:"optional" json:"after" yaml:"after"`
}

