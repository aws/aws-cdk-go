package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseSourceVPCConfigurationProperty := &DatabaseSourceVPCConfigurationProperty{
//   	VpcEndpointServiceName: jsii.String("vpcEndpointServiceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourcevpcconfiguration.html
//
type CfnDeliveryStream_DatabaseSourceVPCConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourcevpcconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourcevpcconfiguration-vpcendpointservicename
	//
	VpcEndpointServiceName *string `field:"required" json:"vpcEndpointServiceName" yaml:"vpcEndpointServiceName"`
}

