package awscdkelasticachealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkelasticachealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticache"
)

// Collection of grant methods for a IServerlessCacheRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import elasticache_alpha "github.com/aws/aws-cdk-go/awscdkelasticachealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serverlessCacheRef IServerlessCacheRef
//
//   serverlessCacheGrants := elasticache_alpha.ServerlessCacheGrants_FromServerlessCache(serverlessCacheRef)
//
// Experimental.
type ServerlessCacheGrants interface {
	// Experimental.
	Resource() interfacesawselasticache.IServerlessCacheRef
	// Grant connect permissions to the cache.
	// Experimental.
	Connect(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for ServerlessCacheGrants
type jsiiProxy_ServerlessCacheGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_ServerlessCacheGrants) Resource() interfacesawselasticache.IServerlessCacheRef {
	var returns interfacesawselasticache.IServerlessCacheRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for ServerlessCacheGrants.
// Experimental.
func ServerlessCacheGrants_FromServerlessCache(resource interfacesawselasticache.IServerlessCacheRef) ServerlessCacheGrants {
	_init_.Initialize()

	if err := validateServerlessCacheGrants_FromServerlessCacheParameters(resource); err != nil {
		panic(err)
	}
	var returns ServerlessCacheGrants

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-elasticache-alpha.ServerlessCacheGrants",
		"fromServerlessCache",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCacheGrants) Connect(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateConnectParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"connect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

