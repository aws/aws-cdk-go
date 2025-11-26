package awsimagebuilderalpha


// Properties for an EC2 Image Builder AWS-managed component.
//
// Example:
//   imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("AwsManagedImageRecipe"), &ImageRecipeProps{
//   	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
//   	Components: []ComponentConfiguration{
//   		&ComponentConfiguration{
//   			Component: imagebuilder.AwsManagedComponent_UpdateOS(this, jsii.String("UpdateOS"), &AwsManagedComponentAttributes{
//   				Platform: imagebuilder.Platform_LINUX,
//   			}),
//   		},
//   		&ComponentConfiguration{
//   			Component: imagebuilder.AwsManagedComponent_AwsCliV2(this, jsii.String("AwsCli"), &AwsManagedComponentAttributes{
//   				Platform: imagebuilder.Platform_LINUX,
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type AwsManagedComponentAttributes struct {
	// The name of the AWS-managed component.
	//
	// The name of the AWS-managed component. This is a required attribute when using the
	// `this.fromAwsManagedComponentAttributes()` method. This parameter should not be provided when
	// using the pre-defined managed component methods, such as `AwsManagedComponent.updateOS()` and
	// `AwsManagedComponent.reboot()`.
	// Default: - none if using  the pre-defined managed component methods, otherwise a platform is required when using
	//   `this.fromAwsManagedComponentAttributes()`
	//
	// Experimental.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The version of the AWS-managed component.
	// Default: - the latest version of the component, x.x.x
	//
	// Experimental.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
	// The platform of the AWS-managed component.
	//
	// This is a required attribute when using the pre-defined managed
	// component methods, such as `AwsManagedComponent.updateOS()` and `AwsManagedComponent.reboot()`. This parameter
	// should not be provided when using the `this.fromAwsManagedComponentAttributes()` method.
	// Default: - none if using `this.fromAwsManagedComponentAttributes()`, otherwise a platform is
	// required when using the pre-defined managed component methods.
	//
	// Experimental.
	Platform Platform `field:"optional" json:"platform" yaml:"platform"`
}

