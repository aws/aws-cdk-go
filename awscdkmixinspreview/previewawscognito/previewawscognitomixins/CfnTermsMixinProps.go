package previewawscognitomixins


// Properties for CfnTermsPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTermsMixinProps := &CfnTermsMixinProps{
//   	ClientId: jsii.String("clientId"),
//   	Enforcement: jsii.String("enforcement"),
//   	Links: map[string]*string{
//   		"linksKey": jsii.String("links"),
//   	},
//   	TermsName: jsii.String("termsName"),
//   	TermsSource: jsii.String("termsSource"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html
//
type CfnTermsMixinProps struct {
	// The ID of the app client that the terms documents are assigned to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// This parameter is reserved for future use and currently accepts one value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-enforcement
	//
	Enforcement *string `field:"optional" json:"enforcement" yaml:"enforcement"`
	// A map of URLs to languages.
	//
	// For each localized language that will view the requested `TermsName` , assign a URL. A selection of `cognito:default` displays for all languages that don't have a language-specific URL.
	//
	// For example, `"cognito:default": "https://terms.example.com", "cognito:spanish": "https://terms.example.com/es"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-links
	//
	Links interface{} `field:"optional" json:"links" yaml:"links"`
	// The type and friendly name of the terms documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-termsname
	//
	TermsName *string `field:"optional" json:"termsName" yaml:"termsName"`
	// This parameter is reserved for future use and currently accepts one value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-termssource
	//
	TermsSource *string `field:"optional" json:"termsSource" yaml:"termsSource"`
	// The ID of the user pool that contains the terms documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-terms.html#cfn-cognito-terms-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

