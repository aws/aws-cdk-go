package previewawskinesisfirehosemixins


// The structure for details of the VPC Endpoint Service which Firehose uses to create a PrivateLink to the database.
//
// Amazon Data Firehose is in preview release and is subject to change.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   databaseSourceVPCConfigurationProperty := &DatabaseSourceVPCConfigurationProperty{
//   	VpcEndpointServiceName: jsii.String("vpcEndpointServiceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourcevpcconfiguration.html
//
type CfnDeliveryStreamPropsMixin_DatabaseSourceVPCConfigurationProperty struct {
	// The VPC endpoint service name which Firehose uses to create a PrivateLink to the database.
	//
	// The endpoint service must have the Firehose service principle `firehose.amazonaws.com` as an allowed principal on the VPC endpoint service. The VPC endpoint service name is a string that looks like `com.amazonaws.vpce.<region>.<vpc-endpoint-service-id>` .
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourcevpcconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourcevpcconfiguration-vpcendpointservicename
	//
	VpcEndpointServiceName *string `field:"optional" json:"vpcEndpointServiceName" yaml:"vpcEndpointServiceName"`
}

