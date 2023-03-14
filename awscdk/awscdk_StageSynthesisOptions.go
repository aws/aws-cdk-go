// An experiment to bundle the entire CDK into a single module
package awscdk


// Options for assembly synthesis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   stageSynthesisOptions := &stageSynthesisOptions{
//   	force: jsii.Boolean(false),
//   	skipValidation: jsii.Boolean(false),
//   	validateOnSynthesis: jsii.Boolean(false),
//   }
//
// Experimental.
type StageSynthesisOptions struct {
	// Force a re-synth, even if the stage has already been synthesized.
	//
	// This is used by tests to allow for incremental verification of the output.
	// Do not use in production.
	// Experimental.
	Force *bool `field:"optional" json:"force" yaml:"force"`
	// Should we skip construct validation.
	// Experimental.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// Whether the stack should be validated after synthesis to check for error metadata.
	// Experimental.
	ValidateOnSynthesis *bool `field:"optional" json:"validateOnSynthesis" yaml:"validateOnSynthesis"`
}

