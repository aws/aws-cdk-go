package awswafv2


// Configures inspection of the response status code. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` .
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnWebACL_ResponseInspectionStatusCodeProperty struct {
	// Status codes in the response that indicate a failed login attempt.
	//
	// To be counted as a failed login, the response status code must match one of these. Each code must be unique among the success and failure status codes.
	//
	// JSON example: `"FailureCodes": [ 400, 404 ]`.
	FailureCodes interface{} `field:"required" json:"failureCodes" yaml:"failureCodes"`
	// Status codes in the response that indicate a successful login attempt.
	//
	// To be counted as a successful login, the response status code must match one of these. Each code must be unique among the success and failure status codes.
	//
	// JSON example: `"SuccessCodes": [ 200, 201 ]`.
	SuccessCodes interface{} `field:"required" json:"successCodes" yaml:"successCodes"`
}

