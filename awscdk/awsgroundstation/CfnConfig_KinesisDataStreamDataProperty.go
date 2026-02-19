package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisDataStreamDataProperty := &KinesisDataStreamDataProperty{
//   	KinesisDataStreamArn: jsii.String("kinesisDataStreamArn"),
//   	KinesisRoleArn: jsii.String("kinesisRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-kinesisdatastreamdata.html
//
type CfnConfig_KinesisDataStreamDataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-kinesisdatastreamdata.html#cfn-groundstation-config-kinesisdatastreamdata-kinesisdatastreamarn
	//
	KinesisDataStreamArn *string `field:"required" json:"kinesisDataStreamArn" yaml:"kinesisDataStreamArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-kinesisdatastreamdata.html#cfn-groundstation-config-kinesisdatastreamdata-kinesisrolearn
	//
	KinesisRoleArn *string `field:"required" json:"kinesisRoleArn" yaml:"kinesisRoleArn"`
}

