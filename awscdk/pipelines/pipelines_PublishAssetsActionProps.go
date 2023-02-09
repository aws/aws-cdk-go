package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Props for a PublishAssetsAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var buildSpec buildSpec
//   var dependable iDependable
//   var role role
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   publishAssetsActionProps := &publishAssetsActionProps{
//   	actionName: jsii.String("actionName"),
//   	assetType: awscdk.Pipelines.assetType_FILE,
//   	cloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	buildSpec: buildSpec,
//   	cdkCliVersion: jsii.String("cdkCliVersion"),
//   	createBuildspecFile: jsii.Boolean(false),
//   	dependable: dependable,
//   	preInstallCommands: []*string{
//   		jsii.String("preInstallCommands"),
//   	},
//   	projectName: jsii.String("projectName"),
//   	role: role,
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   	},
//   	vpc: vpc,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type PublishAssetsActionProps struct {
	// Name of publishing action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// AssetType we're publishing.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetType AssetType `field:"required" json:"assetType" yaml:"assetType"`
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `field:"required" json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Custom BuildSpec that is merged with generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Version of CDK CLI to 'npm install'.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CdkCliVersion *string `field:"optional" json:"cdkCliVersion" yaml:"cdkCliVersion"`
	// Use a file buildspec written to the cloud assembly instead of an inline buildspec.
	//
	// This prevents size limitation errors as inline specs have a max length of 25600 characters.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CreateBuildspecFile *bool `field:"optional" json:"createBuildspecFile" yaml:"createBuildspecFile"`
	// Any Dependable construct that the CodeBuild project needs to take a dependency on.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Dependable awscdk.IDependable `field:"optional" json:"dependable" yaml:"dependable"`
	// Additional commands to run before installing cdk-assert Use this to setup proxies or npm mirrors.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PreInstallCommands *[]*string `field:"optional" json:"preInstallCommands" yaml:"preInstallCommands"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// Role to use for CodePipeline and CodeBuild to build and publish the assets.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where to execute the PublishAssetsAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

