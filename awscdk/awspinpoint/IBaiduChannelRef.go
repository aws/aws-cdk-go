package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BaiduChannel.
// Experimental.
type IBaiduChannelRef interface {
	constructs.IConstruct
	// A reference to a BaiduChannel resource.
	// Experimental.
	BaiduChannelRef() *BaiduChannelReference
}

// The jsii proxy for IBaiduChannelRef
type jsiiProxy_IBaiduChannelRef struct {
	internal.Type__constructsIConstruct
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

