package awselasticloadbalancing


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sourceSecurityGroupProperty := &SourceSecurityGroupProperty{
//   	GroupName: jsii.String("groupName"),
//   	OwnerAlias: jsii.String("ownerAlias"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-sourcesecuritygroup.html
//
type CfnLoadBalancerPropsMixin_SourceSecurityGroupProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-sourcesecuritygroup.html#cfn-elasticloadbalancing-loadbalancer-sourcesecuritygroup-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-sourcesecuritygroup.html#cfn-elasticloadbalancing-loadbalancer-sourcesecuritygroup-owneralias
	//
	OwnerAlias *string `field:"optional" json:"ownerAlias" yaml:"ownerAlias"`
}

