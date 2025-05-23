package awsses


// An object containing additional settings for your VDM configuration as applicable to the Guardian.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardianAttributesProperty := &GuardianAttributesProperty{
//   	OptimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-vdmattributes-guardianattributes.html
//
type CfnVdmAttributes_GuardianAttributesProperty struct {
	// Specifies the status of your VDM optimized shared delivery. Can be one of the following:.
	//
	// - `ENABLED` – Amazon SES enables optimized shared delivery for your account.
	// - `DISABLED` – Amazon SES disables optimized shared delivery for your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-vdmattributes-guardianattributes.html#cfn-ses-vdmattributes-guardianattributes-optimizedshareddelivery
	//
	OptimizedSharedDelivery *string `field:"optional" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

