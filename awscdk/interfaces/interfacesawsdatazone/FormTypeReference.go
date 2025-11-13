package interfacesawsdatazone


// A reference to a FormType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formTypeReference := &FormTypeReference{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	FormTypeIdentifier: jsii.String("formTypeIdentifier"),
//   }
//
type FormTypeReference struct {
	// The DomainIdentifier of the FormType resource.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The FormTypeIdentifier of the FormType resource.
	FormTypeIdentifier *string `field:"required" json:"formTypeIdentifier" yaml:"formTypeIdentifier"`
}

