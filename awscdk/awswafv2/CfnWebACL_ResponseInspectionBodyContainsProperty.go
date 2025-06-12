package awswafv2


// Configures inspection of the response body.
//
// AWS WAF can inspect the first 65,536 bytes (64 KB) of the response body. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet` .
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseInspectionBodyContainsProperty := &ResponseInspectionBodyContainsProperty{
//   	FailureStrings: []*string{
//   		jsii.String("failureStrings"),
//   	},
//   	SuccessStrings: []*string{
//   		jsii.String("successStrings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionbodycontains.html
//
type CfnWebACL_ResponseInspectionBodyContainsProperty struct {
	// Strings in the body of the response that indicate a failed login or account creation attempt.
	//
	// To be counted as a failure, the string can be anywhere in the body and must be an exact match, including case. Each string must be unique among the success and failure strings.
	//
	// JSON example: `"FailureStrings": [ "Request failed" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionbodycontains.html#cfn-wafv2-webacl-responseinspectionbodycontains-failurestrings
	//
	FailureStrings *[]*string `field:"required" json:"failureStrings" yaml:"failureStrings"`
	// Strings in the body of the response that indicate a successful login or account creation attempt.
	//
	// To be counted as a success, the string can be anywhere in the body and must be an exact match, including case. Each string must be unique among the success and failure strings.
	//
	// JSON examples: `"SuccessStrings": [ "Login successful" ]` and `"SuccessStrings": [ "Account creation successful", "Welcome to our site!" ]`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionbodycontains.html#cfn-wafv2-webacl-responseinspectionbodycontains-successstrings
	//
	SuccessStrings *[]*string `field:"required" json:"successStrings" yaml:"successStrings"`
}

