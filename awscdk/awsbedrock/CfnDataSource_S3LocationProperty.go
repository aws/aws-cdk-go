package awsbedrock


// An Amazon S3 location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3location.html
//
type CfnDataSource_S3LocationProperty struct {
	// The location's URI.
	//
	// For example, `s3://my-bucket/chunk-processor/` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3location.html#cfn-bedrock-datasource-s3location-uri
	//
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

