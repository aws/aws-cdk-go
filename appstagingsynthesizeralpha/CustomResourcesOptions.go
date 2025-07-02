package appstagingsynthesizeralpha


// Properties for customResources static method.
//
// Example:
//   resourceApp := awscdk.NewApp()
//   resources := NewCustomStagingStack(resourceApp, jsii.String("CustomStagingStack"), &customStagingStackProps{
//   })
//
//   app := awscdk.NewApp(&AppProps{
//   	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_CustomResources(&CustomResourcesOptions{
//   		Resources: *Resources,
//   	}),
//   })
//
// Experimental.
type CustomResourcesOptions struct {
	// Qualifier to disambiguate multiple bootstrapped environments in the same account.
	//
	// This qualifier is only used to reference bootstrapped resources. It will not
	// be used in the creation of app-specific staging resources: `appId` is used for that
	// instead.
	// Default: - Value of context key '@aws-cdk/core:bootstrapQualifier' if set, otherwise `DEFAULT_QUALIFIER`.
	//
	// Experimental.
	BootstrapQualifier *string `field:"optional" json:"bootstrapQualifier" yaml:"bootstrapQualifier"`
	// What roles to use to deploy applications.
	//
	// These are the roles that have permissions to interact with CloudFormation
	// on your behalf. By default these are the standard bootstrapped CDK roles,
	// but you can customize them or turn them off and use the CLI credentials
	// to deploy.
	// Default: - The standard bootstrapped CDK roles.
	//
	// Experimental.
	DeploymentIdentities DeploymentIdentities `field:"optional" json:"deploymentIdentities" yaml:"deploymentIdentities"`
	// Use these exact staging resources for every stack that this synthesizer is used for.
	// Experimental.
	Resources IStagingResources `field:"required" json:"resources" yaml:"resources"`
}

