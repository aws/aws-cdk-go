// The CDK Construct Library for AWS::S3ObjectLambda
package awscdks3objectlambdaalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdks3objectlambdaalpha/v2/internal"
)

// The interface that represents the AccessPoint resource.
// Experimental.
type IAccessPoint interface {
	awscdk.IResource
	// The virtual hosted-style URL of an S3 object through this access point.
	//
	// Specify `regional: false` at the options for non-regional URL.
	//
	// Returns: an ObjectS3Url token.
	// Experimental.
	VirtualHostedUrlForObject(key *string, options *awss3.VirtualHostedStyleUrlOptions) *string
	// The ARN of the access point.
	// Experimental.
	AccessPointArn() *string
	// The creation data of the access point.
	// Experimental.
	AccessPointCreationDate() *string
	// The IPv4 DNS name of the access point.
	// Experimental.
	DomainName() *string
	// The regional domain name of the access point.
	// Experimental.
	RegionalDomainName() *string
}

// The jsii proxy for IAccessPoint
type jsiiProxy_IAccessPoint struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAccessPoint) VirtualHostedUrlForObject(key *string, options *awss3.VirtualHostedStyleUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"virtualHostedUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAccessPoint) AccessPointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) AccessPointCreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointCreationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

