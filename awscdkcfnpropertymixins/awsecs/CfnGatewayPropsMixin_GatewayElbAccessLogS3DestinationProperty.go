package awsecs


// The S3 destination for gateway ALB access logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gatewayElbAccessLogS3DestinationProperty := &GatewayElbAccessLogS3DestinationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbaccesslogs3destination.html
//
type CfnGatewayPropsMixin_GatewayElbAccessLogS3DestinationProperty struct {
	// The S3 bucket for access logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbaccesslogs3destination.html#cfn-ecs-gateway-gatewayelbaccesslogs3destination-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix for access logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbaccesslogs3destination.html#cfn-ecs-gateway-gatewayelbaccesslogs3destination-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

