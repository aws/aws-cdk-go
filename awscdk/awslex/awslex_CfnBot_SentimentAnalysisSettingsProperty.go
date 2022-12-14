package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sentimentAnalysisSettingsProperty := &sentimentAnalysisSettingsProperty{
//   	detectSentiment: jsii.Boolean(false),
//   }
//
type CfnBot_SentimentAnalysisSettingsProperty struct {
	// `CfnBot.SentimentAnalysisSettingsProperty.DetectSentiment`.
	DetectSentiment interface{} `field:"required" json:"detectSentiment" yaml:"detectSentiment"`
}

