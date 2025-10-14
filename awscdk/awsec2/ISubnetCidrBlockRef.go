package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubnetCidrBlock.
// Experimental.
type ISubnetCidrBlockRef interface {
	constructs.IConstruct
	// A reference to a SubnetCidrBlock resource.
	// Experimental.
	SubnetCidrBlockRef() *SubnetCidrBlockReference
}

// The jsii proxy for ISubnetCidrBlockRef
type jsiiProxy_ISubnetCidrBlockRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubnetCidrBlockRef) SubnetCidrBlockRef() *SubnetCidrBlockReference {
	var returns *SubnetCidrBlockReference
	_jsii_.Get(
		j,
		"subnetCidrBlockRef",
		&returns,
	)
	return returns
}

