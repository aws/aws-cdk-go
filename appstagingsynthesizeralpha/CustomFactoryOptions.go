package appstagingsynthesizeralpha


// Properties for customFactory static method.
//
// Example:
//   type customFactory struct {
//   }
//
//   func (this *customFactory) obtainStagingResources(stack Stack, context ObtainStagingResourcesContext) customStagingStack {
//   	myApp := awscdk.App_Of(*stack)
//
//   	return NewCustomStagingStack(myApp, fmt.Sprintf("CustomStagingStack-%v", *context.EnvironmentString), &customStagingStackProps{
//   	})
//   }
//
//   app := awscdk.NewApp(&AppProps{
//   	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_CustomFactory(&CustomFactoryOptions{
//   		Factory: NewCustomFactory(),
//   		OncePerEnv: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type CustomFactoryOptions struct {
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
	// The factory that will be used to return staging resources for each stack.
	// Experimental.
	Factory IStagingResourcesFactory `field:"required" json:"factory" yaml:"factory"`
	// Reuse the answer from the factory for stacks in the same environment.
	// Default: true.
	//
	// Experimental.
	OncePerEnv *bool `field:"optional" json:"oncePerEnv" yaml:"oncePerEnv"`
}

