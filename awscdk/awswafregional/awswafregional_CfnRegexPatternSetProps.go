package awswafregional


// Properties for defining a `CfnRegexPatternSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegexPatternSetProps := &cfnRegexPatternSetProps{
//   	name: jsii.String("name"),
//   	regexPatternStrings: []*string{
//   		jsii.String("regexPatternStrings"),
//   	},
//   }
//
type CfnRegexPatternSetProps struct {
	// A friendly name or description of the `RegexPatternSet` .
	//
	// You can't change `Name` after you create a `RegexPatternSet` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t` .
	RegexPatternStrings *[]*string `field:"required" json:"regexPatternStrings" yaml:"regexPatternStrings"`
}

