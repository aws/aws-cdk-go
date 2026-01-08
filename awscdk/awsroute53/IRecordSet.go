package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

// A record set.
type IRecordSet interface {
	interfacesawsroute53.IRecordSetRef
	awscdk.IResource
	// The domain name of the record.
	DomainName() *string
}

// The jsii proxy for IRecordSet
type jsiiProxy_IRecordSet struct {
	internal.Type__interfacesawsroute53IRecordSetRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRecordSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IRecordSet) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecordSet) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecordSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecordSet) RecordSetRef() *interfacesawsroute53.RecordSetReference {
	var returns *interfacesawsroute53.RecordSetReference
	_jsii_.Get(
		j,
		"recordSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecordSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

