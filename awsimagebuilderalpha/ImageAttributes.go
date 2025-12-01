package awsimagebuilderalpha


// Properties for an EC2 Image Builder image.
//
// Example:
//   // Import by name
//   existingImageByName := imagebuilder.Image_FromImageName(this, jsii.String("ExistingImageByName"), jsii.String("my-existing-image"))
//
//   // Import by ARN
//   existingImageByArn := imagebuilder.Image_FromImageArn(this, jsii.String("ExistingImageByArn"), jsii.String("arn:aws:imagebuilder:us-east-1:123456789012:image/imported-image/1.0.0"))
//
//   // Import by attributes
//   existingImageByAttributes := imagebuilder.Image_FromImageAttributes(this, jsii.String("ExistingImageByAttributes"), &ImageAttributes{
//   	ImageName: jsii.String("shared-base-image"),
//   	ImageVersion: jsii.String("2024.11.25"),
//   })
//
//   // Grant permissions to imported images
//   role := iam.NewRole(this, jsii.String("ImageAccessRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//
//   existingImageByName.GrantRead(role)
//   existingImageByArn.Grant(role, jsii.String("imagebuilder:GetImage"), jsii.String("imagebuilder:ListImagePackages"))
//
// Experimental.
type ImageAttributes struct {
	// The ARN of the image.
	// Default: - derived from the imageName.
	//
	// Experimental.
	ImageArn *string `field:"optional" json:"imageArn" yaml:"imageArn"`
	// The name of the image.
	// Default: - derived from the imageArn.
	//
	// Experimental.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// The version of the image.
	// Default: - the latest version of the image, x.x.x
	//
	// Experimental.
	ImageVersion *string `field:"optional" json:"imageVersion" yaml:"imageVersion"`
}

