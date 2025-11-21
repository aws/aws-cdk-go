package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
)

// Collection of grant methods for a IDeliveryStreamRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryStreamRef IDeliveryStreamRef
//
//   deliveryStreamGrants := awscdk.Aws_kinesisfirehose.DeliveryStreamGrants_FromDeliveryStream(deliveryStreamRef)
//
type DeliveryStreamGrants interface {
	Resource() interfacesawskinesisfirehose.IDeliveryStreamRef
	// Grant the `grantee` identity permissions to perform `actions`.
	PutRecords(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for DeliveryStreamGrants
type jsiiProxy_DeliveryStreamGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_DeliveryStreamGrants) Resource() interfacesawskinesisfirehose.IDeliveryStreamRef {
	var returns interfacesawskinesisfirehose.IDeliveryStreamRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for DeliveryStreamGrants.
func DeliveryStreamGrants_FromDeliveryStream(resource interfacesawskinesisfirehose.IDeliveryStreamRef) DeliveryStreamGrants {
	_init_.Initialize()

	if err := validateDeliveryStreamGrants_FromDeliveryStreamParameters(resource); err != nil {
		panic(err)
	}
	var returns DeliveryStreamGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.DeliveryStreamGrants",
		"fromDeliveryStream",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStreamGrants) PutRecords(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validatePutRecordsParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"putRecords",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

