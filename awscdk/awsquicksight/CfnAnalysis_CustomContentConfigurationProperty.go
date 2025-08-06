package awsquicksight


// The configuration of a `CustomContentVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customContentConfigurationProperty := &CustomContentConfigurationProperty{
//   	ContentType: jsii.String("contentType"),
//   	ContentUrl: jsii.String("contentUrl"),
//   	ImageScaling: jsii.String("imageScaling"),
//   	Interactions: &VisualInteractionOptionsProperty{
//   		ContextMenuOption: &ContextMenuOptionProperty{
//   			AvailabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   		VisualMenuOption: &VisualMenuOptionProperty{
//   			AvailabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customcontentconfiguration.html
//
type CfnAnalysis_CustomContentConfigurationProperty struct {
	// The content type of the custom content visual.
	//
	// You can use this to have the visual render as an image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customcontentconfiguration.html#cfn-quicksight-analysis-customcontentconfiguration-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The input URL that links to the custom content that you want in the custom visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customcontentconfiguration.html#cfn-quicksight-analysis-customcontentconfiguration-contenturl
	//
	ContentUrl *string `field:"optional" json:"contentUrl" yaml:"contentUrl"`
	// The sizing options for the size of the custom content visual.
	//
	// This structure is required when the `ContentType` of the visual is `'IMAGE'` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customcontentconfiguration.html#cfn-quicksight-analysis-customcontentconfiguration-imagescaling
	//
	ImageScaling *string `field:"optional" json:"imageScaling" yaml:"imageScaling"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customcontentconfiguration.html#cfn-quicksight-analysis-customcontentconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
}

