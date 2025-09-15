package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OriginAccessControl.
// Experimental.
type IOriginAccessControlRef interface {
	constructs.IConstruct
	// A reference to a OriginAccessControl resource.
	// Experimental.
	OriginAccessControlRef() *OriginAccessControlReference
}

// The jsii proxy for IOriginAccessControlRef
type jsiiProxy_IOriginAccessControlRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOriginAccessControlRef) OriginAccessControlRef() *OriginAccessControlReference {
	var returns *OriginAccessControlReference
	_jsii_.Get(
		j,
		"originAccessControlRef",
		&returns,
	)
	return returns
}

