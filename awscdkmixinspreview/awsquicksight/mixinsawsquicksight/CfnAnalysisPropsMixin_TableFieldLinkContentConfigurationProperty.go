package mixinsawsquicksight


// The URL content (text, icon) for the table link configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableFieldLinkContentConfigurationProperty := &TableFieldLinkContentConfigurationProperty{
//   	CustomIconContent: &TableFieldCustomIconContentProperty{
//   		Icon: jsii.String("icon"),
//   	},
//   	CustomTextContent: &TableFieldCustomTextContentProperty{
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablefieldlinkcontentconfiguration.html
//
type CfnAnalysisPropsMixin_TableFieldLinkContentConfigurationProperty struct {
	// The custom icon content for the table link content configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablefieldlinkcontentconfiguration.html#cfn-quicksight-analysis-tablefieldlinkcontentconfiguration-customiconcontent
	//
	CustomIconContent interface{} `field:"optional" json:"customIconContent" yaml:"customIconContent"`
	// The custom text content (value, font configuration) for the table link content configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablefieldlinkcontentconfiguration.html#cfn-quicksight-analysis-tablefieldlinkcontentconfiguration-customtextcontent
	//
	CustomTextContent interface{} `field:"optional" json:"customTextContent" yaml:"customTextContent"`
}

