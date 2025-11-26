package previewawsquicksightmixins


// The custom narrative options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customNarrativeOptionsProperty := &CustomNarrativeOptionsProperty{
//   	Narrative: jsii.String("narrative"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customnarrativeoptions.html
//
type CfnDashboardPropsMixin_CustomNarrativeOptionsProperty struct {
	// The string input of custom narrative.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customnarrativeoptions.html#cfn-quicksight-dashboard-customnarrativeoptions-narrative
	//
	Narrative *string `field:"optional" json:"narrative" yaml:"narrative"`
}

