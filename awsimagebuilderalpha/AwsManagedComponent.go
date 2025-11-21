package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for working with AWS-managed components.
//
// Example:
//   // Install AWS CLI v2
//   awsCliComponent := imagebuilder.AwsManagedComponent_AwsCliV2(this, jsii.String("AwsCli"), &AwsManagedComponentAttributes{
//   	Platform: imagebuilder.Platform_LINUX,
//   })
//
//   // Update the operating system
//   updateComponent := imagebuilder.AwsManagedComponent_UpdateOS(this, jsii.String("UpdateOS"), &AwsManagedComponentAttributes{
//   	Platform: imagebuilder.Platform_LINUX,
//   })
//
//   // Reference any AWS-managed component by name
//   customAwsComponent := imagebuilder.AwsManagedComponent_FromAwsManagedComponentName(this, jsii.String("CloudWatchAgent"), jsii.String("amazon-cloudwatch-agent-linux"))
//
// Experimental.
type AwsManagedComponent interface {
}

// The jsii proxy struct for AwsManagedComponent
type jsiiProxy_AwsManagedComponent struct {
	_ byte // padding
}

// Experimental.
func NewAwsManagedComponent_Override(a AwsManagedComponent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		nil, // no parameters
		a,
	)
}

// Imports the AWS CLI v2 AWS-managed component.
// Experimental.
func AwsManagedComponent_AwsCliV2(scope constructs.Construct, id *string, attrs *AwsManagedComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAwsManagedComponent_AwsCliV2Parameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		"awsCliV2",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an AWS-managed component from its attributes.
// Experimental.
func AwsManagedComponent_FromAwsManagedComponentAttributes(scope constructs.Construct, id *string, attrs *AwsManagedComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAwsManagedComponent_FromAwsManagedComponentAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		"fromAwsManagedComponentAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an AWS-managed component from its name.
// Experimental.
func AwsManagedComponent_FromAwsManagedComponentName(scope constructs.Construct, id *string, awsManagedComponentName *string) IComponent {
	_init_.Initialize()

	if err := validateAwsManagedComponent_FromAwsManagedComponentNameParameters(scope, id, awsManagedComponentName); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		"fromAwsManagedComponentName",
		[]interface{}{scope, id, awsManagedComponentName},
		&returns,
	)

	return returns
}

// Imports the hello world AWS-managed component.
// Experimental.
func AwsManagedComponent_HelloWorld(scope constructs.Construct, id *string, attrs *AwsManagedComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAwsManagedComponent_HelloWorldParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		"helloWorld",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports the Python 3 AWS-managed component.
// Experimental.
func AwsManagedComponent_Python3(scope constructs.Construct, id *string, attrs *AwsManagedComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAwsManagedComponent_Python3Parameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		"python3",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports the reboot AWS-managed component.
// Experimental.
func AwsManagedComponent_Reboot(scope constructs.Construct, id *string, attrs *AwsManagedComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAwsManagedComponent_RebootParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		"reboot",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports the STIG hardening AWS-managed component.
// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/ib-stig.html
//
// Experimental.
func AwsManagedComponent_StigBuild(scope constructs.Construct, id *string, attrs *AwsManagedComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAwsManagedComponent_StigBuildParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		"stigBuild",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports the OS update AWS-managed component.
// Experimental.
func AwsManagedComponent_UpdateOS(scope constructs.Construct, id *string, attrs *AwsManagedComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAwsManagedComponent_UpdateOSParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		"updateOS",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

