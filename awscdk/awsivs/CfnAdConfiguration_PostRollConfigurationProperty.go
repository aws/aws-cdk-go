package awsivs


// Configuration for the post-roll ad break to use for this ad configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postRollConfigurationProperty := &PostRollConfigurationProperty{
//   	DurationSeconds: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-adconfiguration-postrollconfiguration.html
//
type CfnAdConfiguration_PostRollConfigurationProperty struct {
	// Duration of the post-roll ad break, in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-adconfiguration-postrollconfiguration.html#cfn-ivs-adconfiguration-postrollconfiguration-durationseconds
	//
	DurationSeconds *float64 `field:"required" json:"durationSeconds" yaml:"durationSeconds"`
	// Whether the post-roll ad configuration is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-adconfiguration-postrollconfiguration.html#cfn-ivs-adconfiguration-postrollconfiguration-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

