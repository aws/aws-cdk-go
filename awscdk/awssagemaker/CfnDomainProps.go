package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &CfnDomainProps{
//   	AuthMode: jsii.String("authMode"),
//   	DefaultUserSettings: &UserSettingsProperty{
//   		ExecutionRole: jsii.String("executionRole"),
//
//   		// the properties below are optional
//   		CodeEditorAppSettings: &CodeEditorAppSettingsProperty{
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			LifecycleConfigArns: []*string{
//   				jsii.String("lifecycleConfigArns"),
//   			},
//   		},
//   		CustomFileSystemConfigs: []interface{}{
//   			&CustomFileSystemConfigProperty{
//   				EfsFileSystemConfig: &EFSFileSystemConfigProperty{
//   					FileSystemId: jsii.String("fileSystemId"),
//
//   					// the properties below are optional
//   					FileSystemPath: jsii.String("fileSystemPath"),
//   				},
//   			},
//   		},
//   		CustomPosixUserConfig: &CustomPosixUserConfigProperty{
//   			Gid: jsii.Number(123),
//   			Uid: jsii.Number(123),
//   		},
//   		DefaultLandingUri: jsii.String("defaultLandingUri"),
//   		JupyterLabAppSettings: &JupyterLabAppSettingsProperty{
//   			CodeRepositories: []interface{}{
//   				&CodeRepositoryProperty{
//   					RepositoryUrl: jsii.String("repositoryUrl"),
//   				},
//   			},
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			LifecycleConfigArns: []*string{
//   				jsii.String("lifecycleConfigArns"),
//   			},
//   		},
//   		JupyterServerAppSettings: &JupyterServerAppSettingsProperty{
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		KernelGatewayAppSettings: &KernelGatewayAppSettingsProperty{
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		RSessionAppSettings: &RSessionAppSettingsProperty{
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		RStudioServerProAppSettings: &RStudioServerProAppSettingsProperty{
//   			AccessStatus: jsii.String("accessStatus"),
//   			UserGroup: jsii.String("userGroup"),
//   		},
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		SharingSettings: &SharingSettingsProperty{
//   			NotebookOutputOption: jsii.String("notebookOutputOption"),
//   			S3KmsKeyId: jsii.String("s3KmsKeyId"),
//   			S3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   		SpaceStorageSettings: &DefaultSpaceStorageSettingsProperty{
//   			DefaultEbsStorageSettings: &DefaultEbsStorageSettingsProperty{
//   				DefaultEbsVolumeSizeInGb: jsii.Number(123),
//   				MaximumEbsVolumeSizeInGb: jsii.Number(123),
//   			},
//   		},
//   		StudioWebPortal: jsii.String("studioWebPortal"),
//   	},
//   	DomainName: jsii.String("domainName"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	AppNetworkAccessType: jsii.String("appNetworkAccessType"),
//   	AppSecurityGroupManagement: jsii.String("appSecurityGroupManagement"),
//   	DefaultSpaceSettings: &DefaultSpaceSettingsProperty{
//   		ExecutionRole: jsii.String("executionRole"),
//
//   		// the properties below are optional
//   		JupyterServerAppSettings: &JupyterServerAppSettingsProperty{
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		KernelGatewayAppSettings: &KernelGatewayAppSettingsProperty{
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   	DomainSettings: &DomainSettingsProperty{
//   		DockerSettings: &DockerSettingsProperty{
//   			EnableDockerAccess: jsii.String("enableDockerAccess"),
//   			VpcOnlyTrustedAccounts: []*string{
//   				jsii.String("vpcOnlyTrustedAccounts"),
//   			},
//   		},
//   		RStudioServerProDomainSettings: &RStudioServerProDomainSettingsProperty{
//   			DomainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   			// the properties below are optional
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			RStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   			RStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html
//
type CfnDomainProps struct {
	// The mode of authentication that members use to access the Domain.
	//
	// *Valid Values* : `SSO | IAM`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-authmode
	//
	AuthMode *string `field:"required" json:"authMode" yaml:"authMode"`
	// The default user settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-defaultusersettings
	//
	DefaultUserSettings interface{} `field:"required" json:"defaultUserSettings" yaml:"defaultUserSettings"`
	// The domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The VPC subnets that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Array members* : Minimum number of 1 item. Maximum number of 16 items.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly` .
	//
	// - `PublicInternetOnly` - Non-EFS traffic is through a VPC managed by Amazon SageMaker , which allows direct internet access
	// - `VpcOnly` - All Studio traffic is through the specified VPC and subnets
	//
	// *Valid Values* : `PublicInternetOnly | VpcOnly`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-appnetworkaccesstype
	//
	AppNetworkAccessType *string `field:"optional" json:"appNetworkAccessType" yaml:"appNetworkAccessType"`
	// The entity that creates and manages the required security groups for inter-app communication in `VpcOnly` mode.
	//
	// Required when `CreateDomain.AppNetworkAccessType` is `VpcOnly` and `DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn` is provided. If setting up the domain for use with RStudio, this value must be set to `Service` .
	//
	// *Allowed Values* : `Service` | `Customer`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-appsecuritygroupmanagement
	//
	AppSecurityGroupManagement *string `field:"optional" json:"appSecurityGroupManagement" yaml:"appSecurityGroupManagement"`
	// A collection of settings that apply to spaces created in the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-defaultspacesettings
	//
	DefaultSpaceSettings interface{} `field:"optional" json:"defaultSpaceSettings" yaml:"defaultSpaceSettings"`
	// A collection of settings that apply to the `SageMaker Domain` .
	//
	// These settings are specified through the `CreateDomain` API call.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-domainsettings
	//
	DomainSettings interface{} `field:"optional" json:"domainSettings" yaml:"domainSettings"`
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the Domain with an AWS managed customer master key (CMK) by default.
	//
	// For more control, specify a customer managed CMK.
	//
	// *Length Constraints* : Maximum length of 2048.
	//
	// *Pattern* : `.*`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Tags to associated with the Domain.
	//
	// Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	//
	// Tags that you specify for the Domain are also added to all apps that are launched in the Domain.
	//
	// *Array members* : Minimum number of 0 items. Maximum number of 50 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

