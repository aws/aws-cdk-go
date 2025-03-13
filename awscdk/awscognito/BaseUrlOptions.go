package awscognito


// Options to customize the behaviour of `baseUrl()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseUrlOptions := &BaseUrlOptions{
//   	Fips: jsii.Boolean(false),
//   }
//
type BaseUrlOptions struct {
	// Whether to return the FIPS-compliant endpoint.
	// Default: return the standard URL.
	//
	Fips *bool `field:"optional" json:"fips" yaml:"fips"`
}

