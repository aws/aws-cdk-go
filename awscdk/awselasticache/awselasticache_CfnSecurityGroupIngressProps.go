package awselasticache


// Properties for defining a `CfnSecurityGroupIngress`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityGroupIngressProps := &cfnSecurityGroupIngressProps{
//   	cacheSecurityGroupName: jsii.String("cacheSecurityGroupName"),
//   	ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//
//   	// the properties below are optional
//   	ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }
//
type CfnSecurityGroupIngressProps struct {
	// The name of the Cache Security Group to authorize.
	CacheSecurityGroupName *string `field:"required" json:"cacheSecurityGroupName" yaml:"cacheSecurityGroupName"`
	// Name of the EC2 Security Group to include in the authorization.
	Ec2SecurityGroupName *string `field:"required" json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// Specifies the Amazon Account ID of the owner of the EC2 security group specified in the EC2SecurityGroupName property.
	//
	// The Amazon access key ID is not an acceptable value.
	Ec2SecurityGroupOwnerId *string `field:"optional" json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

