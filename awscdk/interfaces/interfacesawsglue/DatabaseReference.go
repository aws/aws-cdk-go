package interfacesawsglue


// A reference to a Database resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseReference := &DatabaseReference{
//   	DatabaseName: jsii.String("databaseName"),
//   }
//
type DatabaseReference struct {
	// The DatabaseName of the Database resource.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
}

