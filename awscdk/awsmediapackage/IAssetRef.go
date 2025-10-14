package awsmediapackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Asset.
// Experimental.
type IAssetRef interface {
	constructs.IConstruct
	// A reference to a Asset resource.
	// Experimental.
	AssetRef() *AssetReference
}

// The jsii proxy for IAssetRef
type jsiiProxy_IAssetRef struct {
	internal.Type__constructsIConstruct
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

