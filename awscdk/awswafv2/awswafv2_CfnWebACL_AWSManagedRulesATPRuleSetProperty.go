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
//   aWSManagedRulesATPRuleSetProperty := &aWSManagedRulesATPRuleSetProperty{
//   	loginPath: jsii.String("loginPath"),
//
//   	// the properties below are optional
//   	requestInspection: &requestInspectionProperty{
//   		passwordField: &fieldIdentifierProperty{
//   			identifier: jsii.String("identifier"),
//   		},
//   		payloadType: jsii.String("payloadType"),
//   		usernameField: &fieldIdentifierProperty{
//   			identifier: jsii.String("identifier"),
//   		},
//   	},
//   	responseInspection: &responseInspectionProperty{
//   		bodyContains: &responseInspectionBodyContainsProperty{
//   			failureStrings: []*string{
//   				jsii.String("failureStrings"),
//   			},
//   			successStrings: []*string{
//   				jsii.String("successStrings"),
//   			},
//   		},
//   		header: &responseInspectionHeaderProperty{
//   			failureValues: []*string{
//   				jsii.String("failureValues"),
//   			},
//   			name: jsii.String("name"),
//   			successValues: []*string{
//   				jsii.String("successValues"),
//   			},
//   		},
//   		json: &responseInspectionJsonProperty{
//   			failureValues: []*string{
//   				jsii.String("failureValues"),
//   			},
//   			identifier: jsii.String("identifier"),
//   			successValues: []*string{
//   				jsii.String("successValues"),
//   			},
//   		},
//   		statusCode: &responseInspectionStatusCodeProperty{
//   			failureCodes: []interface{}{
//   				jsii.Number(123),
//   			},
//   			successCodes: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_AWSManagedRulesATPRuleSetProperty struct {
	// The path of the login endpoint for your application.
	//
	// For example, for the URL `https://example.com/web/login` , you would provide the path `/web/login` .
	//
	// The rule group inspects only HTTP `POST` requests to your specified login endpoint.
	LoginPath *string `field:"required" json:"loginPath" yaml:"loginPath"`
	// The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage.
	RequestInspection interface{} `field:"optional" json:"requestInspection" yaml:"requestInspection"`
	// The criteria for inspecting responses to login requests, used by the ATP rule group to track login failure rates.
	//
	// The ATP rule group evaluates the responses that your protected resources send back to client login attempts, keeping count of successful and failed attempts from each IP address and client session. Using this information, the rule group labels and mitigates requests from client sessions and IP addresses that submit too many failed login attempts in a short amount of time.
	//
	// > Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.
	ResponseInspection interface{} `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

