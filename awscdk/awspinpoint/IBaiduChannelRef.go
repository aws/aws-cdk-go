package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BaiduChannel.
// Experimental.
type IBaiduChannelRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a BaiduChannel resource.
	// Experimental.
	BaiduChannelRef() *BaiduChannelReference
}

// The jsii proxy for IBaiduChannelRef
type jsiiProxy_IBaiduChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IBaiduChannelRef) BaiduChannelRef() *BaiduChannelReference {
	var returns *BaiduChannelReference
	_jsii_.Get(
		j,
		"baiduChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaiduChannelRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaiduChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

