// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
)

// Options for synthesis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   synthesisOptions := &synthesisOptions{
//   	outdir: jsii.String("outdir"),
//   	runtimeInfo: &runtimeInfo{
//   		libraries: map[string]*string{
//   			"librariesKey": jsii.String("libraries"),
//   		},
//   	},
//   	skipValidation: jsii.Boolean(false),
//   	validateOnSynthesis: jsii.Boolean(false),
//   }
//
// Deprecated: use `app.synth()` or `stage.synth()` instead
type SynthesisOptions struct {
	// Include the specified runtime information (module versions) in manifest.
	// Deprecated: All template modifications that should result from this should
	// have already been inserted into the template.
	RuntimeInfo *cxapi.RuntimeInfo `field:"optional" json:"runtimeInfo" yaml:"runtimeInfo"`
	// The output directory into which to synthesize the cloud assembly.
	// Deprecated: use `app.synth()` or `stage.synth()` instead
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Whether synthesis should skip the validation phase.
	// Deprecated: use `app.synth()` or `stage.synth()` instead
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// Whether the stack should be validated after synthesis to check for error metadata.
	// Deprecated: use `app.synth()` or `stage.synth()` instead
	ValidateOnSynthesis *bool `field:"optional" json:"validateOnSynthesis" yaml:"validateOnSynthesis"`
}

