package awscodeguruprofiler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodeguruprofiler"
	"github.com/aws/constructs-go/constructs/v10"
)

// IResource represents a Profiling Group.
type IProfilingGroup interface {
	interfacesawscodeguruprofiler.IProfilingGroupRef
	awscdk.IResource
	// Grant access to publish profiling information to the Profiling Group to the given identity.
	//
	// This will grant the following permissions:
	//
	//  - codeguru-profiler:ConfigureAgent
	// - codeguru-profiler:PostAgentProfile.
	GrantPublish(grantee awsiam.IGrantable) awsiam.Grant
	// Grant access to read profiling information from the Profiling Group to the given identity.
	//
	// This will grant the following permissions:
	//
	//  - codeguru-profiler:GetProfile
	// - codeguru-profiler:DescribeProfilingGroup.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the profiling group.
	ProfilingGroupArn() *string
	// The name of the profiling group.
	ProfilingGroupName() *string
}

// The jsii proxy for IProfilingGroup
type jsiiProxy_IProfilingGroup struct {
	internal.Type__interfacesawscodeguruprofilerIProfilingGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IProfilingGroup) GrantPublish(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPublishParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPublish",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProfilingGroup) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProfilingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IProfilingGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IProfilingGroup) ProfilingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profilingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilingGroup) ProfilingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profilingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilingGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilingGroup) ProfilingGroupRef() *interfacesawscodeguruprofiler.ProfilingGroupReference {
	var returns *interfacesawscodeguruprofiler.ProfilingGroupReference
	_jsii_.Get(
		j,
		"profilingGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

