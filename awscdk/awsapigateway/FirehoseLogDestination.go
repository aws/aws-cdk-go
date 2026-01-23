package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway"
)

// Use a Firehose delivery stream as a custom access log destination for API Gateway.
//
// Example:
//   destinationBucket := s3.NewBucket(this, jsii.String("Bucket"))
//   deliveryStreamRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("firehose.amazonaws.com")),
//   })
//
//   stream := firehose.NewCfnDeliveryStream(this, jsii.String("MyStream"), &CfnDeliveryStreamProps{
//   	DeliveryStreamName: jsii.String("amazon-apigateway-delivery-stream"),
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		BucketArn: destinationBucket.bucketArn,
//   		RoleArn: deliveryStreamRole.roleArn,
//   	},
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
//   	DeployOptions: &StageOptions{
//   		AccessLogDestination: apigateway.NewFirehoseLogDestination(stream),
//   		AccessLogFormat: apigateway.AccessLogFormat_JsonWithStandardFields(),
//   	},
//   })
//
type FirehoseLogDestination interface {
	IAccessLogDestination
	// Binds this destination to the Firehose delivery stream.
	Bind(stage interfacesawsapigateway.IStageRef) *AccessLogDestinationConfig
}

// The jsii proxy struct for FirehoseLogDestination
type jsiiProxy_FirehoseLogDestination struct {
	jsiiProxy_IAccessLogDestination
}

func NewFirehoseLogDestination(stream awskinesisfirehose.CfnDeliveryStream) FirehoseLogDestination {
	_init_.Initialize()

	if err := validateNewFirehoseLogDestinationParameters(stream); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehoseLogDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.FirehoseLogDestination",
		[]interface{}{stream},
		&j,
	)

	return &j
}

func NewFirehoseLogDestination_Override(f FirehoseLogDestination, stream awskinesisfirehose.CfnDeliveryStream) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.FirehoseLogDestination",
		[]interface{}{stream},
		f,
	)
}

func (f *jsiiProxy_FirehoseLogDestination) Bind(stage interfacesawsapigateway.IStageRef) *AccessLogDestinationConfig {
	if err := f.validateBindParameters(stage); err != nil {
		panic(err)
	}
	var returns *AccessLogDestinationConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{stage},
		&returns,
	)

	return returns
}

