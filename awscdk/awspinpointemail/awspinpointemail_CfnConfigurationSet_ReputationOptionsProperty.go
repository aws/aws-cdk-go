package awspinpointemail


// Enable or disable collection of reputation metrics for emails that you send using this configuration set in the current AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reputationOptionsProperty := &reputationOptionsProperty{
//   	reputationMetricsEnabled: jsii.Boolean(false),
//   }
//
type CfnConfigurationSet_ReputationOptionsProperty struct {
	// If `true` , tracking of reputation metrics is enabled for the configuration set.
	//
	// If `false` , tracking of reputation metrics is disabled for the configuration set.
	ReputationMetricsEnabled interface{} `field:"optional" json:"reputationMetricsEnabled" yaml:"reputationMetricsEnabled"`
}

