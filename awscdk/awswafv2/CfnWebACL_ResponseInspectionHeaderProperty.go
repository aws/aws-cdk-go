package awswafv2


// Configures inspection of the response header. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet` .
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseInspectionHeaderProperty := &ResponseInspectionHeaderProperty{
//   	FailureValues: []*string{
//   		jsii.String("failureValues"),
//   	},
//   	Name: jsii.String("name"),
//   	SuccessValues: []*string{
//   		jsii.String("successValues"),
//   	},
//   }
//
type CfnWebACL_ResponseInspectionHeaderProperty struct {
	// Values in the response header with the specified name that indicate a failed login or account creation attempt.
	//
	// To be counted as a failure, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON examples: `"FailureValues": [ "LoginFailed", "Failed login" ]` and `"FailureValues": [ "AccountCreationFailed" ]`.
	FailureValues *[]*string `field:"required" json:"failureValues" yaml:"failureValues"`
	// The name of the header to match against. The name must be an exact match, including case.
	//
	// JSON example: `"Name": [ "RequestResult" ]`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Values in the response header with the specified name that indicate a successful login or account creation attempt.
	//
	// To be counted as a success, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON examples: `"SuccessValues": [ "LoginPassed", "Successful login" ]` and `"SuccessValues": [ "AccountCreated", "Successful account creation" ]`.
	SuccessValues *[]*string `field:"required" json:"successValues" yaml:"successValues"`
}

