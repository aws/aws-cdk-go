package awswafv2


// Details for your use of the account creation fraud prevention managed rule group, `AWSManagedRulesACFPRuleSet` .
//
// This configuration is used in `ManagedRuleGroupConfig` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSManagedRulesACFPRuleSetProperty := &AWSManagedRulesACFPRuleSetProperty{
//   	CreationPath: jsii.String("creationPath"),
//   	RegistrationPagePath: jsii.String("registrationPagePath"),
//   	RequestInspection: &RequestInspectionACFPProperty{
//   		PayloadType: jsii.String("payloadType"),
//
//   		// the properties below are optional
//   		AddressFields: []interface{}{
//   			&FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   		},
//   		EmailField: &FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   		PasswordField: &FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   		PhoneNumberFields: []interface{}{
//   			&FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   		},
//   		UsernameField: &FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	EnableRegexInPath: jsii.Boolean(false),
//   	ResponseInspection: &ResponseInspectionProperty{
//   		BodyContains: &ResponseInspectionBodyContainsProperty{
//   			FailureStrings: []*string{
//   				jsii.String("failureStrings"),
//   			},
//   			SuccessStrings: []*string{
//   				jsii.String("successStrings"),
//   			},
//   		},
//   		Header: &ResponseInspectionHeaderProperty{
//   			FailureValues: []*string{
//   				jsii.String("failureValues"),
//   			},
//   			Name: jsii.String("name"),
//   			SuccessValues: []*string{
//   				jsii.String("successValues"),
//   			},
//   		},
//   		Json: &ResponseInspectionJsonProperty{
//   			FailureValues: []*string{
//   				jsii.String("failureValues"),
//   			},
//   			Identifier: jsii.String("identifier"),
//   			SuccessValues: []*string{
//   				jsii.String("successValues"),
//   			},
//   		},
//   		StatusCode: &ResponseInspectionStatusCodeProperty{
//   			FailureCodes: []interface{}{
//   				jsii.Number(123),
//   			},
//   			SuccessCodes: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html
//
type CfnWebACL_AWSManagedRulesACFPRuleSetProperty struct {
	// The path of the account creation endpoint for your application.
	//
	// This is the page on your website that accepts the completed registration form for a new user. This page must accept `POST` requests.
	//
	// For example, for the URL `https://example.com/web/signup` , you would provide the path `/web/signup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-creationpath
	//
	CreationPath *string `field:"required" json:"creationPath" yaml:"creationPath"`
	// The path of the account registration endpoint for your application.
	//
	// This is the page on your website that presents the registration form to new users.
	//
	// > This page must accept `GET` text/html requests.
	//
	// For example, for the URL `https://example.com/web/register` , you would provide the path `/web/register` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-registrationpagepath
	//
	RegistrationPagePath *string `field:"required" json:"registrationPagePath" yaml:"registrationPagePath"`
	// The criteria for inspecting account creation requests, used by the ACFP rule group to validate and track account creation attempts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-requestinspection
	//
	RequestInspection interface{} `field:"required" json:"requestInspection" yaml:"requestInspection"`
	// Allow the use of regular expressions in the registration page path and the account creation path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-enableregexinpath
	//
	EnableRegexInPath interface{} `field:"optional" json:"enableRegexInPath" yaml:"enableRegexInPath"`
	// The criteria for inspecting responses to account creation requests, used by the ACFP rule group to track account creation success rates.
	//
	// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
	//
	// The ACFP rule group evaluates the responses that your protected resources send back to client account creation attempts, keeping count of successful and failed attempts from each IP address and client session. Using this information, the rule group labels and mitigates requests from client sessions and IP addresses that have had too many successful account creation attempts in a short amount of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.html#cfn-wafv2-webacl-awsmanagedrulesacfpruleset-responseinspection
	//
	ResponseInspection interface{} `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

