package awsappstream


// Describes VPC configuration information for fleets and image builders.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &VpcConfigProperty{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnAppBlockBuilder_VpcConfigProperty struct {
	// The identifiers of the security groups for the fleet or image builder.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The identifiers of the subnets to which a network interface is attached from the fleet instance or image builder instance.
	//
	// Fleet instances use one or more subnets. Image builder instances use one subnet.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

