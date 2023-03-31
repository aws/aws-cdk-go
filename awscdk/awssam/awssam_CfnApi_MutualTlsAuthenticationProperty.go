package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mutualTlsAuthenticationProperty := &mutualTlsAuthenticationProperty{
//   	truststoreUri: jsii.String("truststoreUri"),
//   	truststoreVersion: jsii.String("truststoreVersion"),
//   }
//
type CfnApi_MutualTlsAuthenticationProperty struct {
	// `CfnApi.MutualTlsAuthenticationProperty.TruststoreUri`.
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// `CfnApi.MutualTlsAuthenticationProperty.TruststoreVersion`.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

