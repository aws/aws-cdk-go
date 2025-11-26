package previewawsconnectmixins


// Contains information about a phone number for a quick connect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   phoneNumberQuickConnectConfigProperty := &PhoneNumberQuickConnectConfigProperty{
//   	PhoneNumber: jsii.String("phoneNumber"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-phonenumberquickconnectconfig.html
//
type CfnQuickConnectPropsMixin_PhoneNumberQuickConnectConfigProperty struct {
	// The phone number in E.164 format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-phonenumberquickconnectconfig.html#cfn-connect-quickconnect-phonenumberquickconnectconfig-phonenumber
	//
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

