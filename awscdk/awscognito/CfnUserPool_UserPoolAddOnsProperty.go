package awscognito


// The user pool add-ons type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolAddOnsProperty := &UserPoolAddOnsProperty{
//   	AdvancedSecurityMode: jsii.String("advancedSecurityMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html
//
type CfnUserPool_UserPoolAddOnsProperty struct {
	// The advanced security mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html#cfn-cognito-userpool-userpooladdons-advancedsecuritymode
	//
	AdvancedSecurityMode *string `field:"optional" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
}

