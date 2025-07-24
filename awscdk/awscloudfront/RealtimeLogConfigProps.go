package awscloudfront


// Properties for defining a RealtimeLogConfig resource.
//
// Example:
//   // Adding realtime logs config to a Cloudfront Distribution on default behavior.
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//   var stream stream
//
//
//   realTimeConfig := cloudfront.NewRealtimeLogConfig(this, jsii.String("realtimeLog"), &RealtimeLogConfigProps{
//   	EndPoints: []endpoint{
//   		cloudfront.*endpoint_FromKinesisStream(stream),
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
type RealtimeLogConfigProps struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration.
	EndPoints *[]Endpoint `field:"required" json:"endPoints" yaml:"endPoints"`
	// A list of fields that are included in each real-time log record.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields
	//
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
	// The sampling rate for this real-time log configuration.
	SamplingRate *float64 `field:"required" json:"samplingRate" yaml:"samplingRate"`
	// The unique name of this real-time log configuration.
	// Default: - the unique construct ID.
	//
	RealtimeLogConfigName *string `field:"optional" json:"realtimeLogConfigName" yaml:"realtimeLogConfigName"`
}

