package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstepfunctions"
)

// Collection of grant methods for a IStateMachineRef.
type StateMachineGrants interface {
	Resource() interfacesawsstepfunctions.IStateMachineRef
	// Grant the given identity custom permissions.
	Actions(identity awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to start an execution of this state machine.
	Execution(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to read results from state machine.
	Read(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permission to redrive the execution of the state machine.
	RedriveExecution(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to start an execution of this state machine.
	StartExecution(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to start a synchronous execution of this state machine.
	StartSyncExecution(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity task response permissions on a state machine.
	TaskResponse(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for StateMachineGrants
type jsiiProxy_StateMachineGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_StateMachineGrants) Resource() interfacesawsstepfunctions.IStateMachineRef {
	var returns interfacesawsstepfunctions.IStateMachineRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


func (s *jsiiProxy_StateMachineGrants) Actions(identity awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := s.validateActionsParameters(identity); err != nil {
		panic(err)
	}
	args := []interface{}{identity}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"actions",
		args,
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineGrants) Execution(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := s.validateExecutionParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"execution",
		args,
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineGrants) Read(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"read",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineGrants) RedriveExecution(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateRedriveExecutionParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"redriveExecution",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineGrants) StartExecution(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateStartExecutionParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"startExecution",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineGrants) StartSyncExecution(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateStartSyncExecutionParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"startSyncExecution",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineGrants) TaskResponse(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateTaskResponseParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"taskResponse",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

