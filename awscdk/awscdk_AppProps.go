// An experiment to bundle the entire CDK into a single module
package awscdk


// Initialization props for apps.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var context interface{}
//
//   appProps := &AppProps{
//   	AnalyticsReporting: jsii.Boolean(false),
//   	AutoSynth: jsii.Boolean(false),
//   	Context: map[string]interface{}{
//   		"contextKey": context,
//   	},
//   	Outdir: jsii.String("outdir"),
//   	RuntimeInfo: jsii.Boolean(false),
//   	StackTraces: jsii.Boolean(false),
//   	TreeMetadata: jsii.Boolean(false),
//   }
//
// Experimental.
type AppProps struct {
	// Include runtime versioning information in the Stacks of this app.
	// Experimental.
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// Automatically call `synth()` before the program exits.
	//
	// If you set this, you don't have to call `synth()` explicitly. Note that
	// this feature is only available for certain programming languages, and
	// calling `synth()` is still recommended.
	// Experimental.
	AutoSynth *bool `field:"optional" json:"autoSynth" yaml:"autoSynth"`
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdk.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// The output directory into which to emit synthesized artifacts.
	//
	// You should never need to set this value. By default, the value you pass to
	// the CLI's `--output` flag will be used, and if you change it to a different
	// directory the CLI will fail to pick up the generated Cloud Assembly.
	//
	// This property is intended for internal and testing use.
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Include runtime versioning information in the Stacks of this app.
	// Deprecated: use `versionReporting` instead.
	RuntimeInfo *bool `field:"optional" json:"runtimeInfo" yaml:"runtimeInfo"`
	// Include construct creation stack trace in the `aws:cdk:trace` metadata key of all constructs.
	// Experimental.
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
	// Include construct tree metadata as part of the Cloud Assembly.
	// Experimental.
	TreeMetadata *bool `field:"optional" json:"treeMetadata" yaml:"treeMetadata"`
}

