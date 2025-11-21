package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// The regional distribution settings to use for an AMI build.
//
// Example:
//   distributionConfiguration := imagebuilder.NewDistributionConfiguration(this, jsii.String("DistributionConfiguration"), &DistributionConfigurationProps{
//   	DistributionConfigurationName: jsii.String("test-distribution-configuration"),
//   	Description: jsii.String("A Distribution Configuration"),
//   	AmiDistributions: []AmiDistribution{
//   		&AmiDistribution{
//   			// Distribute AMI to us-east-2 and publish the AMI ID to an SSM parameter
//   			Region: jsii.String("us-east-2"),
//   			SsmParameters: []SSMParameterConfigurations{
//   				&SSMParameterConfigurations{
//   					Parameter: ssm.StringParameter_FromStringParameterAttributes(this, jsii.String("CrossRegionParameter"), &StringParameterAttributes{
//   						ParameterName: jsii.String("/imagebuilder/ami"),
//   						ForceDynamicReference: jsii.Boolean(true),
//   					}),
//   				},
//   			},
//   		},
//   	},
//   })
//
//   // For AMI-based image builds - add an AMI distribution in the current region
//   distributionConfiguration.AddAmiDistributions(&AmiDistribution{
//   	AmiName: jsii.String("imagebuilder-{{ imagebuilder:buildDate }}"),
//   	AmiDescription: jsii.String("Build AMI"),
//   	AmiKmsKey: kms.Key_FromLookup(this, jsii.String("ComponentKey"), &KeyLookupOptions{
//   		AliasName: jsii.String("alias/distribution-encryption-key"),
//   	}),
//   	// Copy the AMI to different accounts
//   	AmiTargetAccountIds: []*string{
//   		jsii.String("123456789012"),
//   		jsii.String("098765432109"),
//   	},
//   	// Add launch permissions on the AMI
//   	AmiLaunchPermission: &AmiLaunchPermission{
//   		OrganizationArns: []*string{
//   			this.FormatArn(&ArnComponents{
//   				Region: jsii.String(""),
//   				Service: jsii.String("organizations"),
//   				Resource: jsii.String("organization"),
//   				ResourceName: jsii.String("o-1234567abc"),
//   			}),
//   		},
//   		OrganizationalUnitArns: []*string{
//   			this.*FormatArn(&ArnComponents{
//   				Region: jsii.String(""),
//   				Service: jsii.String("organizations"),
//   				Resource: jsii.String("ou"),
//   				ResourceName: jsii.String("o-1234567abc/ou-a123-b4567890"),
//   			}),
//   		},
//   		IsPublicUserGroup: jsii.Boolean(true),
//   		AccountIds: []*string{
//   			jsii.String("234567890123"),
//   		},
//   	},
//   	// Attach tags to the AMI
//   	AmiTags: map[string]*string{
//   		"Environment": jsii.String("production"),
//   		"Version": jsii.String("{{ imagebuilder:buildVersion }}"),
//   	},
//   	// Optional - publish the distributed AMI ID to an SSM parameter
//   	SsmParameters: []SSMParameterConfigurations{
//   		&SSMParameterConfigurations{
//   			Parameter: ssm.StringParameter_*FromStringParameterAttributes(this, jsii.String("Parameter"), &StringParameterAttributes{
//   				ParameterName: jsii.String("/imagebuilder/ami"),
//   				ForceDynamicReference: jsii.Boolean(true),
//   			}),
//   		},
//   		&SSMParameterConfigurations{
//   			AmiAccount: jsii.String("098765432109"),
//   			DataType: ssm.ParameterDataType_TEXT,
//   			Parameter: ssm.StringParameter_*FromStringParameterAttributes(this, jsii.String("CrossAccountParameter"), &StringParameterAttributes{
//   				ParameterName: jsii.String("imagebuilder-prod-ami"),
//   				ForceDynamicReference: jsii.Boolean(true),
//   			}),
//   		},
//   	},
//   	// Optional - create a new launch template version with the distributed AMI ID
//   	LaunchTemplates: []LaunchTemplateConfiguration{
//   		&LaunchTemplateConfiguration{
//   			LaunchTemplate: ec2.LaunchTemplate_FromLaunchTemplateAttributes(this, jsii.String("LaunchTemplate"), &LaunchTemplateAttributes{
//   				LaunchTemplateId: jsii.String("lt-1234"),
//   			}),
//   			SetDefaultVersion: jsii.Boolean(true),
//   		},
//   		&LaunchTemplateConfiguration{
//   			AccountId: jsii.String("123456789012"),
//   			LaunchTemplate: ec2.LaunchTemplate_*FromLaunchTemplateAttributes(this, jsii.String("CrossAccountLaunchTemplate"), &LaunchTemplateAttributes{
//   				LaunchTemplateId: jsii.String("lt-5678"),
//   			}),
//   			SetDefaultVersion: jsii.Boolean(true),
//   		},
//   	},
//   	// Optional - enable Fast Launch on an imported launch template
//   	FastLaunchConfigurations: []FastLaunchConfiguration{
//   		&FastLaunchConfiguration{
//   			Enabled: jsii.Boolean(true),
//   			LaunchTemplate: ec2.LaunchTemplate_*FromLaunchTemplateAttributes(this, jsii.String("FastLaunchLT"), &LaunchTemplateAttributes{
//   				LaunchTemplateName: jsii.String("fast-launch-lt"),
//   			}),
//   			MaxParallelLaunches: jsii.Number(10),
//   			TargetSnapshotCount: jsii.Number(2),
//   		},
//   	},
//   	// Optional - license configurations to apply to the AMI
//   	LicenseConfigurationArns: []*string{
//   		jsii.String("arn:aws:license-manager:us-west-2:123456789012:license-configuration:lic-abcdefghijklmnopqrstuvwxyz"),
//   	},
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

