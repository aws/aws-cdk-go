package assertions


// Options to configure template parsing behavior, such as disregarding circular dependencies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateParsingOptions := &templateParsingOptions{
//   	skipCyclicalDependenciesCheck: jsii.Boolean(false),
//   }
//
type TemplateParsingOptions struct {
	// If set to true, will skip checking for cyclical / circular dependencies.
	//
	// Should be set to false other than for
	// templates that are valid despite containing cycles, such as unprocessed transform stacks.
	SkipCyclicalDependenciesCheck *bool `field:"optional" json:"skipCyclicalDependenciesCheck" yaml:"skipCyclicalDependenciesCheck"`
}

