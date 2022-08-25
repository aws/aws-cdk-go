package awslex


// Provides a regular expression used to validate the value of a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotValueRegexFilterProperty := &slotValueRegexFilterProperty{
//   	pattern: jsii.String("pattern"),
//   }
//
type CfnBot_SlotValueRegexFilterProperty struct {
	// A regular expression used to validate the value of a slot.
	//
	// Use a standard regular expression. Amazon Lex supports the following characters in the regular expression:
	//
	// - A-Z, a-z
	// - 0-9
	// - Unicode characters ("\ u<Unicode>")
	//
	// Represent Unicode characters with four digits, for example "]u0041" or "\ u005A".
	//
	// The following regular expression operators are not supported:
	//
	// - Infinite repeaters: *, +, or {x,} with no upper bound
	// - Wild card (.)
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
}

