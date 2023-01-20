package awsapigateway


// Properties for defining a `CfnRequestValidator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRequestValidatorProps := &cfnRequestValidatorProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	validateRequestBody: jsii.Boolean(false),
//   	validateRequestParameters: jsii.Boolean(false),
//   }
//
type CfnRequestValidatorProps struct {
	// The identifier of the targeted API entity.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The name of this request validator.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody interface{} `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	ValidateRequestParameters interface{} `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

