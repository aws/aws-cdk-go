package interfacesawstimestream


// A reference to a Database resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseReference := &DatabaseReference{
//   	DatabaseArn: jsii.String("databaseArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   }
//
type DatabaseReference struct {
	// The ARN of the Database resource.
	DatabaseArn *string `field:"required" json:"databaseArn" yaml:"databaseArn"`
	// The DatabaseName of the Database resource.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
}

