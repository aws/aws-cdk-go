package previewawsgroundstationmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnConfigPropsMixin_TelemetrySinkConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-telemetrysinkconfig.html#cfn-groundstation-config-telemetrysinkconfig-telemetrysinkdata
	//
	TelemetrySinkData interface{} `field:"optional" json:"telemetrySinkData" yaml:"telemetrySinkData"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-telemetrysinkconfig.html#cfn-groundstation-config-telemetrysinkconfig-telemetrysinktype
	//
	TelemetrySinkType *string `field:"optional" json:"telemetrySinkType" yaml:"telemetrySinkType"`
}

