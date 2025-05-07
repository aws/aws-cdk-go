package appstagingsynthesizeralpha


// Options for `DeploymentIdentities.defaultBootstrappedRoles`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := awscdk.NewApp(&AppProps{
//   	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
//   		AppId: jsii.String("my-app-id"),
//   		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
//
//   		// The following line is optional. By default it is assumed you have bootstrapped in the same
//   		// region(s) as the stack(s) you are deploying.
//   		DeploymentIdentities: appstagingsynthesizeralpha.DeploymentIdentities_DefaultBootstrapRoles(&DefaultBootstrapRolesOptions{
//   			BootstrapRegion: jsii.String("us-east-1"),
//   		}),
//   	}),
//   })
//
// Experimental.
type DefaultBootstrapRolesOptions struct {
	// The region where the default bootstrap roles have been created.
	//
	// By default, the region in which the stack is deployed is used.
	// Default: - the stack's current region.
	//
	// Experimental.
	BootstrapRegion *string `field:"optional" json:"bootstrapRegion" yaml:"bootstrapRegion"`
}

