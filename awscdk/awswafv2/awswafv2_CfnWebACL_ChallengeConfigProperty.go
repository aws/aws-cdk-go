package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   challengeConfigProperty := &challengeConfigProperty{
//   	immunityTimeProperty: &immunityTimePropertyProperty{
//   		immunityTime: jsii.Number(123),
//   	},
//   }
//
type CfnWebACL_ChallengeConfigProperty struct {
	// `CfnWebACL.ChallengeConfigProperty.ImmunityTimeProperty`.
	ImmunityTimeProperty interface{} `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

