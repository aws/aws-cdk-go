package awsquicksight


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
//   						FontSize: &FontSizeProperty{
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
type CfnTemplate_TableFieldOptionProperty struct {
	// `CfnTemplate.TableFieldOptionProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnTemplate.TableFieldOptionProperty.CustomLabel`.
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// `CfnTemplate.TableFieldOptionProperty.URLStyling`.
	UrlStyling interface{} `field:"optional" json:"urlStyling" yaml:"urlStyling"`
	// `CfnTemplate.TableFieldOptionProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// `CfnTemplate.TableFieldOptionProperty.Width`.
	Width *string `field:"optional" json:"width" yaml:"width"`
}

