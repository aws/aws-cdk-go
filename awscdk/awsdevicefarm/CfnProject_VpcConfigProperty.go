package awsdevicefarm


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
//   	VpcId: jsii.String("vpcId"),
//   }
//
type CfnProject_VpcConfigProperty struct {
	// `CfnProject.VpcConfigProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `CfnProject.VpcConfigProperty.SubnetIds`.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// `CfnProject.VpcConfigProperty.VpcId`.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

