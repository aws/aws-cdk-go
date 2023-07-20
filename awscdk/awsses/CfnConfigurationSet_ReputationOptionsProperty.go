package awsses


// Contains information about the reputation settings for a configuration set.
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
	// Describes whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch.
	//
	// If the value is `true` , reputation metrics are published. If the value is `false` , reputation metrics are not published. The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-reputationoptions.html#cfn-ses-configurationset-reputationoptions-reputationmetricsenabled
	//
	ReputationMetricsEnabled interface{} `field:"optional" json:"reputationMetricsEnabled" yaml:"reputationMetricsEnabled"`
}

