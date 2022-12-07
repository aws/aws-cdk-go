package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jWTConfigurationProperty := &jWTConfigurationProperty{
//   	audience: []*string{
//   		jsii.String("audience"),
//   	},
//   	issuer: jsii.String("issuer"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-authorizer-jwtconfiguration.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnAuthorizerV2_JWTConfigurationProperty struct {
	// `CfnAuthorizerV2.JWTConfigurationProperty.Audience`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-authorizer-jwtconfiguration.html#cfn-apigatewayv2-authorizer-jwtconfiguration-audience
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Audience *[]*string `field:"optional" json:"audience" yaml:"audience"`
	// `CfnAuthorizerV2.JWTConfigurationProperty.Issuer`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-authorizer-jwtconfiguration.html#cfn-apigatewayv2-authorizer-jwtconfiguration-issuer
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

