package previewawss3outpostsmixins


// Properties for CfnEndpointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEndpointMixinProps := &CfnEndpointMixinProps{
//   	AccessType: jsii.String("accessType"),
//   	CustomerOwnedIpv4Pool: jsii.String("customerOwnedIpv4Pool"),
//   	FailedReason: &FailedReasonProperty{
//   		ErrorCode: jsii.String("errorCode"),
//   		Message: jsii.String("message"),
//   	},
//   	OutpostId: jsii.String("outpostId"),
//   	SecurityGroupId: jsii.String("securityGroupId"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-endpoint.html
//
type CfnEndpointMixinProps struct {
	// The container for the type of connectivity used to access the Amazon S3 on Outposts endpoint.
	//
	// To use the Amazon VPC , choose `Private` . To use the endpoint with an on-premises network, choose `CustomerOwnedIp` . If you choose `CustomerOwnedIp` , you must also provide the customer-owned IP address pool (CoIP pool).
	//
	// > `Private` is the default access type value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-endpoint.html#cfn-s3outposts-endpoint-accesstype
	//
	// Default: - "Private".
	//
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// The ID of the customer-owned IPv4 address pool (CoIP pool) for the endpoint.
	//
	// IP addresses are allocated from this pool for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-endpoint.html#cfn-s3outposts-endpoint-customerownedipv4pool
	//
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// The failure reason, if any, for a create or delete endpoint operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-endpoint.html#cfn-s3outposts-endpoint-failedreason
	//
	FailedReason interface{} `field:"optional" json:"failedReason" yaml:"failedReason"`
	// The ID of the Outpost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-endpoint.html#cfn-s3outposts-endpoint-outpostid
	//
	OutpostId *string `field:"optional" json:"outpostId" yaml:"outpostId"`
	// The ID of the security group used for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-endpoint.html#cfn-s3outposts-endpoint-securitygroupid
	//
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// The ID of the subnet used for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-endpoint.html#cfn-s3outposts-endpoint-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

