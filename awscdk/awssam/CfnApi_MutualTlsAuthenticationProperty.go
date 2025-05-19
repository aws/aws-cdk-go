package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mutualTlsAuthenticationProperty := &MutualTlsAuthenticationProperty{
//   	TruststoreUri: jsii.String("truststoreUri"),
//   	TruststoreVersion: jsii.String("truststoreVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-mutualtlsauthentication.html
//
type CfnApi_MutualTlsAuthenticationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-mutualtlsauthentication.html#cfn-serverless-api-mutualtlsauthentication-truststoreuri
	//
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-mutualtlsauthentication.html#cfn-serverless-api-mutualtlsauthentication-truststoreversion
	//
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

