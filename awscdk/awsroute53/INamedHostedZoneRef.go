package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53"
)

type INamedHostedZoneRef interface {
	interfacesawsroute53.IHostedZoneRef
	Name() *string
}

// The jsii proxy for INamedHostedZoneRef
type jsiiProxy_INamedHostedZoneRef struct {
	internal.Type__interfacesawsroute53IHostedZoneRef
}

func (j *jsiiProxy_INamedHostedZoneRef) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

