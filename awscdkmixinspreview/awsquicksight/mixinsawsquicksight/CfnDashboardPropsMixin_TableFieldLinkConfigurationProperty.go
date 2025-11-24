package mixinsawsquicksight


// The link configuration of a table field URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   				FontFamily: jsii.String("fontFamily"),
//   				FontSize: &FontSizeProperty{
//   					Absolute: jsii.String("absolute"),
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Target: jsii.String("target"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablefieldlinkconfiguration.html
//
type CfnDashboardPropsMixin_TableFieldLinkConfigurationProperty struct {
	// The URL content (text, icon) for the table link configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablefieldlinkconfiguration.html#cfn-quicksight-dashboard-tablefieldlinkconfiguration-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// The URL target (new tab, new window, same tab) for the table link configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablefieldlinkconfiguration.html#cfn-quicksight-dashboard-tablefieldlinkconfiguration-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
}

