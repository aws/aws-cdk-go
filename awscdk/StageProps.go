package awscdk


// Initialization props for a stage.
//
// Example:
//   var app app
//
//
//   awscdk.NewStage(app, jsii.String("DevStage"))
//
//   awscdk.NewStage(app, jsii.String("BetaStage"), &StageProps{
//   	PermissionsBoundary: awscdk.PermissionsBoundary_FromName(jsii.String("beta-permissions-boundary")),
//   })
//
//   awscdk.NewStage(app, jsii.String("GammaStage"), &StageProps{
//   	PermissionsBoundary: awscdk.PermissionsBoundary_*FromName(jsii.String("prod-permissions-boundary")),
//   })
//
//   awscdk.NewStage(app, jsii.String("ProdStage"), &StageProps{
//   	PermissionsBoundary: awscdk.PermissionsBoundary_*FromName(jsii.String("prod-permissions-boundary")),
//   })
//
type StageProps struct {
	// Default AWS environment (account/region) for `Stack`s in this `Stage`.
	//
	// Stacks defined inside this `Stage` with either `region` or `account` missing
	// from its env will use the corresponding field given here.
	//
	// If either `region` or `account`is is not configured for `Stack` (either on
	// the `Stack` itself or on the containing `Stage`), the Stack will be
	// *environment-agnostic*.
	//
	// Environment-agnostic stacks can be deployed to any environment, may not be
	// able to take advantage of all features of the CDK. For example, they will
	// not be able to use environmental context lookups, will not automatically
	// translate Service Principals to the right format based on the environment's
	// AWS partition, and other such enhancements.
	//
	// Example:
	//   // Use a concrete account and region to deploy this Stage to
	//   // Use a concrete account and region to deploy this Stage to
	//   awscdk.NewStage(app, jsii.String("Stage1"), &StageProps{
	//   	Env: &Environment{
	//   		Account: jsii.String("123456789012"),
	//   		Region: jsii.String("us-east-1"),
	//   	},
	//   })
	//
	//   // Use the CLI's current credentials to determine the target environment
	//   // Use the CLI's current credentials to determine the target environment
	//   awscdk.NewStage(app, jsii.String("Stage2"), &StageProps{
	//   	Env: &Environment{
	//   		Account: process.env.cDK_DEFAULT_ACCOUNT,
	//   		Region: process.env.cDK_DEFAULT_REGION,
	//   	},
	//   })
	//
	// Default: - The environments should be configured on the `Stack`s.
	//
	Env *Environment `field:"optional" json:"env" yaml:"env"`
	// The output directory into which to emit synthesized artifacts.
	//
	// Can only be specified if this stage is the root stage (the app). If this is
	// specified and this stage is nested within another stage, an error will be
	// thrown.
	// Default: - for nested stages, outdir will be determined as a relative
	// directory to the outdir of the app. For apps, if outdir is not specified, a
	// temporary directory will be created.
	//
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Options for applying a permissions boundary to all IAM Roles and Users created within this Stage.
	// Default: - no permissions boundary is applied.
	//
	PermissionsBoundary PermissionsBoundary `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Validation plugins to run during synthesis.
	//
	// If any plugin reports any violation,
	// synthesis will be interrupted and the report displayed to the user.
	// Default: - no validation plugins are used.
	//
	PolicyValidationBeta1 *[]IPolicyValidationPluginBeta1 `field:"optional" json:"policyValidationBeta1" yaml:"policyValidationBeta1"`
	// Name of this stage.
	// Default: - Derived from the id.
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

