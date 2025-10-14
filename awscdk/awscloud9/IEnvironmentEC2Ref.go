package awscloud9

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloud9/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentEC2.
// Experimental.
type IEnvironmentEC2Ref interface {
	constructs.IConstruct
	// A reference to a EnvironmentEC2 resource.
	// Experimental.
	EnvironmentEc2Ref() *EnvironmentEC2Reference
}

// The jsii proxy for IEnvironmentEC2Ref
type jsiiProxy_IEnvironmentEC2Ref struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEnvironmentEC2Ref) EnvironmentEc2Ref() *EnvironmentEC2Reference {
	var returns *EnvironmentEC2Reference
	_jsii_.Get(
		j,
		"environmentEc2Ref",
		&returns,
	)
	return returns
}

