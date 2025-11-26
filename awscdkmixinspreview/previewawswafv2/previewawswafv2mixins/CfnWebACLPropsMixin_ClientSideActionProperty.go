package previewawswafv2mixins


// This is part of the `AWSManagedRulesAntiDDoSRuleSet` `ClientSideActionConfig` configuration in `ManagedRuleGroupConfig` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clientSideActionProperty := &ClientSideActionProperty{
//   	ExemptUriRegularExpressions: []interface{}{
//   		&RegexProperty{
//   			RegexString: jsii.String("regexString"),
//   		},
//   	},
//   	Sensitivity: jsii.String("sensitivity"),
//   	UsageOfAction: jsii.String("usageOfAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideaction.html
//
type CfnWebACLPropsMixin_ClientSideActionProperty struct {
	// The regular expression to match against the web request URI, used to identify requests that can't handle a silent browser challenge.
	//
	// When the `ClientSideAction` setting `UsageOfAction` is enabled, the managed rule group uses this setting to determine which requests to label with `awswaf:managed:aws:anti-ddos:challengeable-request` . If `UsageOfAction` is disabled, this setting has no effect and the managed rule group doesn't add the label to any requests.
	//
	// The anti-DDoS managed rule group doesn't evaluate the rules `ChallengeDDoSRequests` or `ChallengeAllDuringEvent` for web requests whose URIs match this regex. This is true regardless of whether you override the rule action for either of the rules in your web ACL configuration.
	//
	// AWS recommends using a regular expression.
	//
	// This setting is required if `UsageOfAction` is set to `ENABLED` . If required, you can provide between 1 and 5 regex objects in the array of settings.
	//
	// AWS recommends starting with the following setting. Review and update it for your application's needs:
	//
	// `\/api\/|\.(acc|avi|css|gif|jpe?g|js|mp[34]|ogg|otf|pdf|png|tiff?|ttf|webm|webp|woff2?)$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideaction.html#cfn-wafv2-webacl-clientsideaction-exempturiregularexpressions
	//
	ExemptUriRegularExpressions interface{} `field:"optional" json:"exemptUriRegularExpressions" yaml:"exemptUriRegularExpressions"`
	// The sensitivity that the rule group rule `ChallengeDDoSRequests` uses when matching against the DDoS suspicion labeling on a request.
	//
	// The managed rule group adds the labeling during DDoS events, before the `ChallengeDDoSRequests` rule runs.
	//
	// The higher the sensitivity, the more levels of labeling that the rule matches:
	//
	// - Low sensitivity is less sensitive, causing the rule to match only on the most likely participants in an attack, which are the requests with the high suspicion label `awswaf:managed:aws:anti-ddos:high-suspicion-ddos-request` .
	// - Medium sensitivity causes the rule to match on the medium and high suspicion labels.
	// - High sensitivity causes the rule to match on all of the suspicion labels: low, medium, and high.
	//
	// Default: `HIGH`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideaction.html#cfn-wafv2-webacl-clientsideaction-sensitivity
	//
	Sensitivity *string `field:"optional" json:"sensitivity" yaml:"sensitivity"`
	// Determines whether to use the `AWSManagedRulesAntiDDoSRuleSet` rules `ChallengeAllDuringEvent` and `ChallengeDDoSRequests` in the rule group evaluation and the related label `awswaf:managed:aws:anti-ddos:challengeable-request` .
	//
	// - If usage is enabled:
	//
	// - The managed rule group adds the label `awswaf:managed:aws:anti-ddos:challengeable-request` to any web request whose URL does *NOT* match the regular expressions provided in the `ClientSideAction` setting `ExemptUriRegularExpressions` .
	// - The two rules are evaluated against web requests for protected resources that are experiencing a DDoS attack. The two rules only apply their action to matching requests that have the label `awswaf:managed:aws:anti-ddos:challengeable-request` .
	// - If usage is disabled:
	//
	// - The managed rule group doesn't add the label `awswaf:managed:aws:anti-ddos:challengeable-request` to any web requests.
	// - The two rules are not evaluated.
	// - None of the other `ClientSideAction` settings have any effect.
	//
	// > This setting only enables or disables the use of the two anti-DDOS rules `ChallengeAllDuringEvent` and `ChallengeDDoSRequests` in the anti-DDoS managed rule group.
	// >
	// > This setting doesn't alter the action setting in the two rules. To override the actions used by the rules `ChallengeAllDuringEvent` and `ChallengeDDoSRequests` , enable this setting, and then override the rule actions in the usual way, in your managed rule group configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-clientsideaction.html#cfn-wafv2-webacl-clientsideaction-usageofaction
	//
	UsageOfAction *string `field:"optional" json:"usageOfAction" yaml:"usageOfAction"`
}

