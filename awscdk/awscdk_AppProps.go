// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Initialization props for apps.
//
// Example:
//   awscdk.NewApp(&appProps{
//   	context: map[string]interface{}{
//   		awscdk.PERMISSIONS_BOUNDARY_CONTEXT_KEY: map[string]*string{
//   			"name": jsii.String("cdk-${Qualifier}-PermissionsBoundary"),
//   		},
//   	},
//   })
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
	// Additional context values for the application.
	//
	// Context provided here has precedence over context set by:
	//
	// - The CLI via --context
	// - The `context` key in `cdk.json`
	// - The {@link AppProps.context} property
	//
	// This property is recommended over the {@link AppProps.context} property since you
	// can make final decision over which context value to take in your app.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   // context from the CLI and from `cdk.json` are stored in the
	//   // CDK_CONTEXT env variable
	//   cliContext := jSON.parse(process.env.cDK_CONTEXT)
	//
	//   // determine whether to take the context passed in the CLI or not
	//   determineValue := process.env.PROD ? cliContext.SOMEKEY : 'my-prod-value'
	//   awscdk.NewApp(&appProps{
	//   	postCliContext: map[string]interface{}{
	//   		"SOMEKEY": determineValue,
	//   	},
	//   })
	//
	PostCliContext *map[string]interface{} `field:"optional" json:"postCliContext" yaml:"postCliContext"`
	// Include construct creation stack trace in the `aws:cdk:trace` metadata key of all constructs.
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
	// Include construct tree metadata as part of the Cloud Assembly.
	TreeMetadata *bool `field:"optional" json:"treeMetadata" yaml:"treeMetadata"`
}

