package awsapigateway


// Properties for defining a `CfnRequestValidator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRequestValidatorProps := &CfnRequestValidatorProps{
//   	RestApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	ValidateRequestBody: jsii.Boolean(false),
//   	ValidateRequestParameters: jsii.Boolean(false),
//   }
//
type CfnRequestValidatorProps struct {
	// The string identifier of the associated RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The name of this RequestValidator.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A Boolean flag to indicate whether to validate a request body according to the configured Model schema.
	ValidateRequestBody interface{} `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// A Boolean flag to indicate whether to validate request parameters ( `true` ) or not ( `false` ).
	ValidateRequestParameters interface{} `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

