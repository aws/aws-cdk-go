package awswafv2


// Details for your use of the account takeover prevention managed rule group, `AWSManagedRulesATPRuleSet` .
//
// This configuration is used in `ManagedRuleGroupConfig` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSManagedRulesATPRuleSetProperty := &AWSManagedRulesATPRuleSetProperty{
//   	LoginPath: jsii.String("loginPath"),
//
//   	// the properties below are optional
//   	EnableRegexInPath: jsii.Boolean(false),
//   	RequestInspection: &RequestInspectionProperty{
//   		PasswordField: &FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   		PayloadType: jsii.String("payloadType"),
//   		UsernameField: &FieldIdentifierProperty{
//   			Identifier: jsii.String("identifier"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesatpruleset.html
//
type CfnWebACL_AWSManagedRulesATPRuleSetProperty struct {
	// The path of the login endpoint for your application.
	//
	// For example, for the URL `https://example.com/web/login` , you would provide the path `/web/login` . Login paths that start with the path that you provide are considered a match. For example `/web/login` matches the login paths `/web/login` , `/web/login/` , `/web/loginPage` , and `/web/login/thisPage` , but doesn't match the login path `/home/web/login` or `/website/login` .
	//
	// The rule group inspects only HTTP `POST` requests to your specified login endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesatpruleset.html#cfn-wafv2-webacl-awsmanagedrulesatpruleset-loginpath
	//
	LoginPath *string `field:"required" json:"loginPath" yaml:"loginPath"`
	// Allow the use of regular expressions in the login page path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesatpruleset.html#cfn-wafv2-webacl-awsmanagedrulesatpruleset-enableregexinpath
	//
	EnableRegexInPath interface{} `field:"optional" json:"enableRegexInPath" yaml:"enableRegexInPath"`
	// The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesatpruleset.html#cfn-wafv2-webacl-awsmanagedrulesatpruleset-requestinspection
	//
	RequestInspection interface{} `field:"optional" json:"requestInspection" yaml:"requestInspection"`
	// The criteria for inspecting responses to login requests, used by the ATP rule group to track login failure rates.
	//
	// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
	//
	// The ATP rule group evaluates the responses that your protected resources send back to client login attempts, keeping count of successful and failed attempts for each IP address and client session. Using this information, the rule group labels and mitigates requests from client sessions and IP addresses that have had too many failed login attempts in a short amount of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesatpruleset.html#cfn-wafv2-webacl-awsmanagedrulesatpruleset-responseinspection
	//
	ResponseInspection interface{} `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

