package awsmsk


// Record schema configuration for a topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordSchemaProperty := &RecordSchemaProperty{
//   	GsrArn: jsii.String("gsrArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-recordschema.html
//
type CfnChannel_RecordSchemaProperty struct {
	// ARN of Glue Schema Registry resource used for table schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-recordschema.html#cfn-msk-channel-recordschema-gsrarn
	//
	GsrArn *string `field:"required" json:"gsrArn" yaml:"gsrArn"`
}

