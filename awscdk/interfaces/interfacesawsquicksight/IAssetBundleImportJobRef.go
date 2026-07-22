package interfacesawsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssetBundleImportJob.
// Experimental.
type IAssetBundleImportJobRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AssetBundleImportJob resource.
	// Experimental.
	AssetBundleImportJobRef() *AssetBundleImportJobReference
}

// The jsii proxy for IAssetBundleImportJobRef
type jsiiProxy_IAssetBundleImportJobRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAssetBundleImportJobRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAssetBundleImportJobRef) AssetBundleImportJobRef() *AssetBundleImportJobReference {
	var returns *AssetBundleImportJobReference
	_jsii_.Get(
		j,
		"assetBundleImportJobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssetBundleImportJobRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssetBundleImportJobRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

