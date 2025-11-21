package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
)

// Collection of grant methods for a IDistributionRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var distributionRef IDistributionRef
//
//   distributionGrants := awscdk.Aws_cloudfront.DistributionGrants_FromDistribution(distributionRef)
//
type DistributionGrants interface {
	Resource() interfacesawscloudfront.IDistributionRef
	// Grant to create invalidations for this bucket to an IAM principal (Role/Group/User).
	CreateInvalidation(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for DistributionGrants
type jsiiProxy_DistributionGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_DistributionGrants) Resource() interfacesawscloudfront.IDistributionRef {
	var returns interfacesawscloudfront.IDistributionRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for DistributionGrants.
func DistributionGrants_FromDistribution(resource interfacesawscloudfront.IDistributionRef) DistributionGrants {
	_init_.Initialize()

	if err := validateDistributionGrants_FromDistributionParameters(resource); err != nil {
		panic(err)
	}
	var returns DistributionGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.DistributionGrants",
		"fromDistribution",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributionGrants) CreateInvalidation(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateCreateInvalidationParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"createInvalidation",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

