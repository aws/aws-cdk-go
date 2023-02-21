package awsrds


// Properties for defining a `CfnDBProxyEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBProxyEndpointProps := &CfnDBProxyEndpointProps{
//   	DbProxyEndpointName: jsii.String("dbProxyEndpointName"),
//   	DbProxyName: jsii.String("dbProxyName"),
//   	VpcSubnetIds: []*string{
//   		jsii.String("vpcSubnetIds"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []tagFormatProperty{
//   		&tagFormatProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnDBProxyEndpointProps struct {
	// The name of the DB proxy endpoint to create.
	DbProxyEndpointName *string `field:"required" json:"dbProxyEndpointName" yaml:"dbProxyEndpointName"`
	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The VPC subnet IDs for the DB proxy endpoint that you create.
	//
	// You can specify a different set of subnet IDs than for the original DB proxy.
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
	Tags *[]*CfnDBProxyEndpoint_TagFormatProperty `field:"optional" json:"tags" yaml:"tags"`
	// The VPC security group IDs for the DB proxy endpoint that you create.
	//
	// You can specify a different set of security group IDs than for the original DB proxy. The default is the default security group for the VPC.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

