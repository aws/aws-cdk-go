package awsquicksight


// The options for a table field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldOptionProperty := &TableFieldOptionProperty{
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	CustomLabel: jsii.String("customLabel"),
//   	UrlStyling: &TableFieldURLConfigurationProperty{
//   		ImageConfiguration: &TableFieldImageConfigurationProperty{
//   			SizingOptions: &TableCellImageSizingConfigurationProperty{
//   				TableCellImageScalingConfiguration: jsii.String("tableCellImageScalingConfiguration"),
//   			},
//   		},
//   		LinkConfiguration: &TableFieldLinkConfigurationProperty{
//   			Content: &TableFieldLinkContentConfigurationProperty{
//   				CustomIconContent: &TableFieldCustomIconContentProperty{
//   					Icon: jsii.String("icon"),
//   				},
//   				CustomTextContent: &TableFieldCustomTextContentProperty{
//   					FontConfiguration: &FontConfigurationProperty{
//   						FontColor: jsii.String("fontColor"),
//   						FontDecoration: jsii.String("fontDecoration"),
//   						FontFamily: jsii.String("fontFamily"),
//   						FontSize: &FontSizeProperty{
//   							Absolute: jsii.String("absolute"),
//   							Relative: jsii.String("relative"),
//   						},
//   						FontStyle: jsii.String("fontStyle"),
//   						FontWeight: &FontWeightProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Target: jsii.String("target"),
//   		},
//   	},
//   	Visibility: jsii.String("visibility"),
//   	Width: jsii.String("width"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoption.html
//
type CfnTemplate_TableFieldOptionProperty struct {
	// The field ID for a table field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoption.html#cfn-quicksight-template-tablefieldoption-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The custom label for a table field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoption.html#cfn-quicksight-template-tablefieldoption-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The URL configuration for a table field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoption.html#cfn-quicksight-template-tablefieldoption-urlstyling
	//
	UrlStyling interface{} `field:"optional" json:"urlStyling" yaml:"urlStyling"`
	// The visibility of a table field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoption.html#cfn-quicksight-template-tablefieldoption-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// The width for a table field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoption.html#cfn-quicksight-template-tablefieldoption-width
	//
	Width *string `field:"optional" json:"width" yaml:"width"`
}

