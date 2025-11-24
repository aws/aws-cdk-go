package mixinsawselasticache


// Properties for CfnSecurityGroupIngressPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityGroupIngressMixinProps := &CfnSecurityGroupIngressMixinProps{
//   	CacheSecurityGroupName: jsii.String("cacheSecurityGroupName"),
//   	Ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	Ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-securitygroupingress.html
//
type CfnSecurityGroupIngressMixinProps struct {
	// The name of the Cache Security Group to authorize.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-securitygroupingress.html#cfn-elasticache-securitygroupingress-cachesecuritygroupname
	//
	CacheSecurityGroupName *string `field:"optional" json:"cacheSecurityGroupName" yaml:"cacheSecurityGroupName"`
	// Name of the EC2 Security Group to include in the authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-securitygroupingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupname
	//
	Ec2SecurityGroupName *string `field:"optional" json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// Specifies the Amazon Account ID of the owner of the EC2 security group specified in the EC2SecurityGroupName property.
	//
	// The Amazon access key ID is not an acceptable value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-securitygroupingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupownerid
	//
	Ec2SecurityGroupOwnerId *string `field:"optional" json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

