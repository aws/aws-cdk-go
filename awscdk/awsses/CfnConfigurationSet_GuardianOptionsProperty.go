package awsses


// An object containing additional settings for your VDM configuration as applicable to the Guardian.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardianOptionsProperty := &GuardianOptionsProperty{
//   	OptimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-guardianoptions.html
//
type CfnConfigurationSet_GuardianOptionsProperty struct {
	// Specifies the status of your VDM optimized shared delivery. Can be one of the following:.
	//
	// - `ENABLED` – Amazon SES enables optimized shared delivery for the configuration set.
	// - `DISABLED` – Amazon SES disables optimized shared delivery for the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-guardianoptions.html#cfn-ses-configurationset-guardianoptions-optimizedshareddelivery
	//
	OptimizedSharedDelivery *string `field:"required" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

