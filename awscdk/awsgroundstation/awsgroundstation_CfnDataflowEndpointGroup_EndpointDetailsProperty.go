package awsgroundstation


// The security details and endpoint information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointDetailsProperty := &endpointDetailsProperty{
//   	endpoint: &dataflowEndpointProperty{
//   		address: &socketAddressProperty{
//   			name: jsii.String("name"),
//   			port: jsii.Number(123),
//   		},
//   		mtu: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   	securityDetails: &securityDetailsProperty{
//   		roleArn: jsii.String("roleArn"),
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
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

