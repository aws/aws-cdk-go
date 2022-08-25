package awsconnect


// Contains information about the phone configuration settings for a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPhoneConfigProperty := &userPhoneConfigProperty{
//   	phoneType: jsii.String("phoneType"),
//
//   	// the properties below are optional
//   	afterContactWorkTimeLimit: jsii.Number(123),
//   	autoAccept: jsii.Boolean(false),
//   	deskPhoneNumber: jsii.String("deskPhoneNumber"),
//   }
//
type CfnUser_UserPhoneConfigProperty struct {
	// The phone type.
	PhoneType *string `field:"required" json:"phoneType" yaml:"phoneType"`
	// The After Call Work (ACW) timeout setting, in seconds.
	AfterContactWorkTimeLimit *float64 `field:"optional" json:"afterContactWorkTimeLimit" yaml:"afterContactWorkTimeLimit"`
	// The Auto accept setting.
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// The phone number for the user's desk phone.
	DeskPhoneNumber *string `field:"optional" json:"deskPhoneNumber" yaml:"deskPhoneNumber"`
}

