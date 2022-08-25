package awslex


// Provides configuration information for the AMAZON.KendraSearchIntent intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra index and returns documents from the index that match the user's utterance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kendraConfigurationProperty := &kendraConfigurationProperty{
//   	kendraIndex: jsii.String("kendraIndex"),
//
//   	// the properties below are optional
//   	queryFilterString: jsii.String("queryFilterString"),
//   	queryFilterStringEnabled: jsii.Boolean(false),
//   }
//
type CfnBot_KendraConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the AMAZON.KendraSearchIntent intent to search. The index must be in the same account and Region as the Amazon Lex bot.
	KendraIndex *string `field:"required" json:"kendraIndex" yaml:"kendraIndex"`
	// A query filter that Amazon Lex sends to Amazon Kendra to filter the response from a query.
	//
	// The filter is in the format defined by Amazon Kendra.
	QueryFilterString *string `field:"optional" json:"queryFilterString" yaml:"queryFilterString"`
	// Determines whether the AMAZON.KendraSearchIntent intent uses a custom query string to query the Amazon Kendra index.
	QueryFilterStringEnabled interface{} `field:"optional" json:"queryFilterStringEnabled" yaml:"queryFilterStringEnabled"`
}

