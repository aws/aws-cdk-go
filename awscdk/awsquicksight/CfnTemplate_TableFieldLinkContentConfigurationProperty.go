package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldLinkContentConfigurationProperty := &TableFieldLinkContentConfigurationProperty{
//   	CustomIconContent: &TableFieldCustomIconContentProperty{
//   		Icon: jsii.String("icon"),
//   	},
//   	CustomTextContent: &TableFieldCustomTextContentProperty{
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Value: jsii.String("value"),
//   	},
//   }
//
type CfnTemplate_TableFieldLinkContentConfigurationProperty struct {
	// `CfnTemplate.TableFieldLinkContentConfigurationProperty.CustomIconContent`.
	CustomIconContent interface{} `field:"optional" json:"customIconContent" yaml:"customIconContent"`
	// `CfnTemplate.TableFieldLinkContentConfigurationProperty.CustomTextContent`.
	CustomTextContent interface{} `field:"optional" json:"customTextContent" yaml:"customTextContent"`
}

