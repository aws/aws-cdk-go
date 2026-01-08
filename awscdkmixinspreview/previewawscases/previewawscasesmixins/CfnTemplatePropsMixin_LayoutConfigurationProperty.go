package previewawscasesmixins


// Object to store configuration of layouts associated to the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   layoutConfigurationProperty := &LayoutConfigurationProperty{
//   	DefaultLayout: jsii.String("defaultLayout"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-template-layoutconfiguration.html
//
type CfnTemplatePropsMixin_LayoutConfigurationProperty struct {
	// Unique identifier of a layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-template-layoutconfiguration.html#cfn-cases-template-layoutconfiguration-defaultlayout
	//
	DefaultLayout *string `field:"optional" json:"defaultLayout" yaml:"defaultLayout"`
}

