package awsemrserverless


// The conﬁguration for an application to automatically start on job submission.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoStartConfigurationProperty := &autoStartConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnApplication_AutoStartConfigurationProperty struct {
	// Enables the application to automatically start on job submission.
	//
	// Defaults to true.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

