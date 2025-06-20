package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtGroupSettingsProperty := &SrtGroupSettingsProperty{
//   	InputLossAction: jsii.String("inputLossAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtgroupsettings.html
//
type CfnChannel_SrtGroupSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtgroupsettings.html#cfn-medialive-channel-srtgroupsettings-inputlossaction
	//
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
}

