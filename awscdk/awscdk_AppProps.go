// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Initialization props for apps.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var context interface{}
//
//   appProps := &appProps{
//   	analyticsReporting: jsii.Boolean(false),
//   	autoSynth: jsii.Boolean(false),
//   	context: map[string]interface{}{
//   		"contextKey": context,
//   	},
//   	outdir: jsii.String("outdir"),
//   	stackTraces: jsii.Boolean(false),
//   	treeMetadata: jsii.Boolean(false),
//   }
//
type AppProps struct {
	// Include runtime versioning information in the Stacks of this app.
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// Automatically call `synth()` before the program exits.
	//
	// If you set this, you don't have to call `synth()` explicitly. Note that
	// this feature is only available for certain programming languages, and
	// calling `synth()` is still recommended.
	AutoSynth *bool `field:"optional" json:"autoSynth" yaml:"autoSynth"`
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdk.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// The output directory into which to emit synthesized artifacts.
	//
	// You should never need to set this value. By default, the value you pass to
	// the CLI's `--output` flag will be used, and if you change it to a different
	// directory the CLI will fail to pick up the generated Cloud Assembly.
	//
	// This property is intended for internal and testing use.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Include construct creation stack trace in the `aws:cdk:trace` metadata key of all constructs.
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
	// Include construct tree metadata as part of the Cloud Assembly.
	TreeMetadata *bool `field:"optional" json:"treeMetadata" yaml:"treeMetadata"`
}

