package mixinsawselasticloadbalancingv2


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationactionadditionalclaim-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

