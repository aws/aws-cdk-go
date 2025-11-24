package mixinsawssecurityhub


// Properties for CfnProductSubscriptionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProductSubscriptionMixinProps := &CfnProductSubscriptionMixinProps{
//   	ProductArn: jsii.String("productArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-productsubscription.html
//
type CfnProductSubscriptionMixinProps struct {
	// The ARN of the product to enable the integration for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-productsubscription.html#cfn-securityhub-productsubscription-productarn
	//
	ProductArn *string `field:"optional" json:"productArn" yaml:"productArn"`
}

