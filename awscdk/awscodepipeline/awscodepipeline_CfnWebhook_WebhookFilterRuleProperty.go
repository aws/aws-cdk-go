package awscodepipeline


// The event criteria that specify when a webhook notification is sent to your URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookFilterRuleProperty := &webhookFilterRuleProperty{
//   	jsonPath: jsii.String("jsonPath"),
//
//   	// the properties below are optional
//   	matchEquals: jsii.String("matchEquals"),
//   }
//
type CfnWebhook_WebhookFilterRuleProperty struct {
	// A JsonPath expression that is applied to the body/payload of the webhook.
	//
	// The value selected by the JsonPath expression must match the value specified in the `MatchEquals` field. Otherwise, the request is ignored. For more information, see [Java JsonPath implementation](https://docs.aws.amazon.com/https://github.com/json-path/JsonPath) in GitHub.
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// The value selected by the `JsonPath` expression must match what is supplied in the `MatchEquals` field.
	//
	// Otherwise, the request is ignored. Properties from the target action configuration can be included as placeholders in this value by surrounding the action configuration key with curly brackets. For example, if the value supplied here is "refs/heads/{Branch}" and the target action has an action configuration property called "Branch" with a value of "main", the `MatchEquals` value is evaluated as "refs/heads/main". For a list of action configuration properties for built-in action types, see [Pipeline Structure Reference Action Requirements](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) .
	MatchEquals *string `field:"optional" json:"matchEquals" yaml:"matchEquals"`
}

