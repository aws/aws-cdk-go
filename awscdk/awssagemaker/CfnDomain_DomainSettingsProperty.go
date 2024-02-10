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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html
//
type CfnDomain_DomainSettingsProperty struct {
	// A collection of settings that configure the domain's Docker interaction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-dockersettings
	//
	DockerSettings interface{} `field:"optional" json:"dockerSettings" yaml:"dockerSettings"`
	// A collection of settings that configure the `RStudioServerPro` Domain-level app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-rstudioserverprodomainsettings
	//
	RStudioServerProDomainSettings interface{} `field:"optional" json:"rStudioServerProDomainSettings" yaml:"rStudioServerProDomainSettings"`
	// The security groups for the Amazon Virtual Private Cloud that the `Domain` uses for communication between Domain-level apps and user apps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-domainsettings.html#cfn-sagemaker-domain-domainsettings-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

