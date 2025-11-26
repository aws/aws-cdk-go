package previewawsdeadlinemixins


// Provides a script that runs as a worker is starting up that you can use to provide additional configuration for workers in your fleet.
//
// To remove a script from a fleet, use the [UpdateFleet](https://docs.aws.amazon.com/deadline-cloud/latest/APIReference/API_UpdateFleet.html) operation with the `hostConfiguration` `scriptBody` parameter set to an empty string ("").
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hostConfigurationProperty := &HostConfigurationProperty{
//   	ScriptBody: jsii.String("scriptBody"),
//   	ScriptTimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-hostconfiguration.html
//
type CfnFleetPropsMixin_HostConfigurationProperty struct {
	// The text of the script that runs as a worker is starting up that you can use to provide additional configuration for workers in your fleet.
	//
	// The script runs after a worker enters the `STARTING` state and before the worker processes tasks.
	//
	// For more information about using the script, see [Run scripts as an administrator to configure workers](https://docs.aws.amazon.com/deadline-cloud/latest/developerguide/smf-admin.html) in the *Deadline Cloud Developer Guide* .
	//
	// > The script runs as an administrative user ( `sudo root` on Linux, as an Administrator on Windows).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-hostconfiguration.html#cfn-deadline-fleet-hostconfiguration-scriptbody
	//
	ScriptBody *string `field:"optional" json:"scriptBody" yaml:"scriptBody"`
	// The maximum time that the host configuration can run.
	//
	// If the timeout expires, the worker enters the `NOT RESPONDING` state and shuts down. You are charged for the time that the worker is running the host configuration script.
	//
	// > You should configure your fleet for a maximum of one worker while testing your host configuration script to avoid starting additional workers.
	//
	// The default is 300 seconds (5 minutes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-hostconfiguration.html#cfn-deadline-fleet-hostconfiguration-scripttimeoutseconds
	//
	// Default: - 300.
	//
	ScriptTimeoutSeconds *float64 `field:"optional" json:"scriptTimeoutSeconds" yaml:"scriptTimeoutSeconds"`
}

