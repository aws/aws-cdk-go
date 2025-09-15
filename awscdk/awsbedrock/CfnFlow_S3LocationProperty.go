package awsbedrock


// The S3 location of the flow definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-s3location.html
//
type CfnFlow_S3LocationProperty struct {
	// The S3 bucket containing the flow definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-s3location.html#cfn-bedrock-flow-s3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The object key for the S3 location containing the definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-s3location.html#cfn-bedrock-flow-s3location-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Amazon S3 location from which to retrieve data for an S3 retrieve node or to which to store data for an S3 storage node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-s3location.html#cfn-bedrock-flow-s3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

