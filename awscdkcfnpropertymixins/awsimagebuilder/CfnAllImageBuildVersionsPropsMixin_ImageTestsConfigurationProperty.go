package awsimagebuilder


// The image tests configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   imageTestsConfigurationProperty := &ImageTestsConfigurationProperty{
//   	ImageTestsEnabled: jsii.Boolean(false),
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allimagebuildversions-imagetestsconfiguration.html
//
type CfnAllImageBuildVersionsPropsMixin_ImageTestsConfigurationProperty struct {
	// Determines if tests should run after building the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allimagebuildversions-imagetestsconfiguration.html#cfn-imagebuilder-allimagebuildversions-imagetestsconfiguration-imagetestsenabled
	//
	ImageTestsEnabled interface{} `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// The maximum time in minutes that tests are permitted to run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allimagebuildversions-imagetestsconfiguration.html#cfn-imagebuilder-allimagebuildversions-imagetestsconfiguration-timeoutminutes
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

