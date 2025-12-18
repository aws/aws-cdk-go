package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for working with Amazon-managed components.
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
type AmazonManagedComponent interface {
}

// The jsii proxy struct for AmazonManagedComponent
type jsiiProxy_AmazonManagedComponent struct {
	_ byte // padding
}

// Experimental.
func NewAmazonManagedComponent_Override(a AmazonManagedComponent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		nil, // no parameters
		a,
	)
}

// Imports the AWS CLI v2 Amazon-managed component.
// Experimental.
func AmazonManagedComponent_AwsCliV2(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) IComponent {
	_init_.Initialize()

	if err := validateAmazonManagedComponent_AwsCliV2Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		"awsCliV2",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports an Amazon-managed component from its attributes.
// Experimental.
func AmazonManagedComponent_FromAmazonManagedComponentAttributes(scope constructs.Construct, id *string, attrs *AmazonManagedComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAmazonManagedComponent_FromAmazonManagedComponentAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		"fromAmazonManagedComponentAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an Amazon-managed component from its name.
// Experimental.
func AmazonManagedComponent_FromAmazonManagedComponentName(scope constructs.Construct, id *string, amazonManagedComponentName *string) IComponent {
	_init_.Initialize()

	if err := validateAmazonManagedComponent_FromAmazonManagedComponentNameParameters(scope, id, amazonManagedComponentName); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		"fromAmazonManagedComponentName",
		[]interface{}{scope, id, amazonManagedComponentName},
		&returns,
	)

	return returns
}

// Imports the hello world Amazon-managed component.
// Experimental.
func AmazonManagedComponent_HelloWorld(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) IComponent {
	_init_.Initialize()

	if err := validateAmazonManagedComponent_HelloWorldParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		"helloWorld",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the Python 3 Amazon-managed component.
// Experimental.
func AmazonManagedComponent_Python3(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) IComponent {
	_init_.Initialize()

	if err := validateAmazonManagedComponent_Python3Parameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		"python3",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the reboot Amazon-managed component.
// Experimental.
func AmazonManagedComponent_Reboot(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) IComponent {
	_init_.Initialize()

	if err := validateAmazonManagedComponent_RebootParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		"reboot",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the STIG hardening Amazon-managed component.
// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/ib-stig.html
//
// Experimental.
func AmazonManagedComponent_StigBuild(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) IComponent {
	_init_.Initialize()

	if err := validateAmazonManagedComponent_StigBuildParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		"stigBuild",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

// Imports the OS update Amazon-managed component.
// Experimental.
func AmazonManagedComponent_UpdateOs(scope constructs.Construct, id *string, opts *AmazonManagedComponentOptions) IComponent {
	_init_.Initialize()

	if err := validateAmazonManagedComponent_UpdateOsParameters(scope, id, opts); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedComponent",
		"updateOs",
		[]interface{}{scope, id, opts},
		&returns,
	)

	return returns
}

