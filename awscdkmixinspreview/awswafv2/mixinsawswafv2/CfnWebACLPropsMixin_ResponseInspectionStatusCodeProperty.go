package mixinsawswafv2


// Configures inspection of the response status code. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet` .
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseInspectionStatusCodeProperty := &ResponseInspectionStatusCodeProperty{
//   	FailureCodes: []interface{}{
//   		jsii.Number(123),
//   	},
//   	SuccessCodes: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionstatuscode.html
//
type CfnWebACLPropsMixin_ResponseInspectionStatusCodeProperty struct {
	// Status codes in the response that indicate a failed login or account creation attempt.
	//
	// To be counted as a failure, the response status code must match one of these. Each code must be unique among the success and failure status codes.
	//
	// JSON example: `"FailureCodes": [ 400, 404 ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionstatuscode.html#cfn-wafv2-webacl-responseinspectionstatuscode-failurecodes
	//
	FailureCodes interface{} `field:"optional" json:"failureCodes" yaml:"failureCodes"`
	// Status codes in the response that indicate a successful login or account creation attempt.
	//
	// To be counted as a success, the response status code must match one of these. Each code must be unique among the success and failure status codes.
	//
	// JSON example: `"SuccessCodes": [ 200, 201 ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionstatuscode.html#cfn-wafv2-webacl-responseinspectionstatuscode-successcodes
	//
	SuccessCodes interface{} `field:"optional" json:"successCodes" yaml:"successCodes"`
}

