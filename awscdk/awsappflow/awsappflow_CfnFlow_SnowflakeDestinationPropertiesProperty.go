package awsappflow


// The properties that are applied when Snowflake is being used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeDestinationPropertiesProperty := &snowflakeDestinationPropertiesProperty{
//   	intermediateBucketName: jsii.String("intermediateBucketName"),
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   }
//
type CfnFlow_SnowflakeDestinationPropertiesProperty struct {
	// The intermediate bucket that Amazon AppFlow uses when moving data into Snowflake.
	IntermediateBucketName *string `field:"required" json:"intermediateBucketName" yaml:"intermediateBucketName"`
	// The object specified in the Snowflake flow destination.
	Object *string `field:"required" json:"object" yaml:"object"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Snowflake destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

