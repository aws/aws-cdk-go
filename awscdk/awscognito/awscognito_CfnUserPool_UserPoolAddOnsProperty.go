package awscognito


// The user pool add-ons type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolAddOnsProperty := &userPoolAddOnsProperty{
//   	advancedSecurityMode: jsii.String("advancedSecurityMode"),
//   }
//
type CfnUserPool_UserPoolAddOnsProperty struct {
	// The advanced security mode.
	AdvancedSecurityMode *string `field:"optional" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
}

