package awswafv2


// The criteria for inspecting responses to login requests, used by the ATP rule group to track login failure rates.
//
// The ATP rule group evaluates the responses that your protected resources send back to client login attempts, keeping count of successful and failed attempts from each IP address and client session. Using this information, the rule group labels and mitigates requests from client sessions and IP addresses that submit too many failed login attempts in a short amount of time.
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// This is part of the `AWSManagedRulesATPRuleSet` configuration in `ManagedRuleGroupConfig` .
//
// Enable login response inspection by configuring exactly one component of the response to inspect. You can't configure more than one. If you don't configure any of the response inspection options, response inspection is disabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseInspectionProperty := &responseInspectionProperty{
//   	bodyContains: &responseInspectionBodyContainsProperty{
//   		failureStrings: []*string{
//   			jsii.String("failureStrings"),
//   		},
//   		successStrings: []*string{
//   			jsii.String("successStrings"),
//   		},
//   	},
//   	header: &responseInspectionHeaderProperty{
//   		failureValues: []*string{
//   			jsii.String("failureValues"),
//   		},
//   		name: jsii.String("name"),
//   		successValues: []*string{
//   			jsii.String("successValues"),
//   		},
//   	},
//   	json: &responseInspectionJsonProperty{
//   		failureValues: []*string{
//   			jsii.String("failureValues"),
//   		},
//   		identifier: jsii.String("identifier"),
//   		successValues: []*string{
//   			jsii.String("successValues"),
//   		},
//   	},
//   	statusCode: &responseInspectionStatusCodeProperty{
//   		failureCodes: []interface{}{
//   			jsii.Number(123),
//   		},
//   		successCodes: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnWebACL_ResponseInspectionProperty struct {
	// Configures inspection of the response body.
	//
	// AWS WAF can inspect the first 65,536 bytes (64 KB) of the response body.
	BodyContains interface{} `field:"optional" json:"bodyContains" yaml:"bodyContains"`
	// Configures inspection of the response header.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Configures inspection of the response JSON.
	//
	// AWS WAF can inspect the first 65,536 bytes (64 KB) of the response JSON.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Configures inspection of the response status code.
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

