package awselasticloadbalancing


// Specifies where and how access logs are stored for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLoggingPolicyProperty := &AccessLoggingPolicyProperty{
//   	Enabled: jsii.Boolean(false),
//   	S3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	EmitInterval: jsii.Number(123),
//   	S3BucketPrefix: jsii.String("s3BucketPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-accessloggingpolicy.html
//
type CfnLoadBalancer_AccessLoggingPolicyProperty struct {
	// Specifies whether access logs are enabled for the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-accessloggingpolicy.html#cfn-elasticloadbalancing-loadbalancer-accessloggingpolicy-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the Amazon S3 bucket where the access logs are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-accessloggingpolicy.html#cfn-elasticloadbalancing-loadbalancer-accessloggingpolicy-s3bucketname
	//
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The interval for publishing the access logs. You can specify an interval of either 5 minutes or 60 minutes.
	//
	// Default: 60 minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-accessloggingpolicy.html#cfn-elasticloadbalancing-loadbalancer-accessloggingpolicy-emitinterval
	//
	EmitInterval *float64 `field:"optional" json:"emitInterval" yaml:"emitInterval"`
	// The logical hierarchy you created for your Amazon S3 bucket, for example `my-bucket-prefix/prod` .
	//
	// If the prefix is not provided, the log is placed at the root level of the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-accessloggingpolicy.html#cfn-elasticloadbalancing-loadbalancer-accessloggingpolicy-s3bucketprefix
	//
	S3BucketPrefix *string `field:"optional" json:"s3BucketPrefix" yaml:"s3BucketPrefix"`
}

