package appstagingsynthesizeralpha


// Roles that are bootstrapped to your account.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := awscdk.NewApp(&AppProps{
//   	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
//   		AppId: jsii.String("my-app-id"),
//   		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
//   		DeploymentIdentities: appstagingsynthesizeralpha.DeploymentIdentities_SpecifyRoles(&BootstrapRoles{
//   			CloudFormationExecutionRole: appstagingsynthesizeralpha.BootstrapRole_FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Execute")),
//   			DeploymentRole: appstagingsynthesizeralpha.BootstrapRole_*FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Deploy")),
//   			LookupRole: appstagingsynthesizeralpha.BootstrapRole_*FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Lookup")),
//   		}),
//   	}),
//   })
//
// Experimental.
type BootstrapRoles struct {
	// CloudFormation Execution Role.
	// Default: - use bootstrapped role.
	//
	// Experimental.
	CloudFormationExecutionRole BootstrapRole `field:"optional" json:"cloudFormationExecutionRole" yaml:"cloudFormationExecutionRole"`
	// Deployment Action Role.
	// Default: - use boostrapped role.
	//
	// Experimental.
	DeploymentRole BootstrapRole `field:"optional" json:"deploymentRole" yaml:"deploymentRole"`
	// Lookup Role.
	// Default: - use bootstrapped role.
	//
	// Experimental.
	LookupRole BootstrapRole `field:"optional" json:"lookupRole" yaml:"lookupRole"`
}

