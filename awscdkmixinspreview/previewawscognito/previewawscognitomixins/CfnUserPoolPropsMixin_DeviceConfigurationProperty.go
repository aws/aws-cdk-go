package previewawscognitomixins


// The device-remembering configuration for a user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deviceConfigurationProperty := &DeviceConfigurationProperty{
//   	ChallengeRequiredOnNewDevice: jsii.Boolean(false),
//   	DeviceOnlyRememberedOnUserPrompt: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html
//
type CfnUserPoolPropsMixin_DeviceConfigurationProperty struct {
	// When true, a remembered device can sign in with device authentication instead of SMS and time-based one-time password (TOTP) factors for multi-factor authentication (MFA).
	//
	// > Whether or not `ChallengeRequiredOnNewDevice` is true, users who sign in with devices that have not been confirmed or remembered must still provide a second factor in a user pool that requires MFA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html#cfn-cognito-userpool-deviceconfiguration-challengerequiredonnewdevice
	//
	ChallengeRequiredOnNewDevice interface{} `field:"optional" json:"challengeRequiredOnNewDevice" yaml:"challengeRequiredOnNewDevice"`
	// When true, Amazon Cognito doesn't automatically remember a user's device when your app sends a `ConfirmDevice` API request.
	//
	// In your app, create a prompt for your user to choose whether they want to remember their device. Return the user's choice in an `UpdateDeviceStatus` API request.
	//
	// When `DeviceOnlyRememberedOnUserPrompt` is `false` , Amazon Cognito immediately remembers devices that you register in a `ConfirmDevice` API request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html#cfn-cognito-userpool-deviceconfiguration-deviceonlyrememberedonuserprompt
	//
	DeviceOnlyRememberedOnUserPrompt interface{} `field:"optional" json:"deviceOnlyRememberedOnUserPrompt" yaml:"deviceOnlyRememberedOnUserPrompt"`
}

