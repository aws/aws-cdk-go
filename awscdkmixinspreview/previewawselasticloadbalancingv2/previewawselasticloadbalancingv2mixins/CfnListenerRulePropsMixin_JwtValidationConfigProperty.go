package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jwtValidationConfigProperty := &JwtValidationConfigProperty{
//   	AdditionalClaims: []interface{}{
//   		&JwtValidationActionAdditionalClaimProperty{
//   			Format: jsii.String("format"),
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	Issuer: jsii.String("issuer"),
//   	JwksEndpoint: jsii.String("jwksEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationconfig.html
//
type CfnListenerRulePropsMixin_JwtValidationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationconfig.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationconfig-additionalclaims
	//
	AdditionalClaims interface{} `field:"optional" json:"additionalClaims" yaml:"additionalClaims"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationconfig.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationconfig-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationconfig.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationconfig-jwksendpoint
	//
	JwksEndpoint *string `field:"optional" json:"jwksEndpoint" yaml:"jwksEndpoint"`
}

