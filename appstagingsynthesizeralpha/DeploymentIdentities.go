package appstagingsynthesizeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Deployment identities are the class of roles to be assumed by the CDK when deploying the App.
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
type DeploymentIdentities interface {
	// CloudFormation Execution Role.
	// Experimental.
	CloudFormationExecutionRole() BootstrapRole
	// Deployment Action Role.
	// Experimental.
	DeploymentRole() BootstrapRole
	// Lookup Role.
	// Default: - use bootstrapped role.
	//
	// Experimental.
	LookupRole() BootstrapRole
}

// The jsii proxy struct for DeploymentIdentities
type jsiiProxy_DeploymentIdentities struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentIdentities) CloudFormationExecutionRole() BootstrapRole {
	var returns BootstrapRole
	_jsii_.Get(
		j,
		"cloudFormationExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentIdentities) DeploymentRole() BootstrapRole {
	var returns BootstrapRole
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentIdentities) LookupRole() BootstrapRole {
	var returns BootstrapRole
	_jsii_.Get(
		j,
		"lookupRole",
		&returns,
	)
	return returns
}


// Use CLI credentials for all deployment identities.
// Experimental.
func DeploymentIdentities_CliCredentials() DeploymentIdentities {
	_init_.Initialize()

	var returns DeploymentIdentities

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.DeploymentIdentities",
		"cliCredentials",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Use the Roles that have been created by the default bootstrap stack.
// Experimental.
func DeploymentIdentities_DefaultBootstrapRoles(options *DefaultBootstrapRolesOptions) DeploymentIdentities {
	_init_.Initialize()

	if err := validateDeploymentIdentities_DefaultBootstrapRolesParameters(options); err != nil {
		panic(err)
	}
	var returns DeploymentIdentities

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.DeploymentIdentities",
		"defaultBootstrapRoles",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Specify your own roles for all deployment identities.
//
// These roles
// must already exist.
// Experimental.
func DeploymentIdentities_SpecifyRoles(roles *BootstrapRoles) DeploymentIdentities {
	_init_.Initialize()

	if err := validateDeploymentIdentities_SpecifyRolesParameters(roles); err != nil {
		panic(err)
	}
	var returns DeploymentIdentities

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.DeploymentIdentities",
		"specifyRoles",
		[]interface{}{roles},
		&returns,
	)

	return returns
}

