package awswafv2


// Specifies how AWS WAF should handle `Challenge` evaluations.
//
// This is available at the web ACL level and in each rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   challengeConfigProperty := &ChallengeConfigProperty{
//   	ImmunityTimeProperty: &ImmunityTimePropertyProperty{
//   		ImmunityTime: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-challengeconfig.html
//
type CfnWebACL_ChallengeConfigProperty struct {
	// Determines how long a challenge timestamp in the token remains valid after the client successfully responds to a challenge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-challengeconfig.html#cfn-wafv2-webacl-challengeconfig-immunitytimeproperty
	//
	ImmunityTimeProperty interface{} `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

