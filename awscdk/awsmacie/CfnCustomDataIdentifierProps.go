package awsmacie


// Properties for defining a `CfnCustomDataIdentifier`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomDataIdentifierProps := &CfnCustomDataIdentifierProps{
//   	Name: jsii.String("name"),
//   	Regex: jsii.String("regex"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	IgnoreWords: []*string{
//   		jsii.String("ignoreWords"),
//   	},
//   	Keywords: []*string{
//   		jsii.String("keywords"),
//   	},
//   	MaximumMatchDistance: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html
//
type CfnCustomDataIdentifierProps struct {
	// A custom name for the custom data identifier. The name can contain 1-128 characters.
	//
	// Avoid including sensitive data in the name of a custom data identifier. Users of the account might be able to see the name, depending on the actions that they're allowed to perform in Amazon Macie .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The regular expression ( *regex* ) that defines the text pattern to match.
	//
	// The expression can contain 1-512 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-regex
	//
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// A custom description of the custom data identifier. The description can contain 1-512 characters.
	//
	// Avoid including sensitive data in the description. Users of the account might be able to see the description, depending on the actions that they're allowed to perform in Amazon Macie .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of character sequences ( *ignore words* ) to exclude from the results.
	//
	// If text matches the regular expression ( `Regex` ) but it contains a string in this array, Amazon Macie ignores the text and doesn't include it in the results.
	//
	// The array can contain 1-10 ignore words. Each ignore word can contain 4-90 UTF-8 characters. Ignore words are case sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-ignorewords
	//
	IgnoreWords *[]*string `field:"optional" json:"ignoreWords" yaml:"ignoreWords"`
	// An array of character sequences ( *keywords* ), one of which must precede and be in proximity ( `MaximumMatchDistance` ) of the regular expression ( `Regex` ) to match.
	//
	// The array can contain 1-50 keywords. Each keyword can contain 3-90 UTF-8 characters. Keywords aren't case sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-keywords
	//
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// The maximum number of characters that can exist between the end of at least one complete character sequence specified by the `Keywords` array and the end of text that matches the regular expression ( `Regex` ).
	//
	// If a complete keyword precedes all the text that matches the regular expression and the keyword is within the specified distance, Amazon Macie includes the result.
	//
	// The distance can be 1-300 characters. The default value is 50.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-maximummatchdistance
	//
	MaximumMatchDistance *float64 `field:"optional" json:"maximumMatchDistance" yaml:"maximumMatchDistance"`
}

