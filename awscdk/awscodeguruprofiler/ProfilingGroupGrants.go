package awscodeguruprofiler

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodeguruprofiler"
)

// Collection of grant methods for a IProfilingGroupRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var profilingGroupRef IProfilingGroupRef
//
//   profilingGroupGrants := awscdk.Aws_codeguruprofiler.ProfilingGroupGrants_FromProfilingGroup(profilingGroupRef)
//
type ProfilingGroupGrants interface {
	Resource() interfacesawscodeguruprofiler.IProfilingGroupRef
	// Grant access to publish profiling information to the Profiling Group to the given identity.
	//
	// This will grant the following permissions:
	//
	//  - codeguru-profiler:ConfigureAgent
	// - codeguru-profiler:PostAgentProfile.
	Publish(grantee awsiam.IGrantable) awsiam.Grant
	// Grant access to read profiling information from the Profiling Group to the given identity.
	//
	// This will grant the following permissions:
	//
	//  - codeguru-profiler:GetProfile
	// - codeguru-profiler:DescribeProfilingGroup.
	Read(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for ProfilingGroupGrants
type jsiiProxy_ProfilingGroupGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_ProfilingGroupGrants) Resource() interfacesawscodeguruprofiler.IProfilingGroupRef {
	var returns interfacesawscodeguruprofiler.IProfilingGroupRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for ProfilingGroupGrants.
func ProfilingGroupGrants_FromProfilingGroup(resource interfacesawscodeguruprofiler.IProfilingGroupRef) ProfilingGroupGrants {
	_init_.Initialize()

	if err := validateProfilingGroupGrants_FromProfilingGroupParameters(resource); err != nil {
		panic(err)
	}
	var returns ProfilingGroupGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codeguruprofiler.ProfilingGroupGrants",
		"fromProfilingGroup",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProfilingGroupGrants) Publish(grantee awsiam.IGrantable) awsiam.Grant {
	if err := p.validatePublishParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		p,
		"publish",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProfilingGroupGrants) Read(grantee awsiam.IGrantable) awsiam.Grant {
	if err := p.validateReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		p,
		"read",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

