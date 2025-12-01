package awsimagebuilderalpha


// Attributes for importing an Amazon-managed image by name (and optionally a version).
//
// Example:
//   // Import by name
//   managedImageByName := imagebuilder.AmazonManagedImage_FromAmazonManagedImageName(this, jsii.String("ManagedImageByName"), jsii.String("amazon-linux-2023-x86"))
//
//   // Import by attributes with specific version
//   managedImageByAttributes := imagebuilder.AmazonManagedImage_FromAmazonManagedImageAttributes(this, jsii.String("ManagedImageByAttributes"), &AmazonManagedImageAttributes{
//   	ImageName: jsii.String("ubuntu-server-22-lts-x86"),
//   	ImageVersion: jsii.String("2024.11.25"),
//   })
//
// Experimental.
type AmazonManagedImageAttributes struct {
	// The name of the Amazon-managed image.
	//
	// The provided name must be normalized by converting all alphabetical
	// characters to lowercase, and replacing all spaces and underscores with hyphens.
	// Experimental.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The version of the Amazon-managed image.
	// Default: x.x.x
	//
	// Experimental.
	ImageVersion *string `field:"optional" json:"imageVersion" yaml:"imageVersion"`
}

