package awslex


// Contains details about the configuration of the Amazon Kendra index used for the AMAZON.QnAIntent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   qnAKendraConfigurationProperty := &QnAKendraConfigurationProperty{
//   	ExactResponse: jsii.Boolean(false),
//   	KendraIndex: jsii.String("kendraIndex"),
//   	QueryFilterStringEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	QueryFilterString: jsii.String("queryFilterString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnakendraconfiguration.html
//
type CfnBot_QnAKendraConfigurationProperty struct {
	// Specifies whether to return an exact response from the Amazon Kendra index or to let the Amazon Bedrock model you select generate a response based on the results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnakendraconfiguration.html#cfn-lex-bot-qnakendraconfiguration-exactresponse
	//
	ExactResponse interface{} `field:"required" json:"exactResponse" yaml:"exactResponse"`
	// The ARN of the Amazon Kendra index to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnakendraconfiguration.html#cfn-lex-bot-qnakendraconfiguration-kendraindex
	//
	KendraIndex *string `field:"required" json:"kendraIndex" yaml:"kendraIndex"`
	// Specifies whether to enable an Amazon Kendra filter string or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnakendraconfiguration.html#cfn-lex-bot-qnakendraconfiguration-queryfilterstringenabled
	//
	QueryFilterStringEnabled interface{} `field:"required" json:"queryFilterStringEnabled" yaml:"queryFilterStringEnabled"`
	// Contains the Amazon Kendra filter string to use if enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnakendraconfiguration.html#cfn-lex-bot-qnakendraconfiguration-queryfilterstring
	//
	QueryFilterString *string `field:"optional" json:"queryFilterString" yaml:"queryFilterString"`
}

