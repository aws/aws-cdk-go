package interfacesawstimestream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InfluxDBInstance.
// Experimental.
type IInfluxDBInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InfluxDBInstance resource.
	// Experimental.
	InfluxDbInstanceRef() *InfluxDBInstanceReference
}

// The jsii proxy for IInfluxDBInstanceRef
type jsiiProxy_IInfluxDBInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IInfluxDBInstanceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IInfluxDBInstanceRef) InfluxDbInstanceRef() *InfluxDBInstanceReference {
	var returns *InfluxDBInstanceReference
	_jsii_.Get(
		j,
		"influxDbInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInfluxDBInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInfluxDBInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

