package awsmediapackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackagingConfiguration.
// Experimental.
type IPackagingConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PackagingConfiguration resource.
	// Experimental.
	PackagingConfigurationRef() *PackagingConfigurationReference
}

// The jsii proxy for IPackagingConfigurationRef
type jsiiProxy_IPackagingConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPackagingConfigurationRef) PackagingConfigurationRef() *PackagingConfigurationReference {
	var returns *PackagingConfigurationReference
	_jsii_.Get(
		j,
		"packagingConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPackagingConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPackagingConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

