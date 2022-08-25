package awswafv2


// Determines how long a `CAPTCHA` token remains valid after the client successfully solves a `CAPTCHA` puzzle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   immunityTimePropertyProperty := &immunityTimePropertyProperty{
//   	immunityTime: jsii.Number(123),
//   }
//
type CfnWebACL_ImmunityTimePropertyProperty struct {
	// The amount of time, in seconds, that a `CAPTCHA` token is valid.
	//
	// The default setting is 300.
	ImmunityTime *float64 `field:"required" json:"immunityTime" yaml:"immunityTime"`
}

