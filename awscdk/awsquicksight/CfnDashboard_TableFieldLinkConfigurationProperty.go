package awsquicksight


// The link configuration of a table field URL.
//
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
	// The URL content (text, icon) for the table link configuration.
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// The URL target (new tab, new window, same tab) for the table link configuration.
	Target *string `field:"required" json:"target" yaml:"target"`
}

