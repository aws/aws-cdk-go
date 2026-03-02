package previewawscloudwatchmixins


// Targets to be muted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   muteTargetsProperty := &MuteTargetsProperty{
//   	AlarmNames: []*string{
//   		jsii.String("alarmNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarmmuterule-mutetargets.html
//
type CfnAlarmMuteRulePropsMixin_MuteTargetsProperty struct {
	// The alarm names to be mute by the AlarmMuteRule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarmmuterule-mutetargets.html#cfn-cloudwatch-alarmmuterule-mutetargets-alarmnames
	//
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
}

