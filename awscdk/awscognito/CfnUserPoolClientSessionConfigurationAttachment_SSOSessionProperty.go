package awscognito


// Managed login single sign-on session settings for the app client.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSOSessionProperty := &SSOSessionProperty{
//   	MaxValidityMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclientsessionconfigurationattachment-ssosession.html
//
type CfnUserPoolClientSessionConfigurationAttachment_SSOSessionProperty struct {
	// The maximum time in minutes since the last authentication during which this app client can reuse an existing managed login SSO session without asking the user to sign in again.
	//
	// 0 means the app client always asks the user to sign in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclientsessionconfigurationattachment-ssosession.html#cfn-cognito-userpoolclientsessionconfigurationattachment-ssosession-maxvalidityminutes
	//
	MaxValidityMinutes *float64 `field:"required" json:"maxValidityMinutes" yaml:"maxValidityMinutes"`
}

