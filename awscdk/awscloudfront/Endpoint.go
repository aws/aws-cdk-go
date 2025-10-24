package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

// Represents the endpoints available for targetting within a realtime log config resource.
//
// Example:
//   // Adding realtime logs config to a Cloudfront Distribution on default behavior.
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//   var stream Stream
//
//
//   realTimeConfig := cloudfront.NewRealtimeLogConfig(this, jsii.String("realtimeLog"), &RealtimeLogConfigProps{
//   	EndPoints: []Endpoint{
//   		cloudfront.Endpoint_FromKinesisStream(stream),
//   	},
//   	Fields: []*string{
//   		jsii.String("timestamp"),
//   		jsii.String("c-ip"),
//   		jsii.String("time-to-first-byte"),
//   		jsii.String("sc-status"),
//   	},
//   	RealtimeLogConfigName: jsii.String("my-delivery-stream"),
//   	SamplingRate: jsii.Number(100),
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("myCdn"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   		RealtimeLogConfig: realTimeConfig,
//   	},
//   })
//
type Endpoint interface {
}

// The jsii proxy struct for Endpoint
type jsiiProxy_Endpoint struct {
	_ byte // padding
}

// Configure a Kinesis Stream Endpoint for Realtime Log Config.
// Default: - a role will be created and used across your endpoints.
//
func Endpoint_FromKinesisStream(stream awskinesis.IStream, role awsiam.IRole) Endpoint {
	_init_.Initialize()

	if err := validateEndpoint_FromKinesisStreamParameters(stream); err != nil {
		panic(err)
	}
	var returns Endpoint

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.Endpoint",
		"fromKinesisStream",
		[]interface{}{stream, role},
		&returns,
	)

	return returns
}

