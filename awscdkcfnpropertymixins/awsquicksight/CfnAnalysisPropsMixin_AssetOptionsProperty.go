package awsquicksight


// An array of analysis level configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   assetOptionsProperty := &AssetOptionsProperty{
//   	CustomActionDefaults: &VisualCustomActionDefaultsProperty{
//   		HighlightOperation: &VisualHighlightOperationProperty{
//   			Trigger: jsii.String("trigger"),
//   		},
//   	},
//   	Timezone: jsii.String("timezone"),
//   	VisualMessages: &VisualMessagesProperty{
//   		NoDataMessage: &VisualMessageConfigurationProperty{
//   			Description: jsii.String("description"),
//   			DescriptionVisibility: jsii.String("descriptionVisibility"),
//   			Enabled: jsii.Boolean(false),
//   			LinkText: jsii.String("linkText"),
//   			LinkUrl: jsii.String("linkUrl"),
//   			LinkVisibility: jsii.String("linkVisibility"),
//   			Title: jsii.String("title"),
//   			TitleVisibility: jsii.String("titleVisibility"),
//   		},
//   	},
//   	WeekStart: jsii.String("weekStart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html
//
type CfnAnalysisPropsMixin_AssetOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html#cfn-quicksight-analysis-assetoptions-customactiondefaults
	//
	CustomActionDefaults interface{} `field:"optional" json:"customActionDefaults" yaml:"customActionDefaults"`
	// Determines the timezone for the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html#cfn-quicksight-analysis-assetoptions-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html#cfn-quicksight-analysis-assetoptions-visualmessages
	//
	VisualMessages interface{} `field:"optional" json:"visualMessages" yaml:"visualMessages"`
	// Determines the week start day for an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html#cfn-quicksight-analysis-assetoptions-weekstart
	//
	WeekStart *string `field:"optional" json:"weekStart" yaml:"weekStart"`
}

