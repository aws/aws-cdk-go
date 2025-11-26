package previewawsapigatewayv2mixins


// Specifies whether the parameter is required.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parameterConstraintsProperty := &ParameterConstraintsProperty{
//   	Required: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routeresponse-parameterconstraints.html
//
type CfnRouteResponsePropsMixin_ParameterConstraintsProperty struct {
	// Specifies whether the parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routeresponse-parameterconstraints.html#cfn-apigatewayv2-routeresponse-parameterconstraints-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

