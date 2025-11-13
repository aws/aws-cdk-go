package interfacesawsmacie


// A reference to a CustomDataIdentifier resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDataIdentifierReference := &CustomDataIdentifierReference{
//   	CustomDataIdentifierArn: jsii.String("customDataIdentifierArn"),
//   	CustomDataIdentifierId: jsii.String("customDataIdentifierId"),
//   }
//
type CustomDataIdentifierReference struct {
	// The ARN of the CustomDataIdentifier resource.
	CustomDataIdentifierArn *string `field:"required" json:"customDataIdentifierArn" yaml:"customDataIdentifierArn"`
	// The Id of the CustomDataIdentifier resource.
	CustomDataIdentifierId *string `field:"required" json:"customDataIdentifierId" yaml:"customDataIdentifierId"`
}

