package awssagemaker


// Describes the location of the channel data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceProperty := &dataSourceProperty{
//   	s3DataSource: &s3DataSourceProperty{
//   		s3DataType: jsii.String("s3DataType"),
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnModelPackage_DataSourceProperty struct {
	// The S3 location of the data source that is associated with a channel.
	S3DataSource interface{} `field:"required" json:"s3DataSource" yaml:"s3DataSource"`
}

