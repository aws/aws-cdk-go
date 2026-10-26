package awsmediaconvert


// Properties for defining a `CfnCertificateAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateAssociationProps := &CfnCertificateAssociationProps{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-certificateassociation.html
//
type CfnCertificateAssociationProps struct {
	// The ARN of the ACM certificate to associate with MediaConvert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-certificateassociation.html#cfn-mediaconvert-certificateassociation-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

