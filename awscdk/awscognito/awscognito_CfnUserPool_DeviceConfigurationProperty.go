package awscognito


// The device tracking configuration for a user pool. A user pool with device tracking deactivated returns a null value.
//
// > When you provide values for any DeviceConfiguration field, you activate device tracking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceConfigurationProperty := &deviceConfigurationProperty{
//   	challengeRequiredOnNewDevice: jsii.Boolean(false),
//   	deviceOnlyRememberedOnUserPrompt: jsii.Boolean(false),
//   }
//
type CfnUserPool_DeviceConfigurationProperty struct {
	// When true, device authentication can replace SMS and time-based one-time password (TOTP) factors for multi-factor authentication (MFA).
	//
	// > Users that sign in with devices that have not been confirmed or remembered will still have to provide a second factor, whether or not ChallengeRequiredOnNewDevice is true, when your user pool requires MFA.
	ChallengeRequiredOnNewDevice interface{} `field:"optional" json:"challengeRequiredOnNewDevice" yaml:"challengeRequiredOnNewDevice"`
	// When true, users can opt in to remembering their device.
	//
	// Your app code must use callback functions to return the user's choice.
	DeviceOnlyRememberedOnUserPrompt interface{} `field:"optional" json:"deviceOnlyRememberedOnUserPrompt" yaml:"deviceOnlyRememberedOnUserPrompt"`
}

