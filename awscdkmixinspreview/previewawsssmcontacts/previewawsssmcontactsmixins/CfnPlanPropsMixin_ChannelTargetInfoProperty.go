package previewawsssmcontactsmixins


// Information about the contact channel that Incident Manager uses to engage the contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   channelTargetInfoProperty := &ChannelTargetInfoProperty{
//   	ChannelId: jsii.String("channelId"),
//   	RetryIntervalInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-plan-channeltargetinfo.html
//
type CfnPlanPropsMixin_ChannelTargetInfoProperty struct {
	// The Amazon Resource Name (ARN) of the contact channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-plan-channeltargetinfo.html#cfn-ssmcontacts-plan-channeltargetinfo-channelid
	//
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// The number of minutes to wait before retrying to send engagement if the engagement initially failed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-plan-channeltargetinfo.html#cfn-ssmcontacts-plan-channeltargetinfo-retryintervalinminutes
	//
	RetryIntervalInMinutes *float64 `field:"optional" json:"retryIntervalInMinutes" yaml:"retryIntervalInMinutes"`
}

