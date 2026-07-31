package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
)

// Collection of grant methods for a IRouterInputRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var routerInputRef IRouterInputRef
//
//   routerInputGrants := mediaconnect_alpha.RouterInputGrants_FromRouterInput(routerInputRef)
//
// Experimental.
type RouterInputGrants interface {
	// Experimental.
	Resource() interfacesawsmediaconnect.IRouterInputRef
	// Grant the given identity custom permissions.
	// Experimental.
	Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant
	// Grant permissions to restart this router input.
	// Experimental.
	Restart(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to start this router input.
	// Experimental.
	Start(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to stop this router input.
	// Experimental.
	Stop(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for RouterInputGrants
type jsiiProxy_RouterInputGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_RouterInputGrants) Resource() interfacesawsmediaconnect.IRouterInputRef {
	var returns interfacesawsmediaconnect.IRouterInputRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for RouterInputGrants.
// Experimental.
func RouterInputGrants_FromRouterInput(resource interfacesawsmediaconnect.IRouterInputRef) RouterInputGrants {
	_init_.Initialize()

	if err := validateRouterInputGrants_FromRouterInputParameters(resource); err != nil {
		panic(err)
	}
	var returns RouterInputGrants

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputGrants",
		"fromRouterInput",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouterInputGrants) Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant {
	if err := r.validateActionsParameters(grantee, actions, options); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"actions",
		[]interface{}{grantee, actions, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouterInputGrants) Restart(grantee awsiam.IGrantable) awsiam.Grant {
	if err := r.validateRestartParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"restart",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouterInputGrants) Start(grantee awsiam.IGrantable) awsiam.Grant {
	if err := r.validateStartParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"start",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouterInputGrants) Stop(grantee awsiam.IGrantable) awsiam.Grant {
	if err := r.validateStopParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"stop",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

