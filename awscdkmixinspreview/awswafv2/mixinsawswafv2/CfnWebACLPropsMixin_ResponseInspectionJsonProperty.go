package mixinsawswafv2


// Configures inspection of the response JSON.
//
// AWS WAF can inspect the first 65,536 bytes (64 KB) of the response JSON. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet` .
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnWebACLPropsMixin_ResponseInspectionJsonProperty struct {
	// Values for the specified identifier in the response JSON that indicate a failed login or account creation attempt.
	//
	// To be counted as a failure, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON example: `"FailureValues": [ "False", "Failed" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionjson.html#cfn-wafv2-webacl-responseinspectionjson-failurevalues
	//
	FailureValues *[]*string `field:"optional" json:"failureValues" yaml:"failureValues"`
	// The identifier for the value to match against in the JSON.
	//
	// The identifier must be an exact match, including case.
	//
	// JSON examples: `"Identifier": [ "/login/success" ]` and `"Identifier": [ "/sign-up/success" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionjson.html#cfn-wafv2-webacl-responseinspectionjson-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Values for the specified identifier in the response JSON that indicate a successful login or account creation attempt.
	//
	// To be counted as a success, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON example: `"SuccessValues": [ "True", "Succeeded" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionjson.html#cfn-wafv2-webacl-responseinspectionjson-successvalues
	//
	SuccessValues *[]*string `field:"optional" json:"successValues" yaml:"successValues"`
}

