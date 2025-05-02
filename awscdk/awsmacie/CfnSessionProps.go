package awsmacie


// Properties for defining a `CfnSession`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSessionProps := &CfnSessionProps{
//   	FindingPublishingFrequency: jsii.String("findingPublishingFrequency"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html
//
type CfnSessionProps struct {
	// Specifies how often Amazon Macie publishes updates to policy findings for the account.
	//
	// This includes publishing updates to AWS Security Hub and Amazon EventBridge (formerly Amazon CloudWatch Events ). Valid values are:
	//
	// - FIFTEEN_MINUTES
	// - ONE_HOUR
	// - SIX_HOURS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-findingpublishingfrequency
	//
	// Default: - "SIX_HOURS".
	//
	FindingPublishingFrequency *string `field:"optional" json:"findingPublishingFrequency" yaml:"findingPublishingFrequency"`
	// The status of Amazon Macie for the account.
	//
	// Valid values are: `ENABLED` , start or resume Macie activities for the account; and, `PAUSED` , suspend Macie activities for the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-status
	//
	// Default: - "ENABLED".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

