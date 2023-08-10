package awscdkredshiftalpha


// A column in a Redshift table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import redshift_alpha "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//
//   column := &Column{
//   	DataType: jsii.String("dataType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   	DistKey: jsii.Boolean(false),
//   	Encoding: redshift_alpha.ColumnEncoding_AUTO,
//   	Id: jsii.String("id"),
//   	SortKey: jsii.Boolean(false),
//   }
//
// Experimental.
type Column struct {
	// The data type of the column.
	// Experimental.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The name of the column.
	//
	// This will appear on Amazon Redshift.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to attach to the column.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Boolean value that indicates whether the column is to be configured as DISTKEY.
	// Experimental.
	DistKey *bool `field:"optional" json:"distKey" yaml:"distKey"`
	// The encoding to use for the column.
	// Experimental.
	Encoding ColumnEncoding `field:"optional" json:"encoding" yaml:"encoding"`
	// The unique identifier of the column.
	//
	// This is not the name of the column, and renaming this identifier will cause a new column to be created and the old column to be dropped.
	//
	// **NOTE** - This field will be set, however, only by setting the `@aws-cdk/aws-redshift:columnId` feature flag will this field be used.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Boolean value that indicates whether the column is to be configured as SORTKEY.
	// Experimental.
	SortKey *bool `field:"optional" json:"sortKey" yaml:"sortKey"`
}

