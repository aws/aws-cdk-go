package appstagingsynthesizeralpha


// Options that apply to all AppStagingSynthesizer variants.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import app_staging_synthesizer_alpha "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha"
//
//   var deploymentIdentities DeploymentIdentities
//
//   appStagingSynthesizerOptions := &AppStagingSynthesizerOptions{
//   	BootstrapQualifier: jsii.String("bootstrapQualifier"),
//   	DeploymentIdentities: deploymentIdentities,
//   }
//
// Experimental.
type AppStagingSynthesizerOptions struct {
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
}

