package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for working with Amazon-managed images.
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
type AmazonManagedImage interface {
}

// The jsii proxy struct for AmazonManagedImage
type jsiiProxy_AmazonManagedImage struct {
	_ byte // padding
}

// Experimental.
func NewAmazonManagedImage() AmazonManagedImage {
	_init_.Initialize()

	j := jsiiProxy_AmazonManagedImage{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAmazonManagedImage_Override(a AmazonManagedImage) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		nil, // no parameters
		a,
	)
}

// Imports the Amazon Linux 2 Amazon-managed image.
// See: https://gallery.ecr.aws/amazonlinux/amazonlinux
//
// Experimental.
func AmazonManagedImage_AmazonLinux2(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_AmazonLinux2Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"amazonLinux2",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Amazon Linux 2023 Amazon-managed image.
// See: https://gallery.ecr.aws/amazonlinux/amazonlinux
//
// Experimental.
func AmazonManagedImage_AmazonLinux2023(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_AmazonLinux2023Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"amazonLinux2023",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports an Amazon-managed image from its attributes.
// Experimental.
func AmazonManagedImage_FromAmazonManagedImageAttributes(scope constructs.Construct, id *string, attrs *AmazonManagedImageAttributes) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_FromAmazonManagedImageAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"fromAmazonManagedImageAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an Amazon-managed image from its name.
// Experimental.
func AmazonManagedImage_FromAmazonManagedImageName(scope constructs.Construct, id *string, amazonManagedImageName *string) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_FromAmazonManagedImageNameParameters(scope, id, amazonManagedImageName); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"fromAmazonManagedImageName",
		[]interface{}{scope, id, amazonManagedImageName},
		&returns,
	)

	return returns
}

// Imports the macOS 14 Amazon-managed image.
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-mac-instances.html
//
// Experimental.
func AmazonManagedImage_MacOS14(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_MacOS14Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"macOS14",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the macOS 15 Amazon-managed image.
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-mac-instances.html
//
// Experimental.
func AmazonManagedImage_MacOS15(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_MacOS15Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"macOS15",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Red Hat Enterprise Linux 10 Amazon-managed image.
// See: https://aws.amazon.com/partners/redhat/faqs
//
// Experimental.
func AmazonManagedImage_RedHatEnterpriseLinux10(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_RedHatEnterpriseLinux10Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"redHatEnterpriseLinux10",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the SUSE Linux Enterprise Server 15 Amazon-managed image.
// See: https://aws.amazon.com/linux/commercial-linux/faqs/
//
// Experimental.
func AmazonManagedImage_SuseLinuxEnterpriseServer15(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_SuseLinuxEnterpriseServer15Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"suseLinuxEnterpriseServer15",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Ubuntu 22.04 Amazon-managed image.
// See: https://hub.docker.com/_/ubuntu
//
// Experimental.
func AmazonManagedImage_UbuntuServer2204(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_UbuntuServer2204Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"ubuntuServer2204",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Ubuntu 24.04 Amazon-managed image.
// See: https://hub.docker.com/_/ubuntu
//
// Experimental.
func AmazonManagedImage_UbuntuServer2404(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_UbuntuServer2404Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"ubuntuServer2404",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Windows Server 2016 Core Amazon-managed image.
// See: https://hub.docker.com/r/microsoft/windows-servercore
//
// Experimental.
func AmazonManagedImage_WindowsServer2016Core(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_WindowsServer2016CoreParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"windowsServer2016Core",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Windows Server 2016 Full Amazon-managed image.
// See: https://docs.aws.amazon.com/ec2/latest/windows-ami-reference/index.html
//
// Experimental.
func AmazonManagedImage_WindowsServer2016Full(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_WindowsServer2016FullParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"windowsServer2016Full",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Windows Server 2019 Core Amazon-managed image.
// See: https://hub.docker.com/r/microsoft/windows-servercore
//
// Experimental.
func AmazonManagedImage_WindowsServer2019Core(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_WindowsServer2019CoreParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"windowsServer2019Core",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Windows Server 2019 Full Amazon-managed image.
// See: https://docs.aws.amazon.com/ec2/latest/windows-ami-reference/index.html
//
// Experimental.
func AmazonManagedImage_WindowsServer2019Full(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_WindowsServer2019FullParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"windowsServer2019Full",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Windows Server 2022 Core Amazon-managed image.
// See: https://docs.aws.amazon.com/ec2/latest/windows-ami-reference/index.html
//
// Experimental.
func AmazonManagedImage_WindowsServer2022Core(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_WindowsServer2022CoreParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"windowsServer2022Core",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Windows Server 2022 Full Amazon-managed image.
// See: https://docs.aws.amazon.com/ec2/latest/windows-ami-reference/index.html
//
// Experimental.
func AmazonManagedImage_WindowsServer2022Full(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_WindowsServer2022FullParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"windowsServer2022Full",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Windows Server 2025 Core Amazon-managed image.
// See: https://docs.aws.amazon.com/ec2/latest/windows-ami-reference/index.html
//
// Experimental.
func AmazonManagedImage_WindowsServer2025Core(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_WindowsServer2025CoreParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"windowsServer2025Core",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Windows Server 2025 Full Amazon-managed image.
// See: https://docs.aws.amazon.com/ec2/latest/windows-ami-reference/index.html
//
// Experimental.
func AmazonManagedImage_WindowsServer2025Full(scope constructs.Construct, id *string, opts *AmazonManagedImageOptions) IImage {
	_init_.Initialize()

	if err := validateAmazonManagedImage_WindowsServer2025FullParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		"windowsServer2025Full",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

