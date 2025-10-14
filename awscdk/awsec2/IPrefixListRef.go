package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrefixList.
// Experimental.
type IPrefixListRef interface {
	constructs.IConstruct
	// A reference to a PrefixList resource.
	// Experimental.
	PrefixListRef() *PrefixListReference
}

// The jsii proxy for IPrefixListRef
type jsiiProxy_IPrefixListRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPrefixListRef) PrefixListRef() *PrefixListReference {
	var returns *PrefixListReference
	_jsii_.Get(
		j,
		"prefixListRef",
		&returns,
	)
	return returns
}

