package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldLinkConfigurationProperty := &TableFieldLinkConfigurationProperty{
//   	Content: &TableFieldLinkContentConfigurationProperty{
//   		CustomIconContent: &TableFieldCustomIconContentProperty{
//   			Icon: jsii.String("icon"),
//   		},
//   		CustomTextContent: &TableFieldCustomTextContentProperty{
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Target: jsii.String("target"),
//   }
//
type CfnDashboard_TableFieldLinkConfigurationProperty struct {
	// `CfnDashboard.TableFieldLinkConfigurationProperty.Content`.
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// `CfnDashboard.TableFieldLinkConfigurationProperty.Target`.
	Target *string `field:"required" json:"target" yaml:"target"`
}

