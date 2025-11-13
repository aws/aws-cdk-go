package awscognito


// Properties for defining a `CfnTerms`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTermsProps := &CfnTermsProps{
//   	Enforcement: jsii.String("enforcement"),
//   	Links: map[string]*string{
//   		"linksKey": jsii.String("links"),
//   	},
//   	TermsName: jsii.String("termsName"),
//   	TermsSource: jsii.String("termsSource"),
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	ClientId: jsii.String("clientId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html
//
type CfnTermsProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-enforcement
	//
	Enforcement *string `field:"required" json:"enforcement" yaml:"enforcement"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-links
	//
	Links interface{} `field:"required" json:"links" yaml:"links"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-termsname
	//
	TermsName *string `field:"required" json:"termsName" yaml:"termsName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-termssource
	//
	TermsSource *string `field:"required" json:"termsSource" yaml:"termsSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
}

