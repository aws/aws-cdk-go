package awsglue


// Specifies the sort order of a sorted column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   orderProperty := &orderProperty{
//   	column: jsii.String("column"),
//
//   	// the properties below are optional
//   	sortOrder: jsii.Number(123),
//   }
//
type CfnPartition_OrderProperty struct {
	// The name of the column.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Indicates that the column is sorted in ascending order ( `== 1` ), or in descending order ( `==0` ).
	SortOrder *float64 `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

