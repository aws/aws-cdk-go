package awsmacie


// Properties for defining a `CfnCustomDataIdentifier`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomDataIdentifierProps := &cfnCustomDataIdentifierProps{
//   	name: jsii.String("name"),
//   	regex: jsii.String("regex"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	ignoreWords: []*string{
//   		jsii.String("ignoreWords"),
//   	},
//   	keywords: []*string{
//   		jsii.String("keywords"),
//   	},
//   	maximumMatchDistance: jsii.Number(123),
//   }
//
type CfnCustomDataIdentifierProps struct {
	// A custom name for the custom data identifier. The name can contain as many as 128 characters.
	//
	// We strongly recommend that you avoid including any sensitive data in the name of a custom data identifier. Other users of your account might be able to see the identifier's name, depending on the actions that they're allowed to perform in Amazon Macie .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The regular expression ( *regex* ) that defines the pattern to match.
	//
	// The expression can contain as many as 512 characters.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// The description of the custom data identifier.
	//
	// The description can contain as many as 512 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array that lists specific character sequences (ignore words) to exclude from the results.
	//
	// If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4-90 characters. Ignore words are case sensitive.
	IgnoreWords *[]*string `field:"optional" json:"ignoreWords" yaml:"ignoreWords"`
	// An array that lists specific character sequences (keywords), one of which must be within proximity ( `MaximumMatchDistance` ) of the regular expression to match.
	//
	// The array can contain as many as 50 keywords. Each keyword can contain 3-90 characters. Keywords aren't case sensitive.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the `Keywords` array.
	//
	// Amazon Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1-300 characters. The default value is 50.
	MaximumMatchDistance *float64 `field:"optional" json:"maximumMatchDistance" yaml:"maximumMatchDistance"`
}

