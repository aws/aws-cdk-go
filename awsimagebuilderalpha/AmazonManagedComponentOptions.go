package awsimagebuilderalpha


// Options for selecting a predefined Amazon-managed image.
//
// Example:
//   imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("AmazonManagedImageRecipe"), &ImageRecipeProps{
//   	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
//   	Components: []ComponentConfiguration{
//   		&ComponentConfiguration{
//   			Component: imagebuilder.AmazonManagedComponent_UpdateOs(this, jsii.String("UpdateOS"), &AmazonManagedComponentOptions{
//   				Platform: imagebuilder.Platform_LINUX,
//   			}),
//   		},
//   		&ComponentConfiguration{
//   			Component: imagebuilder.AmazonManagedComponent_AwsCliV2(this, jsii.String("AwsCli"), &AmazonManagedComponentOptions{
//   				Platform: imagebuilder.Platform_LINUX,
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type AmazonManagedComponentOptions struct {
	// The platform of the Amazon-managed component.
	// Experimental.
	Platform Platform `field:"required" json:"platform" yaml:"platform"`
	// The version of the Amazon-managed component.
	// Default: - the latest version of the component, x.x.x
	//
	// Experimental.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
}

