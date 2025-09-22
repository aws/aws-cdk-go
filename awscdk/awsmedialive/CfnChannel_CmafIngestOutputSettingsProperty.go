package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafIngestOutputSettingsProperty := &CmafIngestOutputSettingsProperty{
//   	NameModifier: jsii.String("nameModifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestoutputsettings.html
//
type CfnChannel_CmafIngestOutputSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestoutputsettings.html#cfn-medialive-channel-cmafingestoutputsettings-namemodifier
	//
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

