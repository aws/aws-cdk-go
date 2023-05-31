package awsredshift


// A column in a Redshift table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   column := &Column{
//   	DataType: jsii.String("dataType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DistKey: jsii.Boolean(false),
//   	SortKey: jsii.Boolean(false),
//   }
//
// Experimental.
type Column struct {
	// The data type of the column.
	// Experimental.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The name of the column.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Boolean value that indicates whether the column is to be configured as DISTKEY.
	// Experimental.
	DistKey *bool `field:"optional" json:"distKey" yaml:"distKey"`
	// Boolean value that indicates whether the column is to be configured as SORTKEY.
	// Experimental.
	SortKey *bool `field:"optional" json:"sortKey" yaml:"sortKey"`
}

