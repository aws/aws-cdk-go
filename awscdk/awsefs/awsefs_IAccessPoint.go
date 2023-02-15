package awsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
)

// Represents an EFS AccessPoint.
type IAccessPoint interface {
	awscdk.IResource
	// The ARN of the AccessPoint.
	AccessPointArn() *string
	// The ID of the AccessPoint.
	AccessPointId() *string
	// The EFS file system.
	FileSystem() IFileSystem
}

// The jsii proxy for IAccessPoint
type jsiiProxy_IAccessPoint struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAccessPoint) AccessPointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) AccessPointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) FileSystem() IFileSystem {
	var returns IFileSystem
	_jsii_.Get(
		j,
		"fileSystem",
		&returns,
	)
	return returns
}

