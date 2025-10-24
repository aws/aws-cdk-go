package awssagemaker


// A collection of settings that apply to the `SageMaker Domain` .
//
// These settings are specified through the `CreateDomain` API call.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainSettingsProperty := &DomainSettingsProperty{
//   	DockerSettings: &DockerSettingsProperty{
//   		EnableDockerAccess: jsii.String("enableDockerAccess"),
//   		VpcOnlyTrustedAccounts: []*string{
//   			jsii.String("vpcOnlyTrustedAccounts"),
//   		},
//   	},
//   	ExecutionRoleIdentityConfig: jsii.String("executionRoleIdentityConfig"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	RStudioServerProDomainSettings: &RStudioServerProDomainSettingsProperty{
//   		DomainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   		// the properties below are optional
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   		RStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   		RStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	UnifiedStudioSettings: &UnifiedStudioSettingsProperty{
//   		DomainAccountId: jsii.String("domainAccountId"),
//   		DomainId: jsii.String("domainId"),
//   		DomainRegion: jsii.String("domainRegion"),
//   		EnvironmentId: jsii.String("environmentId"),
//   		ProjectId: jsii.String("projectId"),
//   		ProjectS3Path: jsii.String("projectS3Path"),
//   		StudioWebPortalAccess: jsii.String("studioWebPortalAccess"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html
//
type CfnDomain_DomainSettingsProperty struct {
	// A collection of settings that configure the domain's Docker interaction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-dockersettings
	//
	DockerSettings interface{} `field:"optional" json:"dockerSettings" yaml:"dockerSettings"`
	// The configuration for attaching a SageMaker AI user profile name to the execution role as a [sts:SourceIdentity key](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_monitor.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-executionroleidentityconfig
	//
	ExecutionRoleIdentityConfig *string `field:"optional" json:"executionRoleIdentityConfig" yaml:"executionRoleIdentityConfig"`
	// The IP address type for the domain.
	//
	// Specify `ipv4` for IPv4-only connectivity or `dualstack` for both IPv4 and IPv6 connectivity. When you specify `dualstack` , the subnet must support IPv6 CIDR blocks. If not specified, defaults to `ipv4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A collection of settings that configure the `RStudioServerPro` Domain-level app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-rstudioserverprodomainsettings
	//
	RStudioServerProDomainSettings interface{} `field:"optional" json:"rStudioServerProDomainSettings" yaml:"rStudioServerProDomainSettings"`
	// The security groups for the Amazon Virtual Private Cloud that the `Domain` uses for communication between Domain-level apps and user apps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The settings that apply to an SageMaker AI domain when you use it in Amazon SageMaker Unified Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-unifiedstudiosettings
	//
	UnifiedStudioSettings interface{} `field:"optional" json:"unifiedStudioSettings" yaml:"unifiedStudioSettings"`
}

