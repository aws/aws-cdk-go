package awsimagebuilderalpha


// Options for selecting a predefined Amazon-managed image.
//
// Example:
//   // Amazon Linux 2023 AMI for x86_64
//   amazonLinux2023Ami := imagebuilder.AmazonManagedImage_AmazonLinux2023(this, jsii.String("AmazonLinux2023"), &AmazonManagedImageOptions{
//   	ImageType: imagebuilder.ImageType_AMI,
//   	ImageArchitecture: imagebuilder.ImageArchitecture_X86_64,
//   })
//
//   // Ubuntu 22.04 AMI for ARM64
//   ubuntu2204Ami := imagebuilder.AmazonManagedImage_UbuntuServer2204(this, jsii.String("Ubuntu2204"), &AmazonManagedImageOptions{
//   	ImageType: imagebuilder.ImageType_AMI,
//   	ImageArchitecture: imagebuilder.ImageArchitecture_ARM64,
//   })
//
//   // Windows Server 2022 Full AMI
//   windows2022Ami := imagebuilder.AmazonManagedImage_WindowsServer2022Full(this, jsii.String("Windows2022"), &AmazonManagedImageOptions{
//   	ImageType: imagebuilder.ImageType_AMI,
//   	ImageArchitecture: imagebuilder.ImageArchitecture_X86_64,
//   })
//
//   // Use as base image in recipe
//   managedImageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("ManagedImageRecipe"), &ImageRecipeProps{
//   	BaseImage: amazonLinux2023Ami.ToBaseImage(),
//   })
//
// Experimental.
type AmazonManagedImageOptions struct {
	// The architecture of the Amazon-managed image.
	// Experimental.
	ImageArchitecture ImageArchitecture `field:"required" json:"imageArchitecture" yaml:"imageArchitecture"`
	// The type of the Amazon-managed image.
	// Experimental.
	ImageType ImageType `field:"required" json:"imageType" yaml:"imageType"`
	// The version of the Amazon-managed image.
	// Default: x.x.x
	//
	// Experimental.
	ImageVersion *string `field:"optional" json:"imageVersion" yaml:"imageVersion"`
}

