package awsglue


// A reference to a Table resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableReference := &TableReference{
//   	TableId: jsii.String("tableId"),
//   }
//
type TableReference struct {
	// The Id of the Table resource.
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

