package previewawscloudfrontmixins


// An origin in an origin group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   originGroupMemberProperty := &OriginGroupMemberProperty{
//   	OriginId: jsii.String("originId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroupmember.html
//
type CfnDistributionPropsMixin_OriginGroupMemberProperty struct {
	// The ID for an origin in an origin group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroupmember.html#cfn-cloudfront-distribution-origingroupmember-originid
	//
	OriginId *string `field:"optional" json:"originId" yaml:"originId"`
}

