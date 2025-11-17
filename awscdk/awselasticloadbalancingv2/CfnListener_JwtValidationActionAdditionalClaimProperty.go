package awselasticloadbalancingv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jwtValidationActionAdditionalClaimProperty := &JwtValidationActionAdditionalClaimProperty{
//   	Format: jsii.String("format"),
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-jwtvalidationactionadditionalclaim.html
//
type CfnListener_JwtValidationActionAdditionalClaimProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listener-jwtvalidationactionadditionalclaim-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listener-jwtvalidationactionadditionalclaim-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-jwtvalidationactionadditionalclaim.html#cfn-elasticloadbalancingv2-listener-jwtvalidationactionadditionalclaim-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

