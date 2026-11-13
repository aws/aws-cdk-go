package awsdevopsagent


// A regex-based match condition.
//
// Passes when the value matches any pattern.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   patternFilterProperty := &PatternFilterProperty{
//   	Patterns: []*string{
//   		jsii.String("patterns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-patternfilter.html
//
type CfnAssociationPropsMixin_PatternFilterProperty struct {
	// Anchored full-match regex patterns, for example 'main' or 'release/.*'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-patternfilter.html#cfn-devopsagent-association-patternfilter-patterns
	//
	Patterns *[]*string `field:"optional" json:"patterns" yaml:"patterns"`
}

