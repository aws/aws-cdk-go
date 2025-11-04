package awstimestream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InfluxDBInstance.
// Experimental.
type IInfluxDBInstanceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InfluxDBInstance resource.
	// Experimental.
	InfluxDbInstanceRef() *InfluxDBInstanceReference
}

// The jsii proxy for IInfluxDBInstanceRef
type jsiiProxy_IInfluxDBInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IInfluxDBInstanceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

