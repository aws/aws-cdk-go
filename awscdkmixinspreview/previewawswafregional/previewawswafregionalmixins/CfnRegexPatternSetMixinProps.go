package previewawswafregionalmixins


// Properties for CfnRegexPatternSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRegexPatternSetMixinProps := &CfnRegexPatternSetMixinProps{
//   	Name: jsii.String("name"),
//   	RegexPatternStrings: []*string{
//   		jsii.String("regexPatternStrings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html
//
type CfnRegexPatternSetMixinProps struct {
	// A friendly name or description of the `RegexPatternSet` .
	//
	// You can't change `Name` after you create a `RegexPatternSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-regexpatternstrings
	//
	RegexPatternStrings *[]*string `field:"optional" json:"regexPatternStrings" yaml:"regexPatternStrings"`
}

