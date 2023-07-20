package awswafv2


// Configures inspection of the response JSON.
//
// AWS WAF can inspect the first 65,536 bytes (64 KB) of the response JSON. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` .
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseInspectionJsonProperty := &ResponseInspectionJsonProperty{
//   	FailureValues: []*string{
//   		jsii.String("failureValues"),
//   	},
//   	Identifier: jsii.String("identifier"),
//   	SuccessValues: []*string{
//   		jsii.String("successValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionjson.html
//
type CfnWebACL_ResponseInspectionJsonProperty struct {
	// Values for the specified identifier in the response JSON that indicate a failed login attempt.
	//
	// To be counted as a failed login, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON example: `"FailureValues": [ "False", "Failed" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionjson.html#cfn-wafv2-webacl-responseinspectionjson-failurevalues
	//
	FailureValues *[]*string `field:"required" json:"failureValues" yaml:"failureValues"`
	// The identifier for the value to match against in the JSON.
	//
	// The identifier must be an exact match, including case.
	//
	// JSON example: `"Identifier": [ "/login/success" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionjson.html#cfn-wafv2-webacl-responseinspectionjson-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Values for the specified identifier in the response JSON that indicate a successful login attempt.
	//
	// To be counted as a successful login, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON example: `"SuccessValues": [ "True", "Succeeded" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionjson.html#cfn-wafv2-webacl-responseinspectionjson-successvalues
	//
	SuccessValues *[]*string `field:"required" json:"successValues" yaml:"successValues"`
}

