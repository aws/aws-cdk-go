package awsconnect


// Phone Number configuration per channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   phoneNumberConfigProperty := &PhoneNumberConfigProperty{
//   	Channel: jsii.String("channel"),
//   	PhoneType: jsii.String("phoneType"),
//
//   	// the properties below are optional
//   	PhoneNumber: jsii.String("phoneNumber"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-phonenumberconfig.html
//
type CfnUser_PhoneNumberConfigProperty struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-phonenumberconfig.html#cfn-connect-user-phonenumberconfig-channel
	//
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The phone type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-phonenumberconfig.html#cfn-connect-user-phonenumberconfig-phonetype
	//
	PhoneType *string `field:"required" json:"phoneType" yaml:"phoneType"`
	// The phone number for the user's desk phone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-phonenumberconfig.html#cfn-connect-user-phonenumberconfig-phonenumber
	//
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

