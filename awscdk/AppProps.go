package awscdk


// Initialization props for apps.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := awscdk.NewApp(&AppProps{
//   	DefaultStackSynthesizer: awscdkappstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
//   		AppId: jsii.String("my-app-id"),
//   		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
//   		DeploymentIdentities: *awscdkappstagingsynthesizeralpha.DeploymentIdentities_CliCredentials(),
//   	}),
//   })
//
type AppProps struct {
	// Include runtime versioning information in the Stacks of this app.
	// Default: Value of 'aws:cdk:version-reporting' context key.
	//
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// Automatically call `synth()` before the program exits.
	//
	// If you set this, you don't have to call `synth()` explicitly. Note that
	// this feature is only available for certain programming languages, and
	// calling `synth()` is still recommended.
	// Default: true if running via CDK CLI (`CDK_OUTDIR` is set), `false`
	// otherwise.
	//
	AutoSynth *bool `field:"optional" json:"autoSynth" yaml:"autoSynth"`
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdk.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	// Default: - no additional context.
	//
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// The stack synthesizer to use by default for all Stacks in the App.
	//
	// The Stack Synthesizer controls aspects of synthesis and deployment,
	// like how assets are referenced and what IAM roles to use. For more
	// information, see the README of the main CDK package.
	// Default: - A `DefaultStackSynthesizer` with default settings.
	//
	DefaultStackSynthesizer IReusableStackSynthesizer `field:"optional" json:"defaultStackSynthesizer" yaml:"defaultStackSynthesizer"`
	// The output directory into which to emit synthesized artifacts.
	//
	// You should never need to set this value. By default, the value you pass to
	// the CLI's `--output` flag will be used, and if you change it to a different
	// directory the CLI will fail to pick up the generated Cloud Assembly.
	//
	// This property is intended for internal and testing use.
	// Default: - If this value is _not_ set, considers the environment variable `CDK_OUTDIR`.
	//   If `CDK_OUTDIR` is not defined, uses a temp directory.
	//
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Validation plugins to run after synthesis.
	// Default: - no validation plugins.
	//
	PolicyValidationBeta1 *[]IPolicyValidationPluginBeta1 `field:"optional" json:"policyValidationBeta1" yaml:"policyValidationBeta1"`
	// Additional context values for the application.
	//
	// Context provided here has precedence over context set by:
	//
	// - The CLI via --context
	// - The `context` key in `cdk.json`
	// - The `AppProps.context` property
	//
	// This property is recommended over the `AppProps.context` property since you
	// can make final decision over which context value to take in your app.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	//
	// Example:
	//   // context from the CLI and from `cdk.json` are stored in the
	//   // CDK_CONTEXT env variable
	//   cliContext := jSON.parse(process.env.cDK_CONTEXT)
	//
	//   // determine whether to take the context passed in the CLI or not
	//   determineValue := process.env.PROD ? cliContext.SOMEKEY : 'my-prod-value'
	//   awscdk.NewApp(&AppProps{
	//   	PostCliContext: map[string]interface{}{
	//   		"SOMEKEY": determineValue,
	//   	},
	//   })
	//
	// Default: - no additional context.
	//
	PostCliContext *map[string]interface{} `field:"optional" json:"postCliContext" yaml:"postCliContext"`
	// Include construct creation stack trace in the `aws:cdk:trace` metadata key of all constructs.
	// Default: true stack traces are included unless `aws:cdk:disable-stack-trace` is set in the context.
	//
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
	// Include construct tree metadata as part of the Cloud Assembly.
	// Default: true.
	//
	TreeMetadata *bool `field:"optional" json:"treeMetadata" yaml:"treeMetadata"`
}

