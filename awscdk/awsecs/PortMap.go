package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// PortMap ValueObjectClass having by ContainerDefinition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appProtocol appProtocol
//
//   portMap := awscdk.Aws_ecs.NewPortMap(awscdk.Aws_ecs.NetworkMode_NONE, &PortMapping{
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
type PortMap interface {
	// The networking mode to use for the containers in the task.
	Networkmode() NetworkMode
	// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
	Portmapping() *PortMapping
	// validate invalid portmapping and networkmode parameters.
	//
	// throw Error when invalid parameters.
	Validate()
}

// The jsii proxy struct for PortMap
type jsiiProxy_PortMap struct {
	_ byte // padding
}

func (j *jsiiProxy_PortMap) Networkmode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkmode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PortMap) Portmapping() *PortMapping {
	var returns *PortMapping
	_jsii_.Get(
		j,
		"portmapping",
		&returns,
	)
	return returns
}


func NewPortMap(networkmode NetworkMode, pm *PortMapping) PortMap {
	_init_.Initialize()

	if err := validateNewPortMapParameters(networkmode, pm); err != nil {
		panic(err)
	}
	j := jsiiProxy_PortMap{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.PortMap",
		[]interface{}{networkmode, pm},
		&j,
	)

	return &j
}

func NewPortMap_Override(p PortMap, networkmode NetworkMode, pm *PortMapping) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.PortMap",
		[]interface{}{networkmode, pm},
		p,
	)
}

func (p *jsiiProxy_PortMap) Validate() {
	_jsii_.InvokeVoid(
		p,
		"validate",
		nil, // no parameters
	)
}

