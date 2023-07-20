package awswafv2


// Configures inspection of the response header. This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionheader.html
//
type CfnWebACL_ResponseInspectionHeaderProperty struct {
	// Values in the response header with the specified name that indicate a failed login attempt.
	//
	// To be counted as a failed login, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON example: `"FailureValues": [ "LoginFailed", "Failed login" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionheader.html#cfn-wafv2-webacl-responseinspectionheader-failurevalues
	//
	FailureValues *[]*string `field:"required" json:"failureValues" yaml:"failureValues"`
	// The name of the header to match against. The name must be an exact match, including case.
	//
	// JSON example: `"Name": [ "LoginResult" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionheader.html#cfn-wafv2-webacl-responseinspectionheader-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Values in the response header with the specified name that indicate a successful login attempt.
	//
	// To be counted as a successful login, the value must be an exact match, including case. Each value must be unique among the success and failure values.
	//
	// JSON example: `"SuccessValues": [ "LoginPassed", "Successful login" ]`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-responseinspectionheader.html#cfn-wafv2-webacl-responseinspectionheader-successvalues
	//
	SuccessValues *[]*string `field:"required" json:"successValues" yaml:"successValues"`
}

