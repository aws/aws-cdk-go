// CDK Destinations Constructs for AWS Kinesis Firehose
package awscdkkinesisfirehosedestinationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha/v2"
	"github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An S3 bucket destination for data from a Kinesis Data Firehose delivery stream.
//
// Example:
//   // Specify the roles created above when defining the destination and delivery stream.
//   var bucket bucket
//   // Create service roles for the delivery stream and destination.
//   // These can be used for other purposes and granted access to different resources.
//   // They must include the Kinesis Data Firehose service principal in their trust policies.
//   // Two separate roles are shown below, but the same role can be used for both purposes.
//   deliveryStreamRole := iam.NewRole(this, jsii.String("Delivery Stream Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
//   })
//   destinationRole := iam.NewRole(this, jsii.String("Destination Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
//   })
//   destination := destinations.NewS3Bucket(bucket, &s3BucketProps{
//   	role: destinationRole,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destination,
//   	},
//   	role: deliveryStreamRole,
//   })
//
// Experimental.
type S3Bucket interface {
	awscdkkinesisfirehosealpha.IDestination
	// Binds this destination to the Kinesis Data Firehose delivery stream.
	//
	// Implementers should use this method to bind resources to the stack and initialize values using the provided stream.
	// Experimental.
	Bind(scope constructs.Construct, _options *awscdkkinesisfirehosealpha.DestinationBindOptions) *awscdkkinesisfirehosealpha.DestinationConfig
}

// The jsii proxy struct for S3Bucket
type jsiiProxy_S3Bucket struct {
	internal.Type__awscdkkinesisfirehosealphaIDestination
}

// Experimental.
func NewS3Bucket(bucket awss3.IBucket, props *S3BucketProps) S3Bucket {
	_init_.Initialize()

	if err := validateNewS3BucketParameters(bucket, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Bucket{}

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.S3Bucket",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Bucket_Override(s S3Bucket, bucket awss3.IBucket, props *S3BucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.S3Bucket",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3Bucket) Bind(scope constructs.Construct, _options *awscdkkinesisfirehosealpha.DestinationBindOptions) *awscdkkinesisfirehosealpha.DestinationConfig {
	if err := s.validateBindParameters(scope, _options); err != nil {
		panic(err)
	}
	var returns *awscdkkinesisfirehosealpha.DestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, _options},
		&returns,
	)

	return returns
}

