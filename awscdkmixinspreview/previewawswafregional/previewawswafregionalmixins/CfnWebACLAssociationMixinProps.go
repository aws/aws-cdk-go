package previewawswafregionalmixins


// Properties for CfnWebACLAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLAssociationMixinProps := &CfnWebACLAssociationMixinProps{
//   	ResourceArn: jsii.String("resourceArn"),
//   	WebAclId: jsii.String("webAclId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html
//
type CfnWebACLAssociationMixinProps struct {
	// The Amazon Resource Name (ARN) of the resource to protect with the web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// A unique identifier (ID) for the web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-webaclid
	//
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

