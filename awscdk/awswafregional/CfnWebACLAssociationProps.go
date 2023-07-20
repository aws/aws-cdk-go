package awswafregional


// Properties for defining a `CfnWebACLAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebACLAssociationProps := &CfnWebACLAssociationProps{
//   	ResourceArn: jsii.String("resourceArn"),
//   	WebAclId: jsii.String("webAclId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html
//
type CfnWebACLAssociationProps struct {
	// The Amazon Resource Name (ARN) of the resource to protect with the web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// A unique identifier (ID) for the web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-webaclid
	//
	WebAclId *string `field:"required" json:"webAclId" yaml:"webAclId"`
}

