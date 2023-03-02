package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mutualTlsAuthenticationProperty := &mutualTlsAuthenticationProperty{
//   	truststoreUri: jsii.String("truststoreUri"),
//   	truststoreVersion: jsii.Boolean(false),
//   }
//
type CfnHttpApi_MutualTlsAuthenticationProperty struct {
	// `CfnHttpApi.MutualTlsAuthenticationProperty.TruststoreUri`.
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// `CfnHttpApi.MutualTlsAuthenticationProperty.TruststoreVersion`.
	TruststoreVersion interface{} `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

