package awscodeconnections


// The VPC configuration provisioned for the host.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &VpcConfigurationProperty{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	TlsCertificate: jsii.String("tlsCertificate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeconnections-host-vpcconfiguration.html
//
type CfnHost_VpcConfigurationProperty struct {
	// The ID of the security group or security groups associated with the Amazon VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeconnections-host-vpcconfiguration.html#cfn-codeconnections-host-vpcconfiguration-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet or subnets associated with the Amazon VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeconnections-host-vpcconfiguration.html#cfn-codeconnections-host-vpcconfiguration-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon VPC connected to the infrastructure where your provider type is installed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeconnections-host-vpcconfiguration.html#cfn-codeconnections-host-vpcconfiguration-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The value of the Transport Layer Security (TLS) certificate associated with the infrastructure where your provider type is installed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeconnections-host-vpcconfiguration.html#cfn-codeconnections-host-vpcconfiguration-tlscertificate
	//
	TlsCertificate *string `field:"optional" json:"tlsCertificate" yaml:"tlsCertificate"`
}

