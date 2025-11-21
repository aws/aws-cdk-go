package awsimagebuilderalpha


// The regional distribution settings to use for a container build.
//
// Example:
//   ecrRepository := ecr.Repository_FromRepositoryName(this, jsii.String("ECRRepository"), jsii.String("my-repo"))
//   containerRepository := imagebuilder.Repository_FromEcr(ecrRepository)
//   containerDistributionConfiguration := imagebuilder.NewDistributionConfiguration(this, jsii.String("ContainerDistributionConfiguration"))
//
//   containerDistributionConfiguration.AddContainerDistributions(&ContainerDistribution{
//   	ContainerRepository: ContainerRepository,
//   	ContainerDescription: jsii.String("Test container image"),
//   	ContainerTags: []*string{
//   		jsii.String("latest"),
//   		jsii.String("latest-1.0"),
//   	},
//   })
//
// Experimental.
type ContainerDistribution struct {
	// The destination repository to distribute the output container to.
	// Default: The target repository in the container recipe is used.
	//
	// Experimental.
	ContainerRepository Repository `field:"required" json:"containerRepository" yaml:"containerRepository"`
	// The description of the container image.
	// Default: None.
	//
	// Experimental.
	ContainerDescription *string `field:"optional" json:"containerDescription" yaml:"containerDescription"`
	// The additional tags to apply to the distributed container images.
	// Default: None.
	//
	// Experimental.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// The target region to distribute containers to.
	// Default: The current region is used.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

