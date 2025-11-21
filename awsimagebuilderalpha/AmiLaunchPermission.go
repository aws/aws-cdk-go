package awsimagebuilderalpha


// The launch permissions for the AMI, defining which principals are allowed to access the AMI.
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
type AmiLaunchPermission struct {
	// The AWS account IDs to share the AMI with.
	// Default: None.
	//
	// Experimental.
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// Whether to make the AMI public. Block public access for AMIs must be disabled to make the AMI public.
	//
	// WARNING: Making an AMI public exposes it to any AWS account globally.
	//  Ensure the AMI does not contain:
	//  - Sensitive data or credentials
	//  - Proprietary software or configurations
	//  - Internal network information or security settings
	//
	// For more information on blocking public access for AMIs, see: [Understand block public access for AMIs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-public-access-to-amis.html)
	// Default: false.
	//
	// Experimental.
	IsPublicUserGroup *bool `field:"optional" json:"isPublicUserGroup" yaml:"isPublicUserGroup"`
	// The ARNs for the AWS Organizations organizational units to share the AMI with.
	// Default: None.
	//
	// Experimental.
	OrganizationalUnitArns *[]*string `field:"optional" json:"organizationalUnitArns" yaml:"organizationalUnitArns"`
	// The ARNs for the AWS Organization that you want to share the AMI with.
	// Default: None.
	//
	// Experimental.
	OrganizationArns *[]*string `field:"optional" json:"organizationArns" yaml:"organizationArns"`
}

