package awsimagebuilderalpha


// Properties for creating a Distribution Configuration resource.
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
type DistributionConfigurationProps struct {
	// The list of target regions and associated AMI distribution settings where the built AMI will be distributed.
	//
	// AMI
	// distributions may also be added with the `addAmiDistributions` method.
	// Default: None if container distributions are provided. Otherwise, at least one AMI or container distribution must
	// be provided.
	//
	// Experimental.
	AmiDistributions *[]*AmiDistribution `field:"optional" json:"amiDistributions" yaml:"amiDistributions"`
	// The list of target regions and associated container distribution settings where the built container will be distributed.
	//
	// Container distributions may also be added with the `addContainerDistributions` method.
	// Default: None if AMI distributions are provided. Otherwise, at least one AMI or container distribution must be
	// provided.
	//
	// Experimental.
	ContainerDistributions *[]*ContainerDistribution `field:"optional" json:"containerDistributions" yaml:"containerDistributions"`
	// The description of the distribution configuration.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the distribution configuration.
	// Default: A name is generated.
	//
	// Experimental.
	DistributionConfigurationName *string `field:"optional" json:"distributionConfigurationName" yaml:"distributionConfigurationName"`
	// The tags to apply to the distribution configuration.
	// Default: None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

