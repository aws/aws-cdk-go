package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/internal"
)

// An EC2 Image Builder Image.
// Experimental.
type IImage interface {
	awscdk.IResource
	// Grant custom actions to the given grantee for the image.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants the default permissions for building an image to the provided execution role.
	// Experimental.
	GrantDefaultExecutionRolePermissions(grantee awsiam.IGrantable) *[]awsiam.Grant
	// Grant read permissions to the given grantee for the image.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Converts the image to a BaseImage, to use as the parent image in an image recipe.
	// Experimental.
	ToBaseImage() BaseImage
	// Converts the image to a ContainerBaseImage, to use as the parent image in a container recipe.
	// Experimental.
	ToContainerBaseImage() BaseContainerImage
	// The ARN of the image.
	// Experimental.
	ImageArn() *string
	// The name of the image.
	// Experimental.
	ImageName() *string
	// The version of the image.
	// Experimental.
	ImageVersion() *string
}

// The jsii proxy for IImage
type jsiiProxy_IImage struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IImage) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImage) GrantDefaultExecutionRolePermissions(grantee awsiam.IGrantable) *[]awsiam.Grant {
	if err := i.validateGrantDefaultExecutionRolePermissionsParameters(grantee); err != nil {
		panic(err)
	}
	var returns *[]awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDefaultExecutionRolePermissions",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImage) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImage) ToBaseImage() BaseImage {
	var returns BaseImage

	_jsii_.Invoke(
		i,
		"toBaseImage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImage) ToContainerBaseImage() BaseContainerImage {
	var returns BaseContainerImage

	_jsii_.Invoke(
		i,
		"toContainerBaseImage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IImage) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImage) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImage) ImageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageVersion",
		&returns,
	)
	return returns
}

