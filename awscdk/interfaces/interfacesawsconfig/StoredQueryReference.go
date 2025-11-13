package interfacesawsconfig


// A reference to a StoredQuery resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storedQueryReference := &StoredQueryReference{
//   	QueryName: jsii.String("queryName"),
//   }
//
type StoredQueryReference struct {
	// The QueryName of the StoredQuery resource.
	QueryName *string `field:"required" json:"queryName" yaml:"queryName"`
}

