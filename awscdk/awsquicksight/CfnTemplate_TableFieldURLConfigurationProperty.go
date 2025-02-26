package awsquicksight


// The URL configuration for a table field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldURLConfigurationProperty := &TableFieldURLConfigurationProperty{
//   	ImageConfiguration: &TableFieldImageConfigurationProperty{
//   		SizingOptions: &TableCellImageSizingConfigurationProperty{
//   			TableCellImageScalingConfiguration: jsii.String("tableCellImageScalingConfiguration"),
//   		},
//   	},
//   	LinkConfiguration: &TableFieldLinkConfigurationProperty{
//   		Content: &TableFieldLinkContentConfigurationProperty{
//   			CustomIconContent: &TableFieldCustomIconContentProperty{
//   				Icon: jsii.String("icon"),
//   			},
//   			CustomTextContent: &TableFieldCustomTextContentProperty{
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Target: jsii.String("target"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldurlconfiguration.html
//
type CfnTemplate_TableFieldURLConfigurationProperty struct {
	// The image configuration of a table field URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldurlconfiguration.html#cfn-quicksight-template-tablefieldurlconfiguration-imageconfiguration
	//
	ImageConfiguration interface{} `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// The link configuration of a table field URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldurlconfiguration.html#cfn-quicksight-template-tablefieldurlconfiguration-linkconfiguration
	//
	LinkConfiguration interface{} `field:"optional" json:"linkConfiguration" yaml:"linkConfiguration"`
}

