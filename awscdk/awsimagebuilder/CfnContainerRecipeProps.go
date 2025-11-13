package awsimagebuilder


// Properties for defining a `CfnContainerRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerRecipeProps := &CfnContainerRecipeProps{
//   	ContainerType: jsii.String("containerType"),
//   	Name: jsii.String("name"),
//   	ParentImage: jsii.String("parentImage"),
//   	TargetRepository: &TargetContainerRepositoryProperty{
//   		RepositoryName: jsii.String("repositoryName"),
//   		Service: jsii.String("service"),
//   	},
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
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
//   	Description: jsii.String("description"),
//   	DockerfileTemplateData: jsii.String("dockerfileTemplateData"),
//   	DockerfileTemplateUri: jsii.String("dockerfileTemplateUri"),
//   	ImageOsVersionOverride: jsii.String("imageOsVersionOverride"),
//   	InstanceConfiguration: &InstanceConfigurationProperty{
//   		BlockDeviceMappings: []interface{}{
//   			&InstanceBlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EbsInstanceBlockDeviceSpecificationProperty{
//   					DeleteOnTermination: jsii.Boolean(false),
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					SnapshotId: jsii.String("snapshotId"),
//   					Throughput: jsii.Number(123),
//   					VolumeSize: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				NoDevice: jsii.String("noDevice"),
//   				VirtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		Image: jsii.String("image"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PlatformOverride: jsii.String("platformOverride"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html
//
type CfnContainerRecipeProps struct {
	// Specifies the type of container, such as Docker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-containertype
	//
	ContainerType *string `field:"required" json:"containerType" yaml:"containerType"`
	// The name of the container recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The base image for customizations specified in the container recipe.
	//
	// This can contain an Image Builder image resource ARN or a container image URI, for example `amazonlinux:latest` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-parentimage
	//
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// The destination repository for the container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-targetrepository
	//
	TargetRepository interface{} `field:"required" json:"targetRepository" yaml:"targetRepository"`
	// The semantic version of the container recipe.
	//
	// > The semantic version has four nodes: <major>.<minor>.<patch>/<build>. You can assign values for the first three, and can filter on all of them.
	// >
	// > *Assignment:* For the first three nodes you can assign any positive integer value, including zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder automatically assigns the build number to the fourth node.
	// >
	// > *Patterns:* You can use any numeric pattern that adheres to the assignment requirements for the nodes that you can assign. For example, you might choose a software version pattern, such as 1.0.0, or a date, such as 2021.01.01.
	// >
	// > *Filtering:* With semantic versioning, you have the flexibility to use wildcards (x) to specify the most recent versions or nodes when selecting the base image or components for your recipe. When you use a wildcard in any node, all nodes to the right of the first wildcard must also be wildcards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
	// Build and test components that are included in the container recipe.
	//
	// Recipes require a minimum of one build component, and can have a maximum of 20 build and test components in any combination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-components
	//
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// The description of the container recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside.
	//
	// The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-dockerfiletemplatedata
	//
	DockerfileTemplateData *string `field:"optional" json:"dockerfileTemplateData" yaml:"dockerfileTemplateData"`
	// The S3 URI for the Dockerfile that will be used to build your container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-dockerfiletemplateuri
	//
	DockerfileTemplateUri *string `field:"optional" json:"dockerfileTemplateUri" yaml:"dockerfileTemplateUri"`
	// Specifies the operating system version for the base image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-imageosversionoverride
	//
	ImageOsVersionOverride *string `field:"optional" json:"imageOsVersionOverride" yaml:"imageOsVersionOverride"`
	// A group of options that can be used to configure an instance for building and testing container images.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-instanceconfiguration
	//
	InstanceConfiguration interface{} `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// The Amazon Resource Name (ARN) that uniquely identifies which KMS key is used to encrypt the container image for distribution to the target Region.
	//
	// This can be either the Key ARN or the Alias ARN. For more information, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the operating system platform when you use a custom base image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-platformoverride
	//
	PlatformOverride *string `field:"optional" json:"platformOverride" yaml:"platformOverride"`
	// Tags that are attached to the container recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The working directory for use during build and test workflows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html#cfn-imagebuilder-containerrecipe-workingdirectory
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

