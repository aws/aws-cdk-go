package awswafregional


// A reference to a RegexPatternSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regexPatternSetReference := &RegexPatternSetReference{
//   	RegexPatternSetId: jsii.String("regexPatternSetId"),
//   }
//
type RegexPatternSetReference struct {
	// The Id of the RegexPatternSet resource.
	RegexPatternSetId *string `field:"required" json:"regexPatternSetId" yaml:"regexPatternSetId"`
}

