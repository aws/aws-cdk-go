package awswafv2


// The criteria for inspecting responses to login requests and account creation requests, used by the ATP and ACFP rule groups to track login and account creation success and failure rates.
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// The rule groups evaluates the responses that your protected resources send back to client login and account creation attempts, keeping count of successful and failed attempts from each IP address and client session. Using this information, the rule group labels and mitigates requests from client sessions and IP addresses with too much suspicious activity in a short amount of time.
//
// This is part of the `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet` configurations in `ManagedRuleGroupConfig` .
//
// Enable response inspection by configuring exactly one component of the response to inspect, for example, `Header` or `StatusCode` . You can't configure more than one component for inspection. If you don't configure any of the response inspection options, response inspection is disabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseInspectionProperty := &ResponseInspectionProperty{
//   	BodyContains: &ResponseInspectionBodyContainsProperty{
//   		FailureStrings: []*string{
//   			jsii.String("failureStrings"),
//   		},
//   		SuccessStrings: []*string{
//   			jsii.String("successStrings"),
//   		},
//   	},
//   	Header: &ResponseInspectionHeaderProperty{
//   		FailureValues: []*string{
//   			jsii.String("failureValues"),
//   		},
//   		Name: jsii.String("name"),
//   		SuccessValues: []*string{
//   			jsii.String("successValues"),
//   		},
//   	},
//   	Json: &ResponseInspectionJsonProperty{
//   		FailureValues: []*string{
//   			jsii.String("failureValues"),
//   		},
//   		Identifier: jsii.String("identifier"),
//   		SuccessValues: []*string{
//   			jsii.String("successValues"),
//   		},
//   	},
//   	StatusCode: &ResponseInspectionStatusCodeProperty{
//   		FailureCodes: []interface{}{
//   			jsii.Number(123),
//   		},
//   		SuccessCodes: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspection.html
//
type CfnWebACL_ResponseInspectionProperty struct {
	// Configures inspection of the response body for success and failure indicators.
	//
	// AWS WAF can inspect the first 65,536 bytes (64 KB) of the response body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspection.html#cfn-wafv2-webacl-responseinspection-bodycontains
	//
	BodyContains interface{} `field:"optional" json:"bodyContains" yaml:"bodyContains"`
	// Configures inspection of the response header for success and failure indicators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspection.html#cfn-wafv2-webacl-responseinspection-header
	//
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Configures inspection of the response JSON for success and failure indicators.
	//
	// AWS WAF can inspect the first 65,536 bytes (64 KB) of the response JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspection.html#cfn-wafv2-webacl-responseinspection-json
	//
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Configures inspection of the response status code for success and failure indicators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspection.html#cfn-wafv2-webacl-responseinspection-statuscode
	//
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

