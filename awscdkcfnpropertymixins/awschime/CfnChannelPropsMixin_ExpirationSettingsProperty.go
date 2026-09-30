package awschime


// Settings that control the interval after which the channel is automatically deleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   expirationSettingsProperty := &ExpirationSettingsProperty{
//   	ExpirationCriterion: jsii.String("expirationCriterion"),
//   	ExpirationDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-expirationsettings.html
//
type CfnChannelPropsMixin_ExpirationSettingsProperty struct {
	// The condition the expiration period is measured from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-expirationsettings.html#cfn-chime-channel-expirationsettings-expirationcriterion
	//
	ExpirationCriterion *string `field:"optional" json:"expirationCriterion" yaml:"expirationCriterion"`
	// The period in days after which the system automatically deletes the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-expirationsettings.html#cfn-chime-channel-expirationsettings-expirationdays
	//
	ExpirationDays *float64 `field:"optional" json:"expirationDays" yaml:"expirationDays"`
}

