// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Initialization props for a stage.
//
// Example:
//   var app app
//
//
//   awscdk.NewStage(app, jsii.String("DevStage"))
//
//   awscdk.NewStage(app, jsii.String("BetaStage"), &stageProps{
//   	permissionsBoundary: awscdk.PermissionsBoundary.fromName(jsii.String("beta-permissions-boundary")),
//   })
//
//   awscdk.NewStage(app, jsii.String("GammaStage"), &stageProps{
//   	permissionsBoundary: awscdk.PermissionsBoundary.fromName(jsii.String("prod-permissions-boundary")),
//   })
//
//   awscdk.NewStage(app, jsii.String("ProdStage"), &stageProps{
//   	permissionsBoundary: awscdk.PermissionsBoundary.fromName(jsii.String("prod-permissions-boundary")),
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
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   // Use a concrete account and region to deploy this Stage to
	//   // Use a concrete account and region to deploy this Stage to
	//   awscdk.NewStage(app, jsii.String("Stage1"), &stageProps{
	//   	env: &environment{
	//   		account: jsii.String("123456789012"),
	//   		region: jsii.String("us-east-1"),
	//   	},
	//   })
	//
	//   // Use the CLI's current credentials to determine the target environment
	//   // Use the CLI's current credentials to determine the target environment
	//   awscdk.NewStage(app, jsii.String("Stage2"), &stageProps{
	//   	env: &environment{
	//   		account: process.env.cDK_DEFAULT_ACCOUNT,
	//   		region: process.env.cDK_DEFAULT_REGION,
	//   	},
	//   })
	//
	Env *Environment `field:"optional" json:"env" yaml:"env"`
	// The output directory into which to emit synthesized artifacts.
	//
	// Can only be specified if this stage is the root stage (the app). If this is
	// specified and this stage is nested within another stage, an error will be
	// thrown.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Options for applying a permissions boundary to all IAM Roles and Users created within this Stage.
	PermissionsBoundary PermissionsBoundary `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Name of this stage.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

