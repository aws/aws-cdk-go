package awsemrserverless


// The conﬁguration for an application to automatically stop after a certain amount of time being idle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoStopConfigurationProperty := &autoStopConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   	idleTimeoutMinutes: jsii.Number(123),
//   }
//
type CfnApplication_AutoStopConfigurationProperty struct {
	// Enables the application to automatically stop after a certain amount of time being idle.
	//
	// Defaults to true.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The amount of idle time in minutes after which your application will automatically stop. Defaults to 15 minutes.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 10080.
	IdleTimeoutMinutes *float64 `field:"optional" json:"idleTimeoutMinutes" yaml:"idleTimeoutMinutes"`
}

