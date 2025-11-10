package awsodb


// The configuration for a service network endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNetworkEndpointProperty := &ServiceNetworkEndpointProperty{
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   	VpcEndpointType: jsii.String("vpcEndpointType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-servicenetworkendpoint.html
//
type CfnOdbNetwork_ServiceNetworkEndpointProperty struct {
	// The identifier of the VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-servicenetworkendpoint.html#cfn-odb-odbnetwork-servicenetworkendpoint-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The type of the VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-servicenetworkendpoint.html#cfn-odb-odbnetwork-servicenetworkendpoint-vpcendpointtype
	//
	VpcEndpointType *string `field:"optional" json:"vpcEndpointType" yaml:"vpcEndpointType"`
}

