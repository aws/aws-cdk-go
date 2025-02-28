package awslex


// Provides configuration information for the `AMAZON.KendraSearchIntent` intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra index and returns documents from the index that match the user's utterance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kendraConfigurationProperty := &KendraConfigurationProperty{
//   	KendraIndex: jsii.String("kendraIndex"),
//
//   	// the properties below are optional
//   	QueryFilterString: jsii.String("queryFilterString"),
//   	QueryFilterStringEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-kendraconfiguration.html
//
type CfnBot_KendraConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the `AMAZON.KendraSearchIntent` intent to search. The index must be in the same account and Region as the Amazon Lex bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-kendraconfiguration.html#cfn-lex-bot-kendraconfiguration-kendraindex
	//
	KendraIndex *string `field:"required" json:"kendraIndex" yaml:"kendraIndex"`
	// A query filter that Amazon Lex sends to Amazon Kendra to filter the response from a query.
	//
	// The filter is in the format defined by Amazon Kendra. For more information, see [Filtering queries](https://docs.aws.amazon.com/kendra/latest/dg/filtering.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-kendraconfiguration.html#cfn-lex-bot-kendraconfiguration-queryfilterstring
	//
	QueryFilterString *string `field:"optional" json:"queryFilterString" yaml:"queryFilterString"`
	// Determines whether the `AMAZON.KendraSearchIntent` intent uses a custom query string to query the Amazon Kendra index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-kendraconfiguration.html#cfn-lex-bot-kendraconfiguration-queryfilterstringenabled
	//
	QueryFilterStringEnabled interface{} `field:"optional" json:"queryFilterStringEnabled" yaml:"queryFilterStringEnabled"`
}

