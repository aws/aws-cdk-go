package interfacesawsiam


// A reference to a SAMLProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAMLProviderReference := &SAMLProviderReference{
//   	SamlProviderArn: jsii.String("samlProviderArn"),
//   }
//
type SAMLProviderReference struct {
	// The Arn of the SAMLProvider resource.
	SamlProviderArn *string `field:"required" json:"samlProviderArn" yaml:"samlProviderArn"`
}

