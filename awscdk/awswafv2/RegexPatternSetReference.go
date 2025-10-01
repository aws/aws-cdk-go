package awswafv2


// A reference to a RegexPatternSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regexPatternSetReference := &RegexPatternSetReference{
//   	RegexPatternSetArn: jsii.String("regexPatternSetArn"),
//   	RegexPatternSetId: jsii.String("regexPatternSetId"),
//   	RegexPatternSetName: jsii.String("regexPatternSetName"),
//   	Scope: jsii.String("scope"),
//   }
//
type RegexPatternSetReference struct {
	// The ARN of the RegexPatternSet resource.
	RegexPatternSetArn *string `field:"required" json:"regexPatternSetArn" yaml:"regexPatternSetArn"`
	// The Id of the RegexPatternSet resource.
	RegexPatternSetId *string `field:"required" json:"regexPatternSetId" yaml:"regexPatternSetId"`
	// The Name of the RegexPatternSet resource.
	RegexPatternSetName *string `field:"required" json:"regexPatternSetName" yaml:"regexPatternSetName"`
	// The Scope of the RegexPatternSet resource.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

