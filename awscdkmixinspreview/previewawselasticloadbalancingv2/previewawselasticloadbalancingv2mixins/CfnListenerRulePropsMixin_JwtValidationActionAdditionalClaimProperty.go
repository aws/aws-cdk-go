package previewawselasticloadbalancingv2mixins


// Information about an additional claim to validate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jwtValidationActionAdditionalClaimProperty := &JwtValidationActionAdditionalClaimProperty{
//   	Format: jsii.String("format"),
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim.html
//
type CfnListenerRulePropsMixin_JwtValidationActionAdditionalClaimProperty struct {
	// The format of the claim value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The name of the claim.
	//
	// You can't specify `exp` , `iss` , `nbf` , or `iat` because we validate them by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The claim value.
	//
	// The maximum size of the list is 10. Each value can be up to 256 characters in length. If the format is `space-separated-values` , the values can't include spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

