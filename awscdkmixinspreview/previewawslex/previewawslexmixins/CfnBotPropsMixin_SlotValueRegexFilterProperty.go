package previewawslexmixins


// Provides a regular expression used to validate the value of a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slotValueRegexFilterProperty := &SlotValueRegexFilterProperty{
//   	Pattern: jsii.String("pattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueregexfilter.html
//
type CfnBotPropsMixin_SlotValueRegexFilterProperty struct {
	// A regular expression used to validate the value of a slot.
	//
	// Use a standard regular expression. Amazon Lex supports the following characters in the regular expression:
	//
	// - A-Z, a-z
	// - 0-9
	// - Unicode characters ("\⁠u<Unicode>")
	//
	// Represent Unicode characters with four digits, for example "\⁠u0041" or "\⁠u005A".
	//
	// The following regular expression operators are not supported:
	//
	// - Infinite repeaters: *, +, or {x,} with no upper bound.
	// - Wild card (.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueregexfilter.html#cfn-lex-bot-slotvalueregexfilter-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

