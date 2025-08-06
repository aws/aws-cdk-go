package awsses


// Enable or disable collection of reputation metrics for emails that you send using this configuration set in the current AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reputationOptionsProperty := &ReputationOptionsProperty{
//   	ReputationMetricsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-reputationoptions.html
//
type CfnConfigurationSet_ReputationOptionsProperty struct {
	// If `true` , tracking of reputation metrics is enabled for the configuration set.
	//
	// If `false` , tracking of reputation metrics is disabled for the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-reputationoptions.html#cfn-ses-configurationset-reputationoptions-reputationmetricsenabled
	//
	ReputationMetricsEnabled interface{} `field:"optional" json:"reputationMetricsEnabled" yaml:"reputationMetricsEnabled"`
}

