package pipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// Features that the Stage needs from its environment.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type IStageHost interface {
	// Make sure all the assets from the given manifest are published.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PublishAsset(command *AssetPublishingCommand)
	// Return the Artifact the given stack has to emit its outputs into, if any.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackOutputArtifact(stackArtifactId *string) awscodepipeline.Artifact
}

// The jsii proxy for IStageHost
type jsiiProxy_IStageHost struct {
	_ byte // padding
}

func (i *jsiiProxy_IStageHost) PublishAsset(command *AssetPublishingCommand) {
	if err := i.validatePublishAssetParameters(command); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"publishAsset",
		[]interface{}{command},
	)
}

func (i *jsiiProxy_IStageHost) StackOutputArtifact(stackArtifactId *string) awscodepipeline.Artifact {
	if err := i.validateStackOutputArtifactParameters(stackArtifactId); err != nil {
		panic(err)
	}
	var returns awscodepipeline.Artifact

	_jsii_.Invoke(
		i,
		"stackOutputArtifact",
		[]interface{}{stackArtifactId},
		&returns,
	)

	return returns
}

