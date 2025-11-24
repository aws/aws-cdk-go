package mixinsawslex


// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sentimentAnalysisSettingsProperty := &SentimentAnalysisSettingsProperty{
//   	DetectSentiment: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-sentimentanalysissettings.html
//
type CfnBotAliasPropsMixin_SentimentAnalysisSettingsProperty struct {
	// Sets whether Amazon Lex uses Amazon Comprehend to detect the sentiment of user utterances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-sentimentanalysissettings.html#cfn-lex-botalias-sentimentanalysissettings-detectsentiment
	//
	DetectSentiment interface{} `field:"optional" json:"detectSentiment" yaml:"detectSentiment"`
}

