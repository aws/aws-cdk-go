package awss3tables


// A reference to a Table resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableReference := &TableReference{
//   	TableArn: jsii.String("tableArn"),
//   }
//
type TableReference struct {
	// The TableARN of the Table resource.
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
}

