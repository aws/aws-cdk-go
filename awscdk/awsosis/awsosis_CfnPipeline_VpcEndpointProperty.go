package awsosis


// An OpenSearch Ingestion-managed VPC endpoint that will access one or more pipelines.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcEndpointProperty := &VpcEndpointProperty{
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   	VpcId: jsii.String("vpcId"),
//   	VpcOptions: &VpcOptionsProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnPipeline_VpcEndpointProperty struct {
	// The unique identifier of the endpoint.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The ID for your VPC.
	//
	// AWS PrivateLink generates this value when you create a VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Information about the VPC, including associated subnets and security groups.
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

