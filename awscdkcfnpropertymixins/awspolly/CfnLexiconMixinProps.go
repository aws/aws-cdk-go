package awspolly


// Properties for CfnLexiconPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnLexiconMixinProps := &CfnLexiconMixinProps{
//   	Content: jsii.String("content"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-polly-lexicon.html
//
type CfnLexiconMixinProps struct {
	// The content of the PLS lexicon in string format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-polly-lexicon.html#cfn-polly-lexicon-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The name of the lexicon.
	//
	// Must be an alphanumeric string up to 20 characters long.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-polly-lexicon.html#cfn-polly-lexicon-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

