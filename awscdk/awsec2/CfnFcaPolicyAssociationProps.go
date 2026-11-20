package awsec2


// Properties for defining a `CfnFcaPolicyAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFcaPolicyAssociationProps := &CfnFcaPolicyAssociationProps{
//   	FcaPolicyId: jsii.String("fcaPolicyId"),
//   	NetworkResourceId: jsii.String("networkResourceId"),
//
//   	// the properties below are optional
//   	NetworkResourceScopeType: jsii.String("networkResourceScopeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicyassociation.html
//
type CfnFcaPolicyAssociationProps struct {
	// The ID of the FCA policy to associate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicyassociation.html#cfn-ec2-fcapolicyassociation-fcapolicyid
	//
	FcaPolicyId *string `field:"required" json:"fcaPolicyId" yaml:"fcaPolicyId"`
	// The ID of the network resource to associate the FCA policy with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicyassociation.html#cfn-ec2-fcapolicyassociation-networkresourceid
	//
	NetworkResourceId *string `field:"required" json:"networkResourceId" yaml:"networkResourceId"`
	// Associate by scope instead of directly.
	//
	// When set, NetworkResourceId is the ID of the scope (an account ID, a VPC ID, or a subnet ID) and the association covers every eligible network resource in that scope; when omitted, NetworkResourceId is a direct network resource (for example a NAT gateway or instance) and must match the policy's NetworkResourceType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicyassociation.html#cfn-ec2-fcapolicyassociation-networkresourcescopetype
	//
	NetworkResourceScopeType *string `field:"optional" json:"networkResourceScopeType" yaml:"networkResourceScopeType"`
}

