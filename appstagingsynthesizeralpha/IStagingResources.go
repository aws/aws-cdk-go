package appstagingsynthesizeralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/appstagingsynthesizeralpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Staging Resource interface.
// Experimental.
type IStagingResources interface {
	constructs.IConstruct
	// Return staging resource information for a docker asset.
	// Experimental.
	AddDockerImage(asset *awscdk.DockerImageAssetSource) *ImageStagingLocation
	// Return staging resource information for a file asset.
	// Experimental.
	AddFile(asset *awscdk.FileAssetSource) *FileStagingLocation
}

// The jsii proxy for IStagingResources
type jsiiProxy_IStagingResources struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IStagingResources) AddDockerImage(asset *awscdk.DockerImageAssetSource) *ImageStagingLocation {
	if err := i.validateAddDockerImageParameters(asset); err != nil {
		panic(err)
	}
	var returns *ImageStagingLocation

	_jsii_.Invoke(
		i,
		"addDockerImage",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStagingResources) AddFile(asset *awscdk.FileAssetSource) *FileStagingLocation {
	if err := i.validateAddFileParameters(asset); err != nil {
		panic(err)
	}
	var returns *FileStagingLocation

	_jsii_.Invoke(
		i,
		"addFile",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

