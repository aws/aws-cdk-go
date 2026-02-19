package interfacesawsmpa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmpa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApprovalTeam.
// Experimental.
type IApprovalTeamRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApprovalTeam resource.
	// Experimental.
	ApprovalTeamRef() *ApprovalTeamReference
}

// The jsii proxy for IApprovalTeamRef
type jsiiProxy_IApprovalTeamRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IApprovalTeamRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IApprovalTeamRef) ApprovalTeamRef() *ApprovalTeamReference {
	var returns *ApprovalTeamReference
	_jsii_.Get(
		j,
		"approvalTeamRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApprovalTeamRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApprovalTeamRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

