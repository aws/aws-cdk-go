package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   borderSettingsProperty := &BorderSettingsProperty{
//   	BorderColor: jsii.String("borderColor"),
//   	BorderVisibility: jsii.String("borderVisibility"),
//   	BorderWidth: jsii.String("borderWidth"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-bordersettings.html
//
type CfnTemplatePropsMixin_BorderSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-bordersettings.html#cfn-quicksight-template-bordersettings-bordercolor
	//
	BorderColor *string `field:"optional" json:"borderColor" yaml:"borderColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-bordersettings.html#cfn-quicksight-template-bordersettings-bordervisibility
	//
	BorderVisibility *string `field:"optional" json:"borderVisibility" yaml:"borderVisibility"`
	// String based length that is composed of value and unit in px.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-bordersettings.html#cfn-quicksight-template-bordersettings-borderwidth
	//
	BorderWidth *string `field:"optional" json:"borderWidth" yaml:"borderWidth"`
}

