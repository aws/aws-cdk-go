package previewawsgroundstationmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisDataStreamDataProperty := &KinesisDataStreamDataProperty{
//   	KinesisDataStreamArn: jsii.String("kinesisDataStreamArn"),
//   	KinesisRoleArn: jsii.String("kinesisRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-kinesisdatastreamdata.html
//
type CfnConfigPropsMixin_KinesisDataStreamDataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-kinesisdatastreamdata.html#cfn-groundstation-config-kinesisdatastreamdata-kinesisdatastreamarn
	//
	KinesisDataStreamArn *string `field:"optional" json:"kinesisDataStreamArn" yaml:"kinesisDataStreamArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-kinesisdatastreamdata.html#cfn-groundstation-config-kinesisdatastreamdata-kinesisrolearn
	//
	KinesisRoleArn *string `field:"optional" json:"kinesisRoleArn" yaml:"kinesisRoleArn"`
}

