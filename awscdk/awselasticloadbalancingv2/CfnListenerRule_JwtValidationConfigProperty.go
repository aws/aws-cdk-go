package awselasticloadbalancingv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jwtValidationConfigProperty := &JwtValidationConfigProperty{
//   	Issuer: jsii.String("issuer"),
//   	JwksEndpoint: jsii.String("jwksEndpoint"),
//
//   	// the properties below are optional
//   	AdditionalClaims: []interface{}{
//   		&JwtValidationActionAdditionalClaimProperty{
//   			Format: jsii.String("format"),
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationconfig.html
//
type CfnListenerRule_JwtValidationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationconfig.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationconfig-issuer
	//
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationconfig.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationconfig-jwksendpoint
	//
	JwksEndpoint *string `field:"required" json:"jwksEndpoint" yaml:"jwksEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-jwtvalidationconfig.html#cfn-elasticloadbalancingv2-listenerrule-jwtvalidationconfig-additionalclaims
	//
	AdditionalClaims interface{} `field:"optional" json:"additionalClaims" yaml:"additionalClaims"`
}

