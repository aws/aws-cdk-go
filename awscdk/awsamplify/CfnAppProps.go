package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppProps := &CfnAppProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AccessToken: jsii.String("accessToken"),
//   	AutoBranchCreationConfig: &AutoBranchCreationConfigProperty{
//   		AutoBranchCreationPatterns: []*string{
//   			jsii.String("autoBranchCreationPatterns"),
//   		},
//   		BasicAuthConfig: &BasicAuthConfigProperty{
//   			EnableBasicAuth: jsii.Boolean(false),
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		BuildSpec: jsii.String("buildSpec"),
//   		EnableAutoBranchCreation: jsii.Boolean(false),
//   		EnableAutoBuild: jsii.Boolean(false),
//   		EnablePerformanceMode: jsii.Boolean(false),
//   		EnablePullRequestPreview: jsii.Boolean(false),
//   		EnvironmentVariables: []interface{}{
//   			&EnvironmentVariableProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Framework: jsii.String("framework"),
//   		PullRequestEnvironmentName: jsii.String("pullRequestEnvironmentName"),
//   		Stage: jsii.String("stage"),
//   	},
//   	BasicAuthConfig: &BasicAuthConfigProperty{
//   		EnableBasicAuth: jsii.Boolean(false),
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	BuildSpec: jsii.String("buildSpec"),
//   	CacheConfig: &CacheConfigProperty{
//   		Type: jsii.String("type"),
//   	},
//   	ComputeRoleArn: jsii.String("computeRoleArn"),
//   	CustomHeaders: jsii.String("customHeaders"),
//   	CustomRules: []interface{}{
//   		&CustomRuleProperty{
//   			Source: jsii.String("source"),
//   			Target: jsii.String("target"),
//
//   			// the properties below are optional
//   			Condition: jsii.String("condition"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EnableBranchAutoDeletion: jsii.Boolean(false),
//   	EnvironmentVariables: []interface{}{
//   		&EnvironmentVariableProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	IamServiceRole: jsii.String("iamServiceRole"),
//   	JobConfig: &JobConfigProperty{
//   		BuildComputeType: jsii.String("buildComputeType"),
//   	},
//   	OauthToken: jsii.String("oauthToken"),
//   	Platform: jsii.String("platform"),
//   	Repository: jsii.String("repository"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html
//
type CfnAppProps struct {
	// The name of the Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The personal access token for a GitHub repository for an Amplify app.
	//
	// The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.
	//
	// Use `AccessToken` for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use `OauthToken` .
	//
	// You must specify either `AccessToken` or `OauthToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-accesstoken
	//
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Sets the configuration for your automatic branch creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-autobranchcreationconfig
	//
	AutoBranchCreationConfig interface{} `field:"optional" json:"autoBranchCreationConfig" yaml:"autoBranchCreationConfig"`
	// The credentials for basic authorization for an Amplify app.
	//
	// You must base64-encode the authorization credentials and provide them in the format `user:password` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-basicauthconfig
	//
	BasicAuthConfig interface{} `field:"optional" json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// The build specification (build spec) for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-buildspec
	//
	BuildSpec *string `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// The cache configuration for the Amplify app.
	//
	// If you don't specify the cache configuration `type` , Amplify uses the default `AMPLIFY_MANAGED` setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-cacheconfig
	//
	CacheConfig interface{} `field:"optional" json:"cacheConfig" yaml:"cacheConfig"`
	// The Amazon Resource Name (ARN) of the IAM role for an SSR app.
	//
	// The Compute role allows the Amplify Hosting compute service to securely access specific AWS resources based on the role's permissions. For more information about the SSR Compute role, see [Adding an SSR Compute role](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-SSR-compute-role.html) in the *Amplify User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-computerolearn
	//
	ComputeRoleArn *string `field:"optional" json:"computeRoleArn" yaml:"computeRoleArn"`
	// The custom HTTP headers for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-customheaders
	//
	CustomHeaders *string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// The custom rewrite and redirect rules for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-customrules
	//
	CustomRules interface{} `field:"optional" json:"customRules" yaml:"customRules"`
	// The description of the Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Automatically disconnect a branch in Amplify Hosting when you delete a branch from your Git repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-enablebranchautodeletion
	//
	EnableBranchAutoDeletion interface{} `field:"optional" json:"enableBranchAutoDeletion" yaml:"enableBranchAutoDeletion"`
	// The environment variables for the Amplify app.
	//
	// For a list of the environment variables that are accessible to Amplify by default, see [Amplify Environment variables](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-console-environment-variables.html) in the *Amplify Hosting User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// AWS Identity and Access Management ( IAM ) service role for the Amazon Resource Name (ARN) of the Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-iamservicerole
	//
	IamServiceRole *string `field:"optional" json:"iamServiceRole" yaml:"iamServiceRole"`
	// The configuration details that apply to the jobs for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-jobconfig
	//
	JobConfig interface{} `field:"optional" json:"jobConfig" yaml:"jobConfig"`
	// The OAuth token for a third-party source control system for an Amplify app.
	//
	// The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.
	//
	// Use `OauthToken` for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use `AccessToken` .
	//
	// You must specify either `OauthToken` or `AccessToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-oauthtoken
	//
	OauthToken *string `field:"optional" json:"oauthToken" yaml:"oauthToken"`
	// The platform for the Amplify app.
	//
	// For a static app, set the platform type to `WEB` . For a dynamic server-side rendered (SSR) app, set the platform type to `WEB_COMPUTE` . For an app requiring Amplify Hosting's original SSR support only, set the platform type to `WEB_DYNAMIC` .
	//
	// If you are deploying an SSG only app with Next.js version 14 or later, you must set the platform type to `WEB_COMPUTE` and set the artifacts `baseDirectory` to `.next` in the application's build settings. For an example of the build specification settings, see [Amplify build settings for a Next.js 14 SSG application](https://docs.aws.amazon.com/amplify/latest/userguide/deploy-nextjs-app.html#build-setting-detection-ssg-14) in the *Amplify Hosting User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-platform
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// The Git repository for the Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-repository
	//
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The tag for an Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-app.html#cfn-amplify-app-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

