package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// Properties for a Branch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import amplify_alpha "github.com/aws/aws-cdk-go/awscdkamplifyalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var app App
//   var asset Asset
//   var basicAuth BasicAuth
//   var buildSpec BuildSpec
//   var role Role
//
//   branchProps := &BranchProps{
//   	App: app,
//
//   	// the properties below are optional
//   	Asset: asset,
//   	AutoBuild: jsii.Boolean(false),
//   	BasicAuth: basicAuth,
//   	BranchName: jsii.String("branchName"),
//   	BuildSpec: buildSpec,
//   	ComputeRole: role,
//   	Description: jsii.String("description"),
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	PerformanceMode: jsii.Boolean(false),
//   	PullRequestEnvironmentName: jsii.String("pullRequestEnvironmentName"),
//   	PullRequestPreview: jsii.Boolean(false),
//   	SkewProtection: jsii.Boolean(false),
//   	Stage: jsii.String("stage"),
//   }
//
// Experimental.
type BranchProps struct {
	// Asset for deployment.
	//
	// The Amplify app must not have a sourceCodeProvider configured as this resource uses Amplify's
	// startDeployment API to initiate and deploy a S3 asset onto the App.
	// Default: - no asset.
	//
	// Experimental.
	Asset awss3assets.Asset `field:"optional" json:"asset" yaml:"asset"`
	// Whether to enable auto building for the branch.
	// Default: true.
	//
	// Experimental.
	AutoBuild *bool `field:"optional" json:"autoBuild" yaml:"autoBuild"`
	// The Basic Auth configuration.
	//
	// Use this to set password protection for
	// the branch.
	// Default: - no password protection.
	//
	// Experimental.
	BasicAuth BasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// The name of the branch.
	// Default: - the construct's id.
	//
	// Experimental.
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// BuildSpec for the branch.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html
	//
	// Default: - no build spec.
	//
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// The IAM role to assign to a branch of an SSR app.
	//
	// The SSR Compute role allows the Amplify Hosting compute service to securely access specific AWS resources based on the role's permissions.
	//
	// This role overrides the app-level compute role.
	// Default: undefined - No specific role for the branch. If the app has a compute role, it will be inherited.
	//
	// Experimental.
	ComputeRole awsiam.IRole `field:"optional" json:"computeRole" yaml:"computeRole"`
	// A description for the branch.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Environment variables for the branch.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Default: - application environment variables.
	//
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Enables performance mode for the branch.
	//
	// Performance mode optimizes for faster hosting performance by keeping content cached at the edge
	// for a longer interval. When performance mode is enabled, hosting configuration or code changes
	// can take up to 10 minutes to roll out.
	// Default: false.
	//
	// Experimental.
	PerformanceMode *bool `field:"optional" json:"performanceMode" yaml:"performanceMode"`
	// The dedicated backend environment for the pull request previews.
	// Default: - automatically provision a temporary backend.
	//
	// Experimental.
	PullRequestEnvironmentName *string `field:"optional" json:"pullRequestEnvironmentName" yaml:"pullRequestEnvironmentName"`
	// Whether to enable pull request preview for the branch.
	// Default: true.
	//
	// Experimental.
	PullRequestPreview *bool `field:"optional" json:"pullRequestPreview" yaml:"pullRequestPreview"`
	// Specifies whether the skew protection feature is enabled for the branch.
	//
	// Deployment skew protection is available to Amplify applications to eliminate version skew issues
	// between client and servers in web applications.
	// When you apply skew protection to a branch, you can ensure that your clients always interact
	// with the correct version of server-side assets, regardless of when a deployment occurs.
	// Default: None - Default setting is no skew protection.
	//
	// Experimental.
	SkewProtection *bool `field:"optional" json:"skewProtection" yaml:"skewProtection"`
	// Stage for the branch.
	// Default: - no stage.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// The application within which the branch must be created.
	// Experimental.
	App IApp `field:"required" json:"app" yaml:"app"`
}

