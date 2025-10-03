package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InfluxDBInstance.
// Experimental.
type IInfluxDBInstanceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInfluxDBInstanceRef
type jsiiProxy_IInfluxDBInstanceRef struct {
	internal.Type__constructsIConstruct
}

