package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CidrCollection.
// Experimental.
type ICidrCollectionRef interface {
	constructs.IConstruct
	// A reference to a CidrCollection resource.
	// Experimental.
	CidrCollectionRef() *CidrCollectionReference
}

// The jsii proxy for ICidrCollectionRef
type jsiiProxy_ICidrCollectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICidrCollectionRef) CidrCollectionRef() *CidrCollectionReference {
	var returns *CidrCollectionReference
	_jsii_.Get(
		j,
		"cidrCollectionRef",
		&returns,
	)
	return returns
}

