package awsimagebuilder


// Properties for defining a `CfnImageRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageRecipeProps := &CfnImageRecipeProps{
//   	Components: []interface{}{
//   		&ComponentConfigurationProperty{
//   			ComponentArn: jsii.String("componentArn"),
//   			Parameters: []interface{}{
//   				&ComponentParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	ParentImage: jsii.String("parentImage"),
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	AdditionalInstanceConfiguration: &AdditionalInstanceConfigurationProperty{
//   		SystemsManagerAgent: &SystemsManagerAgentProperty{
//   			UninstallAfterBuild: jsii.Boolean(false),
//   		},
//   		UserDataOverride: jsii.String("userDataOverride"),
//   	},
//   	BlockDeviceMappings: []interface{}{
//   		&InstanceBlockDeviceMappingProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			Ebs: &EbsInstanceBlockDeviceSpecificationProperty{
//   				DeleteOnTermination: jsii.Boolean(false),
//   				Encrypted: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				SnapshotId: jsii.String("snapshotId"),
//   				Throughput: jsii.Number(123),
//   				VolumeSize: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			NoDevice: jsii.String("noDevice"),
//   			VirtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnImageRecipeProps struct {
	// The components of the image recipe.
	//
	// Components are orchestration documents that define a sequence of steps for downloading, installing, configuring, and testing software packages. They also define validation and security hardening steps. A component is defined using a YAML document format.
	Components interface{} `field:"required" json:"components" yaml:"components"`
	// The name of the image recipe.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parent image of the image recipe.
	//
	// The string must be either an Image ARN or an AMI ID.
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// The semantic version of the image recipe.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Before you create a new AMI, Image Builder launches temporary Amazon EC2 instances to build and test your image configuration.
	//
	// Instance configuration adds a layer of control over those instances. You can define settings and add scripts to run when an instance is launched from your AMI.
	AdditionalInstanceConfiguration interface{} `field:"optional" json:"additionalInstanceConfiguration" yaml:"additionalInstanceConfiguration"`
	// The block device mappings to apply when creating images from this recipe.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The description of the image recipe.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags of the image recipe.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

