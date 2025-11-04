package awsiotsitewise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Asset.
// Experimental.
type IAssetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Asset resource.
	// Experimental.
	AssetRef() *AssetReference
}

// The jsii proxy for IAssetRef
type jsiiProxy_IAssetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAssetRef) AssetRef() *AssetReference {
	var returns *AssetReference
	_jsii_.Get(
		j,
		"assetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

