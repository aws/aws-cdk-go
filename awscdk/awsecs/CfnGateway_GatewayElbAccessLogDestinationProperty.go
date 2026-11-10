package awsecs


// The destination for gateway ALB access logs.
//
// Exactly one destination type must be specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayElbAccessLogDestinationProperty := &GatewayElbAccessLogDestinationProperty{
//   	S3: &GatewayElbAccessLogS3DestinationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbaccesslogdestination.html
//
type CfnGateway_GatewayElbAccessLogDestinationProperty struct {
	// The S3 destination for gateway ALB access logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbaccesslogdestination.html#cfn-ecs-gateway-gatewayelbaccesslogdestination-s3
	//
	S3 interface{} `field:"required" json:"s3" yaml:"s3"`
}

