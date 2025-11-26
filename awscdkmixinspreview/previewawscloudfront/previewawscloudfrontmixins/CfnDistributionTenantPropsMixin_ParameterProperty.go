package previewawscloudfrontmixins


// A list of parameter values to add to the resource.
//
// A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parameterProperty := &ParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-parameter.html
//
type CfnDistributionTenantPropsMixin_ParameterProperty struct {
	// The parameter name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-parameter.html#cfn-cloudfront-distributiontenant-parameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-parameter.html#cfn-cloudfront-distributiontenant-parameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

