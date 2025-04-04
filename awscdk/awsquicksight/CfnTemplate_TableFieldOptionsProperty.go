package awsquicksight


// The field options of a table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldOptionsProperty := &TableFieldOptionsProperty{
//   	Order: []*string{
//   		jsii.String("order"),
//   	},
//   	PinnedFieldOptions: &TablePinnedFieldOptionsProperty{
//   		PinnedLeftFields: []*string{
//   			jsii.String("pinnedLeftFields"),
//   		},
//   	},
//   	SelectedFieldOptions: []interface{}{
//   		&TableFieldOptionProperty{
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			CustomLabel: jsii.String("customLabel"),
//   			UrlStyling: &TableFieldURLConfigurationProperty{
//   				ImageConfiguration: &TableFieldImageConfigurationProperty{
//   					SizingOptions: &TableCellImageSizingConfigurationProperty{
//   						TableCellImageScalingConfiguration: jsii.String("tableCellImageScalingConfiguration"),
//   					},
//   				},
//   				LinkConfiguration: &TableFieldLinkConfigurationProperty{
//   					Content: &TableFieldLinkContentConfigurationProperty{
//   						CustomIconContent: &TableFieldCustomIconContentProperty{
//   							Icon: jsii.String("icon"),
//   						},
//   						CustomTextContent: &TableFieldCustomTextContentProperty{
//   							FontConfiguration: &FontConfigurationProperty{
//   								FontColor: jsii.String("fontColor"),
//   								FontDecoration: jsii.String("fontDecoration"),
//   								FontFamily: jsii.String("fontFamily"),
//   								FontSize: &FontSizeProperty{
//   									Absolute: jsii.String("absolute"),
//   									Relative: jsii.String("relative"),
//   								},
//   								FontStyle: jsii.String("fontStyle"),
//   								FontWeight: &FontWeightProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Target: jsii.String("target"),
//   				},
//   			},
//   			Visibility: jsii.String("visibility"),
//   			Width: jsii.String("width"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoptions.html
//
type CfnTemplate_TableFieldOptionsProperty struct {
	// The order of the field IDs that are configured as field options for a table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoptions.html#cfn-quicksight-template-tablefieldoptions-order
	//
	Order *[]*string `field:"optional" json:"order" yaml:"order"`
	// The settings for the pinned columns of a table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoptions.html#cfn-quicksight-template-tablefieldoptions-pinnedfieldoptions
	//
	PinnedFieldOptions interface{} `field:"optional" json:"pinnedFieldOptions" yaml:"pinnedFieldOptions"`
	// The field options to be configured to a table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldoptions.html#cfn-quicksight-template-tablefieldoptions-selectedfieldoptions
	//
	SelectedFieldOptions interface{} `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
}

