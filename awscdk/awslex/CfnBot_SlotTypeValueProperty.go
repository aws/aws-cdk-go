package awslex


// Each slot type can have a set of values.
//
// Each `SlotTypeValue` represents a value that the slot type can take.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotTypeValueProperty := &SlotTypeValueProperty{
//   	SampleValue: &SampleValueProperty{
//   		Value: jsii.String("value"),
//   	},
//
//   	// the properties below are optional
//   	Synonyms: []interface{}{
//   		&SampleValueProperty{
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottypevalue.html
//
type CfnBot_SlotTypeValueProperty struct {
	// The value of the slot type entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottypevalue.html#cfn-lex-bot-slottypevalue-samplevalue
	//
	SampleValue interface{} `field:"required" json:"sampleValue" yaml:"sampleValue"`
	// Additional values related to the slot type entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slottypevalue.html#cfn-lex-bot-slottypevalue-synonyms
	//
	Synonyms interface{} `field:"optional" json:"synonyms" yaml:"synonyms"`
}

