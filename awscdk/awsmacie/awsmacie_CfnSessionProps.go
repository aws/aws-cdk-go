package awsmacie


// Properties for defining a `CfnSession`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSessionProps := &cfnSessionProps{
//   	findingPublishingFrequency: jsii.String("findingPublishingFrequency"),
//   	status: jsii.String("status"),
//   }
//
type CfnSessionProps struct {
	// The frequency with which Amazon Macie publishes updates to policy findings for an account.
	//
	// This includes publishing updates to AWS Security Hub and Amazon EventBridge (formerly called Amazon CloudWatch Events ). Valid values are:
	//
	// - FIFTEEN_MINUTES
	// - ONE_HOUR
	// - SIX_HOURS.
	FindingPublishingFrequency *string `field:"optional" json:"findingPublishingFrequency" yaml:"findingPublishingFrequency"`
	// The `MacieStatus` of the `Session` .
	//
	// Valid values include `ENABLED` and `PAUSED` .
	Status *string `field:"optional" json:"status" yaml:"status"`
}

