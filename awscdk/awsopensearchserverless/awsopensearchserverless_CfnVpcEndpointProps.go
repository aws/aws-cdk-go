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
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnVpcEndpointProps struct {
	// `AWS::OpenSearchServerless::VpcEndpoint.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::OpenSearchServerless::VpcEndpoint.VpcId`.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// `AWS::OpenSearchServerless::VpcEndpoint.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `AWS::OpenSearchServerless::VpcEndpoint.SubnetIds`.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

