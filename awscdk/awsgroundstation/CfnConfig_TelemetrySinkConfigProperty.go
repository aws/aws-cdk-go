package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetrySinkConfigProperty := &TelemetrySinkConfigProperty{
//   	TelemetrySinkData: &TelemetrySinkDataProperty{
//   		KinesisDataStreamData: &KinesisDataStreamDataProperty{
//   			KinesisDataStreamArn: jsii.String("kinesisDataStreamArn"),
//   			KinesisRoleArn: jsii.String("kinesisRoleArn"),
//   		},
//   	},
//   	TelemetrySinkType: jsii.String("telemetrySinkType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-telemetrysinkconfig.html
//
type CfnConfig_TelemetrySinkConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-telemetrysinkconfig.html#cfn-groundstation-config-telemetrysinkconfig-telemetrysinkdata
	//
	TelemetrySinkData interface{} `field:"required" json:"telemetrySinkData" yaml:"telemetrySinkData"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-telemetrysinkconfig.html#cfn-groundstation-config-telemetrysinkconfig-telemetrysinktype
	//
	TelemetrySinkType *string `field:"required" json:"telemetrySinkType" yaml:"telemetrySinkType"`
}

