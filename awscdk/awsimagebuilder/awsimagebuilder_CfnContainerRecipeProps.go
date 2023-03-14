package awsimagebuilder


// Properties for defining a `CfnContainerRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerRecipeProps := &cfnContainerRecipeProps{
//   	components: []interface{}{
//   		&componentConfigurationProperty{
//   			componentArn: jsii.String("componentArn"),
//   			parameters: []interface{}{
//   				&componentParameterProperty{
//   					name: jsii.String("name"),
//   					value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	containerType: jsii.String("containerType"),
//   	name: jsii.String("name"),
//   	parentImage: jsii.String("parentImage"),
//   	targetRepository: &targetContainerRepositoryProperty{
//   		repositoryName: jsii.String("repositoryName"),
//   		service: jsii.String("service"),
//   	},
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	dockerfileTemplateData: jsii.String("dockerfileTemplateData"),
//   	dockerfileTemplateUri: jsii.String("dockerfileTemplateUri"),
//   	imageOsVersionOverride: jsii.String("imageOsVersionOverride"),
//   	instanceConfiguration: &instanceConfigurationProperty{
//   		blockDeviceMappings: []interface{}{
//   			&instanceBlockDeviceMappingProperty{
//   				deviceName: jsii.String("deviceName"),
//   				ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   					deleteOnTermination: jsii.Boolean(false),
//   					encrypted: jsii.Boolean(false),
//   					iops: jsii.Number(123),
//   					kmsKeyId: jsii.String("kmsKeyId"),
//   					snapshotId: jsii.String("snapshotId"),
//   					throughput: jsii.Number(123),
//   					volumeSize: jsii.Number(123),
//   					volumeType: jsii.String("volumeType"),
//   				},
//   				noDevice: jsii.String("noDevice"),
//   				virtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		image: jsii.String("image"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	platformOverride: jsii.String("platformOverride"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnContainerRecipeProps struct {
	// Build and test components that are included in the container recipe.
	//
	// Recipes require a minimum of one build component, and can have a maximum of 20 build and test components in any combination.
	Components interface{} `field:"required" json:"components" yaml:"components"`
	// Specifies the type of container, such as Docker.
	ContainerType *string `field:"required" json:"containerType" yaml:"containerType"`
	// The name of the container recipe.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The base image for the container recipe.
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// The destination repository for the container image.
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
	Version *string `field:"required" json:"version" yaml:"version"`
	// The description of the container recipe.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside.
	//
	// The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
	DockerfileTemplateData *string `field:"optional" json:"dockerfileTemplateData" yaml:"dockerfileTemplateData"`
	// The S3 URI for the Dockerfile that will be used to build your container image.
	DockerfileTemplateUri *string `field:"optional" json:"dockerfileTemplateUri" yaml:"dockerfileTemplateUri"`
	// Specifies the operating system version for the base image.
	ImageOsVersionOverride *string `field:"optional" json:"imageOsVersionOverride" yaml:"imageOsVersionOverride"`
	// A group of options that can be used to configure an instance for building and testing container images.
	InstanceConfiguration interface{} `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Identifies which KMS key is used to encrypt the container image for distribution to the target Region.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride *string `field:"optional" json:"platformOverride" yaml:"platformOverride"`
	// Tags that are attached to the container recipe.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The working directory for use during build and test workflows.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

