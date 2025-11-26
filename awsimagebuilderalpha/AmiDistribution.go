package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// The regional distribution settings to use for an AMI build.
//
// Example:
//   infrastructureConfiguration := imagebuilder.NewInfrastructureConfiguration(this, jsii.String("Infrastructure"), &InfrastructureConfigurationProps{
//   	InfrastructureConfigurationName: jsii.String("production-infrastructure"),
//   	InstanceTypes: []InstanceType{
//   		ec2.InstanceType_*Of(ec2.InstanceClass_COMPUTE7_INTEL, ec2.InstanceSize_LARGE),
//   	},
//   	Vpc: vpc,
//   	SubnetSelection: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   })
//
//   distributionConfiguration := imagebuilder.NewDistributionConfiguration(this, jsii.String("Distribution"))
//   distributionConfiguration.AddAmiDistributions(&AmiDistribution{
//   	AmiName: jsii.String("production-ami-{{ imagebuilder:buildDate }}"),
//   	AmiTargetAccountIds: []*string{
//   		jsii.String("123456789012"),
//   		jsii.String("098765432109"),
//   	},
//   })
//
//   productionPipeline := imagebuilder.NewImagePipeline(this, jsii.String("ProductionPipeline"), &ImagePipelineProps{
//   	Recipe: exampleImageRecipe,
//   	InfrastructureConfiguration: infrastructureConfiguration,
//   	DistributionConfiguration: distributionConfiguration,
//   })
//
// Experimental.
type AmiDistribution struct {
	// The description of the AMI.
	// Default: None.
	//
	// Experimental.
	AmiDescription *string `field:"optional" json:"amiDescription" yaml:"amiDescription"`
	// The KMS key to encrypt the distributed AMI with.
	// Default: None.
	//
	// Experimental.
	AmiKmsKey awskms.IKey `field:"optional" json:"amiKmsKey" yaml:"amiKmsKey"`
	// The launch permissions for the AMI, defining which principals are allowed to access the AMI.
	// Default: None.
	//
	// Experimental.
	AmiLaunchPermission *AmiLaunchPermission `field:"optional" json:"amiLaunchPermission" yaml:"amiLaunchPermission"`
	// The name to use for the distributed AMIs.
	// Default: A name is generated from the image recipe name.
	//
	// Experimental.
	AmiName *string `field:"optional" json:"amiName" yaml:"amiName"`
	// The tags to apply to the distributed AMIs.
	// Default: None.
	//
	// Experimental.
	AmiTags *map[string]*string `field:"optional" json:"amiTags" yaml:"amiTags"`
	// The account IDs to copy the output AMI to.
	// Default: None.
	//
	// Experimental.
	AmiTargetAccountIds *[]*string `field:"optional" json:"amiTargetAccountIds" yaml:"amiTargetAccountIds"`
	// The fast launch configurations to use for enabling EC2 Fast Launch on the distributed Windows AMI.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnableFastLaunch.html
	//
	// Default: None.
	//
	// Experimental.
	FastLaunchConfigurations *[]*FastLaunchConfiguration `field:"optional" json:"fastLaunchConfigurations" yaml:"fastLaunchConfigurations"`
	// The launch templates to apply the distributed AMI to.
	// Default: None.
	//
	// Experimental.
	LaunchTemplates *[]*LaunchTemplateConfiguration `field:"optional" json:"launchTemplates" yaml:"launchTemplates"`
	// The License Manager license configuration ARNs to apply to the distributed AMIs.
	// Default: None.
	//
	// Experimental.
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
	// The target region to distribute AMIs to.
	// Default: The current region is used.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The SSM parameters to create or update for the distributed AMIs.
	// Default: None.
	//
	// Experimental.
	SsmParameters *[]*SSMParameterConfigurations `field:"optional" json:"ssmParameters" yaml:"ssmParameters"`
}

