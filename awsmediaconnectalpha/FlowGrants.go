package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
)

// Collection of grant methods for a IFlowRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var flowRef IFlowRef
//
//   flowGrants := mediaconnect_alpha.FlowGrants_FromFlow(flowRef)
//
// Experimental.
type FlowGrants interface {
	// Experimental.
	Resource() interfacesawsmediaconnect.IFlowRef
	// Grant the given identity custom permissions.
	// Experimental.
	Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant
	// Grant permissions to start this flow.
	// Experimental.
	Start(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to stop this flow.
	// Experimental.
	Stop(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for FlowGrants
type jsiiProxy_FlowGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_FlowGrants) Resource() interfacesawsmediaconnect.IFlowRef {
	var returns interfacesawsmediaconnect.IFlowRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for FlowGrants.
// Experimental.
func FlowGrants_FromFlow(resource interfacesawsmediaconnect.IFlowRef) FlowGrants {
	_init_.Initialize()

	if err := validateFlowGrants_FromFlowParameters(resource); err != nil {
		panic(err)
	}
	var returns FlowGrants

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.FlowGrants",
		"fromFlow",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowGrants) Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant {
	if err := f.validateActionsParameters(grantee, actions, options); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"actions",
		[]interface{}{grantee, actions, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowGrants) Start(grantee awsiam.IGrantable) awsiam.Grant {
	if err := f.validateStartParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"start",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowGrants) Stop(grantee awsiam.IGrantable) awsiam.Grant {
	if err := f.validateStopParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"stop",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

