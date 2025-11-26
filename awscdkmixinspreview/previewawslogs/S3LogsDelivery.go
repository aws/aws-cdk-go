package previewawslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Delivers vended logs to an S3 Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucketRef IBucketRef
//
//   s3LogsDelivery := awscdkmixinspreview.Aws_logs.NewS3LogsDelivery(bucketRef, &S3LogsDeliveryProps{
//   	PermissionsVersion: awscdkmixinspreview.*Aws_logs.S3LogsDeliveryPermissionsVersion_V1,
//   })
//
// Experimental.
type S3LogsDelivery interface {
	ILogsDelivery
	// Binds the S3 destination to a delivery source and creates a delivery connection between them.
	// Experimental.
	Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig
}

// The jsii proxy struct for S3LogsDelivery
type jsiiProxy_S3LogsDelivery struct {
	jsiiProxy_ILogsDelivery
}

// Creates a new S3 Bucket delivery.
// Experimental.
func NewS3LogsDelivery(bucket interfacesawss3.IBucketRef, props *S3LogsDeliveryProps) S3LogsDelivery {
	_init_.Initialize()

	if err := validateNewS3LogsDeliveryParameters(bucket, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3LogsDelivery{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.S3LogsDelivery",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

// Creates a new S3 Bucket delivery.
// Experimental.
func NewS3LogsDelivery_Override(s S3LogsDelivery, bucket interfacesawss3.IBucketRef, props *S3LogsDeliveryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.S3LogsDelivery",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3LogsDelivery) Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig {
	if err := s.validateBindParameters(scope, deliverySource, sourceResourceArn); err != nil {
		panic(err)
	}
	var returns ILogsDeliveryConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, deliverySource, sourceResourceArn},
		&returns,
	)

	return returns
}

