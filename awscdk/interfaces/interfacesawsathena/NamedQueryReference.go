package interfacesawsathena


// A reference to a NamedQuery resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namedQueryReference := &NamedQueryReference{
//   	NamedQueryId: jsii.String("namedQueryId"),
//   }
//
type NamedQueryReference struct {
	// The NamedQueryId of the NamedQuery resource.
	NamedQueryId *string `field:"required" json:"namedQueryId" yaml:"namedQueryId"`
}

