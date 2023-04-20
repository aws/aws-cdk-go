package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldOptionsProperty := &TableFieldOptionsProperty{
//   	Order: []*string{
//   		jsii.String("order"),
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
//   								FontSize: &FontSizeProperty{
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
type CfnDashboard_TableFieldOptionsProperty struct {
	// `CfnDashboard.TableFieldOptionsProperty.Order`.
	Order *[]*string `field:"optional" json:"order" yaml:"order"`
	// `CfnDashboard.TableFieldOptionsProperty.SelectedFieldOptions`.
	SelectedFieldOptions interface{} `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
}

