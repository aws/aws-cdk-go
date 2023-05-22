package awslambda


// Specific settings like the authentication protocol or the VPC components to secure access to your event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceAccessConfigurationType sourceAccessConfigurationType
//
//   sourceAccessConfiguration := &SourceAccessConfiguration{
//   	Type: sourceAccessConfigurationType,
//   	Uri: jsii.String("uri"),
//   }
//
// Experimental.
type SourceAccessConfiguration struct {
	// The type of authentication protocol or the VPC components for your event source.
	//
	// For example: "SASL_SCRAM_512_AUTH".
	// Experimental.
	Type SourceAccessConfigurationType `field:"required" json:"type" yaml:"type"`
	// The value for your chosen configuration in type.
	//
	// For example: "URI": "arn:aws:secretsmanager:us-east-1:01234567890:secret:MyBrokerSecretName".
	// The exact string depends on the type.
	// See: SourceAccessConfigurationType.
	//
	// Experimental.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

