package mixinsawswafv2


// Configures inspection of the response header. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet` .
//
// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionheader.html
//
type CfnWebACLPropsMixin_ResponseInspectionHeaderProperty struct {
	// Values in the response header with the specified name that indicate a failed login or account creation attempt.
	//
	// To be counted as a failure, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON examples: `"FailureValues": [ "LoginFailed", "Failed login" ]` and `"FailureValues": [ "AccountCreationFailed" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionheader.html#cfn-wafv2-webacl-responseinspectionheader-failurevalues
	//
	FailureValues *[]*string `field:"optional" json:"failureValues" yaml:"failureValues"`
	// The name of the header to match against. The name must be an exact match, including case.
	//
	// JSON example: `"Name": [ "RequestResult" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionheader.html#cfn-wafv2-webacl-responseinspectionheader-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Values in the response header with the specified name that indicate a successful login or account creation attempt.
	//
	// To be counted as a success, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON examples: `"SuccessValues": [ "LoginPassed", "Successful login" ]` and `"SuccessValues": [ "AccountCreated", "Successful account creation" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionheader.html#cfn-wafv2-webacl-responseinspectionheader-successvalues
	//
	SuccessValues *[]*string `field:"optional" json:"successValues" yaml:"successValues"`
}

