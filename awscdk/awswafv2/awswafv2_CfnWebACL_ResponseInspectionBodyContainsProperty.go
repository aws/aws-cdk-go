package awswafv2


// Configures inspection of the response body.
//
// AWS WAF can inspect the first 65,536 bytes (64 KB) of the response body. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` .
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseInspectionBodyContainsProperty := &responseInspectionBodyContainsProperty{
//   	failureStrings: []*string{
//   		jsii.String("failureStrings"),
//   	},
//   	successStrings: []*string{
//   		jsii.String("successStrings"),
//   	},
//   }
//
type CfnWebACL_ResponseInspectionBodyContainsProperty struct {
	// Strings in the body of the response that indicate a failed login attempt.
	//
	// To be counted as a failed login, the string can be anywhere in the body and must be an exact match, including case. Each string must be unique among the success and failure strings.
	//
	// JSON example: `"FailureStrings": [ "Login failed" ]`.
	FailureStrings *[]*string `field:"required" json:"failureStrings" yaml:"failureStrings"`
	// Strings in the body of the response that indicate a successful login attempt.
	//
	// To be counted as a successful login, the string can be anywhere in the body and must be an exact match, including case. Each string must be unique among the success and failure strings.
	//
	// JSON example: `"SuccessStrings": [ "Login successful", "Welcome to our site!" ]`
	SuccessStrings *[]*string `field:"required" json:"successStrings" yaml:"successStrings"`
}

