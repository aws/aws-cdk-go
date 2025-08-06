package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for an AwsApi target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var policyStatement policyStatement
//
//   awsApiProps := &AwsApiProps{
//   	Action: jsii.String("action"),
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
//   	ApiVersion: jsii.String("apiVersion"),
//   	CatchErrorPattern: jsii.String("catchErrorPattern"),
//   	Parameters: parameters,
//   	PolicyStatement: policyStatement,
//   }
//
type AwsApiProps struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// API version to use for the service.
	// Deprecated: the handler code was migrated to AWS SDK for JavaScript v3, which does not support this feature anymore.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	// Default: - do not catch errors.
	//
	CatchErrorPattern *string `field:"optional" json:"catchErrorPattern" yaml:"catchErrorPattern"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Default: - no parameters.
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The IAM policy statement to allow the API call.
	//
	// Use only if
	// resource restriction is needed.
	// Default: - extract the permission from the API call.
	//
	PolicyStatement awsiam.PolicyStatement `field:"optional" json:"policyStatement" yaml:"policyStatement"`
}

