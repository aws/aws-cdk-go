package awsimagebuilder


// Define and configure the output AMIs of the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amiDistributionConfigurationProperty := &AmiDistributionConfigurationProperty{
//   	AmiTags: map[string]*string{
//   		"amiTagsKey": jsii.String("amiTags"),
//   	},
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LaunchPermissionConfiguration: &LaunchPermissionConfigurationProperty{
//   		OrganizationalUnitArns: []*string{
//   			jsii.String("organizationalUnitArns"),
//   		},
//   		OrganizationArns: []*string{
//   			jsii.String("organizationArns"),
//   		},
//   		UserGroups: []*string{
//   			jsii.String("userGroups"),
//   		},
//   		UserIds: []*string{
//   			jsii.String("userIds"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	TargetAccountIds: []*string{
//   		jsii.String("targetAccountIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html
//
type CfnDistributionConfiguration_AmiDistributionConfigurationProperty struct {
	// The tags to apply to AMIs distributed to this Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-amitags
	//
	AmiTags interface{} `field:"optional" json:"amiTags" yaml:"amiTags"`
	// The description of the AMI distribution configuration.
	//
	// Minimum and maximum length are in characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key identifier used to encrypt the distributed image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Launch permissions can be used to configure which AWS account s can use the AMI to launch instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-launchpermissionconfiguration
	//
	LaunchPermissionConfiguration interface{} `field:"optional" json:"launchPermissionConfiguration" yaml:"launchPermissionConfiguration"`
	// The name of the output AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of an account to which you want to distribute an image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-targetaccountids
	//
	TargetAccountIds *[]*string `field:"optional" json:"targetAccountIds" yaml:"targetAccountIds"`
}

