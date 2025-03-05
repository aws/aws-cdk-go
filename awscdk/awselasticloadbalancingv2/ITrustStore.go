package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
)

// Represents a Trust Store.
type ITrustStore interface {
	awscdk.IResource
	// The ARN of the trust store.
	TrustStoreArn() *string
	// The name of the trust store.
	TrustStoreName() *string
}

// The jsii proxy for ITrustStore
type jsiiProxy_ITrustStore struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITrustStore) TrustStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStore) TrustStoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreName",
		&returns,
	)
	return returns
}

