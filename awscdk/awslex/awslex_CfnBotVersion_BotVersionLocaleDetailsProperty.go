package awslex


// The version of a bot used for a bot locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botVersionLocaleDetailsProperty := &botVersionLocaleDetailsProperty{
//   	sourceBotVersion: jsii.String("sourceBotVersion"),
//   }
//
type CfnBotVersion_BotVersionLocaleDetailsProperty struct {
	// The version of a bot used for a bot locale.
	SourceBotVersion *string `field:"required" json:"sourceBotVersion" yaml:"sourceBotVersion"`
}

