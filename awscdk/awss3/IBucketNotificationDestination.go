package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Implemented by constructs that can be used as bucket notification destinations.
type IBucketNotificationDestination interface {
	// Registers this resource to receive notifications for the specified bucket.
	//
	// This method will only be called once for each destination/bucket
	// pair and the result will be cached, so there is no need to implement
	// idempotency in each destination.
	Bind(scope constructs.Construct, bucket interfacesawss3.IBucketRef) *BucketNotificationDestinationConfig
}

// The jsii proxy for IBucketNotificationDestination
type jsiiProxy_IBucketNotificationDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IBucketNotificationDestination) Bind(scope constructs.Construct, bucket interfacesawss3.IBucketRef) *BucketNotificationDestinationConfig {
	if err := i.validateBindParameters(scope, bucket); err != nil {
		panic(err)
	}
	var returns *BucketNotificationDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, bucket},
		&returns,
	)

	return returns
}

