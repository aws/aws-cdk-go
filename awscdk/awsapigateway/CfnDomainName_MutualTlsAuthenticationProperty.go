package awsapigateway


// The mutual TLS authentication configuration for a custom domain name.
//
// If specified, API Gateway performs two-way authentication between the client and the server. Clients must present a trusted certificate to access your API.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html
//
type CfnDomainName_MutualTlsAuthenticationProperty struct {
	// An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example `s3://bucket-name/key-name` .
	//
	// The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html#cfn-apigateway-domainname-mutualtlsauthentication-truststoreuri
	//
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// The version of the S3 object that contains your truststore.
	//
	// To specify a version, you must have versioning enabled for the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html#cfn-apigateway-domainname-mutualtlsauthentication-truststoreversion
	//
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

