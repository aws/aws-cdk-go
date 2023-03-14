package awsiot


// Metadata attributes of the time series that are written in each measure record.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestreamDimensionProperty := &timestreamDimensionProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnTopicRule_TimestreamDimensionProperty struct {
	// The metadata dimension name.
	//
	// This is the name of the column in the Amazon Timestream database table record.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value to write in this column of the database record.
	Value *string `field:"required" json:"value" yaml:"value"`
}

