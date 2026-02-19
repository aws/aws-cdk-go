package previewawsconnectmixins


// Phone Number configuration per channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   phoneNumberConfigProperty := &PhoneNumberConfigProperty{
//   	Channel: jsii.String("channel"),
//   	PhoneNumber: jsii.String("phoneNumber"),
//   	PhoneType: jsii.String("phoneType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-phonenumberconfig.html
//
type CfnUserPropsMixin_PhoneNumberConfigProperty struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-phonenumberconfig.html#cfn-connect-user-phonenumberconfig-channel
	//
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The phone number for the user's desk phone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-phonenumberconfig.html#cfn-connect-user-phonenumberconfig-phonenumber
	//
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The phone type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-phonenumberconfig.html#cfn-connect-user-phonenumberconfig-phonetype
	//
	PhoneType *string `field:"optional" json:"phoneType" yaml:"phoneType"`
}

