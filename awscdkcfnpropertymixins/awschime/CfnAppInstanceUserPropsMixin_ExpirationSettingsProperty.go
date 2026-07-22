package awschime


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstanceuser-expirationsettings.html
//
type CfnAppInstanceUserPropsMixin_ExpirationSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstanceuser-expirationsettings.html#cfn-chime-appinstanceuser-expirationsettings-expirationcriterion
	//
	ExpirationCriterion *string `field:"optional" json:"expirationCriterion" yaml:"expirationCriterion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-appinstanceuser-expirationsettings.html#cfn-chime-appinstanceuser-expirationsettings-expirationdays
	//
	ExpirationDays *float64 `field:"optional" json:"expirationDays" yaml:"expirationDays"`
}

