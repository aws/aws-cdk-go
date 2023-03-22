package awsiotanalytics


// Contains information about a column that stores your data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnProperty := &columnProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   }
//
type CfnDatastore_ColumnProperty struct {
	// The name of the column.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of data.
	//
	// For more information about the supported data types, see [Common data types](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-common.html) in the *AWS Glue Developer Guide* .
	Type *string `field:"required" json:"type" yaml:"type"`
}

