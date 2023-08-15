package awscognito


// Device tracking settings.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	DeviceTracking: &DeviceTracking{
//   		ChallengeRequiredOnNewDevice: jsii.Boolean(true),
//   		DeviceOnlyRememberedOnUserPrompt: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
//
type DeviceTracking struct {
	// Indicates whether a challenge is required on a new device.
	//
	// Only applicable to a new device.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
	//
	// Default: false.
	//
	ChallengeRequiredOnNewDevice *bool `field:"required" json:"challengeRequiredOnNewDevice" yaml:"challengeRequiredOnNewDevice"`
	// If true, a device is only remembered on user prompt.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
	//
	// Default: false.
	//
	DeviceOnlyRememberedOnUserPrompt *bool `field:"required" json:"deviceOnlyRememberedOnUserPrompt" yaml:"deviceOnlyRememberedOnUserPrompt"`
}

