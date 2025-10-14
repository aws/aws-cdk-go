package awsiotsitewise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssetModel.
// Experimental.
type IAssetModelRef interface {
	constructs.IConstruct
	// A reference to a AssetModel resource.
	// Experimental.
	AssetModelRef() *AssetModelReference
}

// The jsii proxy for IAssetModelRef
type jsiiProxy_IAssetModelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAssetModelRef) AssetModelRef() *AssetModelReference {
	var returns *AssetModelReference
	_jsii_.Get(
		j,
		"assetModelRef",
		&returns,
	)
	return returns
}

