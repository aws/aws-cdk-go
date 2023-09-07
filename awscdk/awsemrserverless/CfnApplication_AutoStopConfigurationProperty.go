package awsemrserverless


// Configuration for Auto Stop of Application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoStopConfigurationProperty := &AutoStopConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	IdleTimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostopconfiguration.html
//
type CfnApplication_AutoStopConfigurationProperty struct {
	// If set to true, the Application will automatically stop after being idle.
	//
	// Defaults to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostopconfiguration.html#cfn-emrserverless-application-autostopconfiguration-enabled
	//
	// Default: - true.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The amount of time [in minutes] to wait before auto stopping the Application when idle.
	//
	// Defaults to 15 minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostopconfiguration.html#cfn-emrserverless-application-autostopconfiguration-idletimeoutminutes
	//
	IdleTimeoutMinutes *float64 `field:"optional" json:"idleTimeoutMinutes" yaml:"idleTimeoutMinutes"`
}

