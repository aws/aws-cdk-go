package awsopensearchserverless


// Properties for defining a `CfnVpcEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcEndpointProps := &cfnVpcEndpointProps{
//   	name: jsii.String("name"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
type CfnVpcEndpointProps struct {
	// The name of the endpoint.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the subnets from which you access OpenSearch Serverless.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC from which you access OpenSearch Serverless.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The unique identifiers of the security groups that define the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

