package awsrds


// Properties for defining a `CfnDBProxyEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBProxyEndpointProps := &cfnDBProxyEndpointProps{
//   	dbProxyEndpointName: jsii.String("dbProxyEndpointName"),
//   	dbProxyName: jsii.String("dbProxyName"),
//   	vpcSubnetIds: []*string{
//   		jsii.String("vpcSubnetIds"),
//   	},
//
//   	// the properties below are optional
//   	tags: []tagFormatProperty{
//   		&tagFormatProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetRole: jsii.String("targetRole"),
//   	vpcSecurityGroupIds: []*string{
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
	// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
	//
	// Valid Values: `READ_WRITE | READ_ONLY`.
	TargetRole *string `field:"optional" json:"targetRole" yaml:"targetRole"`
	// The VPC security group IDs for the DB proxy endpoint that you create.
	//
	// You can specify a different set of security group IDs than for the original DB proxy. The default is the default security group for the VPC.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

