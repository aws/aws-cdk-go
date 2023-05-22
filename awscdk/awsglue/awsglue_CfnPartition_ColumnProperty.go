package awsglue


// A column in a `Table` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnProperty := &ColumnProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   	Type: jsii.String("type"),
//   }
//
type CfnPartition_ColumnProperty struct {
	// The name of the `Column` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// A free-form text comment.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The data type of the `Column` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

