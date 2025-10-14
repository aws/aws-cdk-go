package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubnetGroup.
// Experimental.
type ISubnetGroupRef interface {
	constructs.IConstruct
	// A reference to a SubnetGroup resource.
	// Experimental.
	SubnetGroupRef() *SubnetGroupReference
}

// The jsii proxy for ISubnetGroupRef
type jsiiProxy_ISubnetGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubnetGroupRef) SubnetGroupRef() *SubnetGroupReference {
	var returns *SubnetGroupReference
	_jsii_.Get(
		j,
		"subnetGroupRef",
		&returns,
	)
	return returns
}

