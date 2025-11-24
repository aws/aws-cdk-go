package mixinsawsapigateway


// Properties for CfnRequestValidatorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRequestValidatorMixinProps := &CfnRequestValidatorMixinProps{
//   	Name: jsii.String("name"),
//   	RestApiId: jsii.String("restApiId"),
//   	ValidateRequestBody: jsii.Boolean(false),
//   	ValidateRequestParameters: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html
//
type CfnRequestValidatorMixinProps struct {
	// The name of this RequestValidator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-restapiid
	//
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// A Boolean flag to indicate whether to validate a request body according to the configured Model schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-validaterequestbody
	//
	ValidateRequestBody interface{} `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// A Boolean flag to indicate whether to validate request parameters ( `true` ) or not ( `false` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-validaterequestparameters
	//
	ValidateRequestParameters interface{} `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

