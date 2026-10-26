package awsmediaconvert


// Properties for CfnCertificateAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCertificateAssociationMixinProps := &CfnCertificateAssociationMixinProps{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-certificateassociation.html
//
type CfnCertificateAssociationMixinProps struct {
	// The ARN of the ACM certificate to associate with MediaConvert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-certificateassociation.html#cfn-mediaconvert-certificateassociation-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

