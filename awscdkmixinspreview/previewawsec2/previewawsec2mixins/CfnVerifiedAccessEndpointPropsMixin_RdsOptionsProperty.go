package previewawsec2mixins


// Describes the RDS options for a Verified Access endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rdsOptionsProperty := &RdsOptionsProperty{
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	RdsDbClusterArn: jsii.String("rdsDbClusterArn"),
//   	RdsDbInstanceArn: jsii.String("rdsDbInstanceArn"),
//   	RdsDbProxyArn: jsii.String("rdsDbProxyArn"),
//   	RdsEndpoint: jsii.String("rdsEndpoint"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html
//
type CfnVerifiedAccessEndpointPropsMixin_RdsOptionsProperty struct {
	// The port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html#cfn-ec2-verifiedaccessendpoint-rdsoptions-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html#cfn-ec2-verifiedaccessendpoint-rdsoptions-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The ARN of the DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html#cfn-ec2-verifiedaccessendpoint-rdsoptions-rdsdbclusterarn
	//
	RdsDbClusterArn *string `field:"optional" json:"rdsDbClusterArn" yaml:"rdsDbClusterArn"`
	// The ARN of the RDS instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html#cfn-ec2-verifiedaccessendpoint-rdsoptions-rdsdbinstancearn
	//
	RdsDbInstanceArn *string `field:"optional" json:"rdsDbInstanceArn" yaml:"rdsDbInstanceArn"`
	// The ARN of the RDS proxy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html#cfn-ec2-verifiedaccessendpoint-rdsoptions-rdsdbproxyarn
	//
	RdsDbProxyArn *string `field:"optional" json:"rdsDbProxyArn" yaml:"rdsDbProxyArn"`
	// The RDS endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html#cfn-ec2-verifiedaccessendpoint-rdsoptions-rdsendpoint
	//
	RdsEndpoint *string `field:"optional" json:"rdsEndpoint" yaml:"rdsEndpoint"`
	// The IDs of the subnets.
	//
	// You can specify only one subnet per Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html#cfn-ec2-verifiedaccessendpoint-rdsoptions-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

