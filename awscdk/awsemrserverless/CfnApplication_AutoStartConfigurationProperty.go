package awsemrserverless


// Configuration for Auto Start of Application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoStartConfigurationProperty := &AutoStartConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostartconfiguration.html
//
type CfnApplication_AutoStartConfigurationProperty struct {
	// If set to true, the Application will automatically start.
	//
	// Defaults to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostartconfiguration.html#cfn-emrserverless-application-autostartconfiguration-enabled
	//
	// Default: - true.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

