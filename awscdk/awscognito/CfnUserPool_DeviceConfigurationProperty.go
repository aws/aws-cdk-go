package awscognito


// The device-remembering configuration for a user pool.
//
// A `API_DescribeUserPool` request returns a null value for this object when the user pool isn't configured to remember devices. When device remembering is active, you can remember a user's device with a `API_ConfirmDevice` API request. Additionally. when the property `DeviceOnlyRememberedOnUserPrompt` is `true` , you must follow `ConfirmDevice` with an `API_UpdateDeviceStatus` API request that sets the user's device to `remembered` or `not_remembered` .
//
// To sign in with a remembered device, include `DEVICE_KEY` in the authentication parameters in your user's `API_InitiateAuth` request. If your app doesn't include a `DEVICE_KEY` parameter, the `API_InitiateAuth` from Amazon Cognito includes newly-generated `DEVICE_KEY` and `DEVICE_GROUP_KEY` values under `NewDeviceMetadata` . Store these values to use in future device-authentication requests.
//
// > When you provide a value for any property of `DeviceConfiguration` , you activate the device remembering for the user pool.
// >
// > This data type is a request and response parameter of `API_CreateUserPool` and `API_UpdateUserPool` , and a response parameter of `API_DescribeUserPool` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceConfigurationProperty := &DeviceConfigurationProperty{
//   	ChallengeRequiredOnNewDevice: jsii.Boolean(false),
//   	DeviceOnlyRememberedOnUserPrompt: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html
//
type CfnUserPool_DeviceConfigurationProperty struct {
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

