package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   perceptualFilterSettingsProperty := &PerceptualFilterSettingsProperty{
//   	Strength: jsii.String("strength"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-perceptualfiltersettings.html
//
type CfnChannelPropsMixin_PerceptualFilterSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-perceptualfiltersettings.html#cfn-medialive-channel-perceptualfiltersettings-strength
	//
	Strength *string `field:"optional" json:"strength" yaml:"strength"`
}

