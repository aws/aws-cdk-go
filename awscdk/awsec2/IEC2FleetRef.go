package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EC2Fleet.
// Experimental.
type IEC2FleetRef interface {
	constructs.IConstruct
	// A reference to a EC2Fleet resource.
	// Experimental.
	Ec2FleetRef() *EC2FleetReference
}

// The jsii proxy for IEC2FleetRef
type jsiiProxy_IEC2FleetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEC2FleetRef) Ec2FleetRef() *EC2FleetReference {
	var returns *EC2FleetReference
	_jsii_.Get(
		j,
		"ec2FleetRef",
		&returns,
	)
	return returns
}

