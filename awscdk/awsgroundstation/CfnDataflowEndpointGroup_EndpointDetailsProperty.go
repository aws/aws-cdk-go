package awsgroundstation


// The security details and endpoint information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointDetailsProperty := &EndpointDetailsProperty{
//   	Endpoint: &DataflowEndpointProperty{
//   		Address: &SocketAddressProperty{
//   			Name: jsii.String("name"),
//   			Port: jsii.Number(123),
//   		},
//   		Mtu: jsii.Number(123),
//   		Name: jsii.String("name"),
//   	},
//   	SecurityDetails: &SecurityDetailsProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDataflowEndpointGroup_EndpointDetailsProperty struct {
	// Information about the endpoint such as name and the endpoint address.
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The role ARN, and IDs for security groups and subnets.
	SecurityDetails interface{} `field:"optional" json:"securityDetails" yaml:"securityDetails"`
}

