package previewawsobservabilityadminmixins


// Condition that matches based on the specific WAF action taken on the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionConditionProperty := &ActionConditionProperty{
//   	Action: jsii.String("action"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-actioncondition.html
//
type CfnTelemetryRulePropsMixin_ActionConditionProperty struct {
	// The WAF action to match against (ALLOW, BLOCK, COUNT, CAPTCHA, CHALLENGE, EXCLUDED_AS_COUNT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-actioncondition.html#cfn-observabilityadmin-telemetryrule-actioncondition-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
}

