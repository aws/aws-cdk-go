package awsquicksight


// An array of analysis level configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetOptionsProperty := &AssetOptionsProperty{
//   	ExcludedDataSetArns: []*string{
//   		jsii.String("excludedDataSetArns"),
//   	},
//   	QBusinessInsightsStatus: jsii.String("qBusinessInsightsStatus"),
//   	Timezone: jsii.String("timezone"),
//   	WeekStart: jsii.String("weekStart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-assetoptions.html
//
type CfnDashboard_AssetOptionsProperty struct {
	// A list of dataset ARNS to exclude from Dashboard Q&A.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-assetoptions.html#cfn-quicksight-dashboard-assetoptions-excludeddatasetarns
	//
	ExcludedDataSetArns *[]*string `field:"optional" json:"excludedDataSetArns" yaml:"excludedDataSetArns"`
	// Determines whether insight summaries from Amazon Q Business are allowed in Dashboard Q&A.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-assetoptions.html#cfn-quicksight-dashboard-assetoptions-qbusinessinsightsstatus
	//
	QBusinessInsightsStatus *string `field:"optional" json:"qBusinessInsightsStatus" yaml:"qBusinessInsightsStatus"`
	// Determines the timezone for the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-assetoptions.html#cfn-quicksight-dashboard-assetoptions-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Determines the week start day for an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-assetoptions.html#cfn-quicksight-dashboard-assetoptions-weekstart
	//
	WeekStart *string `field:"optional" json:"weekStart" yaml:"weekStart"`
}

