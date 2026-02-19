package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deepgramSpeechModelConfigProperty := &DeepgramSpeechModelConfigProperty{
//   	ApiTokenSecretArn: jsii.String("apiTokenSecretArn"),
//
//   	// the properties below are optional
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-deepgramspeechmodelconfig.html
//
type CfnBot_DeepgramSpeechModelConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-deepgramspeechmodelconfig.html#cfn-lex-bot-deepgramspeechmodelconfig-apitokensecretarn
	//
	ApiTokenSecretArn *string `field:"required" json:"apiTokenSecretArn" yaml:"apiTokenSecretArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-deepgramspeechmodelconfig.html#cfn-lex-bot-deepgramspeechmodelconfig-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

