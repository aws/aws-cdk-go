package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The interface for BaseService.
// Experimental.
type IBaseService interface {
	IService
	// The cluster that hosts the service.
	// Experimental.
	Cluster() ICluster
}

// The jsii proxy for IBaseService
type jsiiProxy_IBaseService struct {
	jsiiProxy_IService
}

func (j *jsiiProxy_IBaseService) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

