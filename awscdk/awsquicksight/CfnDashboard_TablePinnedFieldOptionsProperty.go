package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tablePinnedFieldOptionsProperty := &TablePinnedFieldOptionsProperty{
//   	PinnedLeftFields: []*string{
//   		jsii.String("pinnedLeftFields"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablepinnedfieldoptions.html
//
type CfnDashboard_TablePinnedFieldOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablepinnedfieldoptions.html#cfn-quicksight-dashboard-tablepinnedfieldoptions-pinnedleftfields
	//
	PinnedLeftFields *[]*string `field:"optional" json:"pinnedLeftFields" yaml:"pinnedLeftFields"`
}

