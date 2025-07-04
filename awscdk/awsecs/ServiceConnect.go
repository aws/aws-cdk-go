package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// ServiceConnect ValueObjectClass having by ContainerDefinition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appProtocol appProtocol
//
//   serviceConnect := awscdk.Aws_ecs.NewServiceConnect(awscdk.Aws_ecs.NetworkMode_NONE, &PortMapping{
//   	ContainerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	AppProtocol: appProtocol,
//   	ContainerPortRange: jsii.String("containerPortRange"),
//   	HostPort: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Protocol: awscdk.*Aws_ecs.Protocol_TCP,
//   })
//
type ServiceConnect interface {
	// The networking mode to use for the containers in the task.
	Networkmode() NetworkMode
	// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
	Portmapping() *PortMapping
	// Judge parameters can be serviceconnect logick.
	//
	// If parameters can be serviceConnect return true.
	IsServiceConnect() *bool
	// Judge serviceconnect parametes are valid.
	//
	// If invalid, throw Error.
	Validate()
}

// The jsii proxy struct for ServiceConnect
type jsiiProxy_ServiceConnect struct {
	_ byte // padding
}

func (j *jsiiProxy_ServiceConnect) Networkmode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkmode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceConnect) Portmapping() *PortMapping {
	var returns *PortMapping
	_jsii_.Get(
		j,
		"portmapping",
		&returns,
	)
	return returns
}


func NewServiceConnect(networkmode NetworkMode, pm *PortMapping) ServiceConnect {
	_init_.Initialize()

	if err := validateNewServiceConnectParameters(networkmode, pm); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceConnect{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ServiceConnect",
		[]interface{}{networkmode, pm},
		&j,
	)

	return &j
}

func NewServiceConnect_Override(s ServiceConnect, networkmode NetworkMode, pm *PortMapping) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ServiceConnect",
		[]interface{}{networkmode, pm},
		s,
	)
}

func (s *jsiiProxy_ServiceConnect) IsServiceConnect() *bool {
	var returns *bool

	_jsii_.Invoke(
		s,
		"isServiceConnect",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceConnect) Validate() {
	_jsii_.InvokeVoid(
		s,
		"validate",
		nil, // no parameters
	)
}

