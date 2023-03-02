package awsappstream


// The VPC configuration for the image builder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnImageBuilder_VpcConfigProperty struct {
	// The identifiers of the security groups for the image builder.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The identifier of the subnet to which a network interface is attached from the image builder instance.
	//
	// An image builder instance can use one subnet.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

