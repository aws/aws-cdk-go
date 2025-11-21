package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs"
)

// Collection of grant methods for a IClusterRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clusterRef IClusterRef
//
//   clusterGrants := awscdk.Aws_ecs.ClusterGrants_FromCluster(clusterRef)
//
type ClusterGrants interface {
	Resource() interfacesawsecs.IClusterRef
	// Grants an ECS Task Protection API permission to the specified grantee.
	//
	// This method provides a streamlined way to assign the 'ecs:UpdateTaskProtection'
	// permission, enabling the grantee to manage task protection in the ECS cluster.
	TaskProtection(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for ClusterGrants
type jsiiProxy_ClusterGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterGrants) Resource() interfacesawsecs.IClusterRef {
	var returns interfacesawsecs.IClusterRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for ClusterGrants.
func ClusterGrants_FromCluster(resource interfacesawsecs.IClusterRef) ClusterGrants {
	_init_.Initialize()

	if err := validateClusterGrants_FromClusterParameters(resource); err != nil {
		panic(err)
	}
	var returns ClusterGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ClusterGrants",
		"fromCluster",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterGrants) TaskProtection(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateTaskProtectionParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"taskProtection",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

