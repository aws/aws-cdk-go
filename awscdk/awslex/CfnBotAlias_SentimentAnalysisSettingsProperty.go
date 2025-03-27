package awslex


// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sentimentAnalysisSettingsProperty := &SentimentAnalysisSettingsProperty{
//   	DetectSentiment: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-sentimentanalysissettings.html
//
type CfnBotAlias_SentimentAnalysisSettingsProperty struct {
	// Sets whether Amazon Lex uses Amazon Comprehend to detect the sentiment of user utterances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-sentimentanalysissettings.html#cfn-lex-botalias-sentimentanalysissettings-detectsentiment
	//
	DetectSentiment interface{} `field:"required" json:"detectSentiment" yaml:"detectSentiment"`
}

