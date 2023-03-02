package awsimagebuilder


// Define and configure the output AMIs of the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amiDistributionConfigurationProperty := &amiDistributionConfigurationProperty{
//   	amiTags: map[string]*string{
//   		"amiTagsKey": jsii.String("amiTags"),
//   	},
//   	description: jsii.String("description"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	launchPermissionConfiguration: &launchPermissionConfigurationProperty{
//   		organizationalUnitArns: []*string{
//   			jsii.String("organizationalUnitArns"),
//   		},
//   		organizationArns: []*string{
//   			jsii.String("organizationArns"),
//   		},
//   		userGroups: []*string{
//   			jsii.String("userGroups"),
//   		},
//   		userIds: []*string{
//   			jsii.String("userIds"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	targetAccountIds: []*string{
//   		jsii.String("targetAccountIds"),
//   	},
//   }
//
type CfnDistributionConfiguration_AmiDistributionConfigurationProperty struct {
	// The tags to apply to AMIs distributed to this Region.
	AmiTags interface{} `field:"optional" json:"amiTags" yaml:"amiTags"`
	// The description of the AMI distribution configuration.
	//
	// Minimum and maximum length are in characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key identifier used to encrypt the distributed image.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Launch permissions can be used to configure which AWS account s can use the AMI to launch instances.
	LaunchPermissionConfiguration interface{} `field:"optional" json:"launchPermissionConfiguration" yaml:"launchPermissionConfiguration"`
	// The name of the output AMI.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of an account to which you want to distribute an image.
	TargetAccountIds *[]*string `field:"optional" json:"targetAccountIds" yaml:"targetAccountIds"`
}

