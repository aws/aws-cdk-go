package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var restApi restApi
//
//   requestValidatorProps := &requestValidatorProps{
//   	restApi: restApi,
//
//   	// the properties below are optional
//   	requestValidatorName: jsii.String("requestValidatorName"),
//   	validateRequestBody: jsii.Boolean(false),
//   	validateRequestParameters: jsii.Boolean(false),
//   }
//
type RequestValidatorProps struct {
	// The name of this request validator.
	RequestValidatorName *string `field:"optional" json:"requestValidatorName" yaml:"requestValidatorName"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody *bool `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	ValidateRequestParameters *bool `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
	// The rest API that this model is part of.
	//
	// The reason we need the RestApi object itself and not just the ID is because the model
	// is being tracked by the top-level RestApi object for the purpose of calculating it's
	// hash to determine the ID of the deployment. This allows us to automatically update
	// the deployment when the model of the REST API changes.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

