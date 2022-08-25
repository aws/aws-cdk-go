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
//   sourceAccessConfiguration := &sourceAccessConfiguration{
//   	type: sourceAccessConfigurationType,
//   	uri: jsii.String("uri"),
//   }
//
type SourceAccessConfiguration struct {
	// The type of authentication protocol or the VPC components for your event source.
	//
	// For example: "SASL_SCRAM_512_AUTH".
	Type SourceAccessConfigurationType `field:"required" json:"type" yaml:"type"`
	// The value for your chosen configuration in type.
	//
	// For example: "URI": "arn:aws:secretsmanager:us-east-1:01234567890:secret:MyBrokerSecretName".
	// The exact string depends on the type.
	// See: SourceAccessConfigurationType.
	//
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

