package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetrySinkDataProperty := &TelemetrySinkDataProperty{
//   	KinesisDataStreamData: &KinesisDataStreamDataProperty{
//   		KinesisDataStreamArn: jsii.String("kinesisDataStreamArn"),
//   		KinesisRoleArn: jsii.String("kinesisRoleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-telemetrysinkdata.html
//
type CfnConfig_TelemetrySinkDataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-telemetrysinkdata.html#cfn-groundstation-config-telemetrysinkdata-kinesisdatastreamdata
	//
	KinesisDataStreamData interface{} `field:"required" json:"kinesisDataStreamData" yaml:"kinesisDataStreamData"`
}

