package awsquicksight


// The custom narrative options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customNarrativeOptionsProperty := &CustomNarrativeOptionsProperty{
//   	Narrative: jsii.String("narrative"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customnarrativeoptions.html
//
type CfnDashboard_CustomNarrativeOptionsProperty struct {
	// The string input of custom narrative.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customnarrativeoptions.html#cfn-quicksight-dashboard-customnarrativeoptions-narrative
	//
	Narrative *string `field:"required" json:"narrative" yaml:"narrative"`
}

