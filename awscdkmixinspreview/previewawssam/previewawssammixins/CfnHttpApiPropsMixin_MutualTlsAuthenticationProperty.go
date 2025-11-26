package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mutualTlsAuthenticationProperty := &MutualTlsAuthenticationProperty{
//   	TruststoreUri: jsii.String("truststoreUri"),
//   	TruststoreVersion: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-mutualtlsauthentication.html
//
type CfnHttpApiPropsMixin_MutualTlsAuthenticationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-mutualtlsauthentication.html#cfn-serverless-httpapi-mutualtlsauthentication-truststoreuri
	//
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-mutualtlsauthentication.html#cfn-serverless-httpapi-mutualtlsauthentication-truststoreversion
	//
	TruststoreVersion interface{} `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

