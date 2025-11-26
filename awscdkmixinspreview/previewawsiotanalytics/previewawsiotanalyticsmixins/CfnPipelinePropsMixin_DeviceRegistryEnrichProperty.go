package previewawsiotanalyticsmixins


// An activity that adds data from the AWS IoT device registry to your message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deviceRegistryEnrichProperty := &DeviceRegistryEnrichProperty{
//   	Attribute: jsii.String("attribute"),
//   	Name: jsii.String("name"),
//   	Next: jsii.String("next"),
//   	RoleArn: jsii.String("roleArn"),
//   	ThingName: jsii.String("thingName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html
//
type CfnPipelinePropsMixin_DeviceRegistryEnrichProperty struct {
	// The name of the attribute that is added to the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// The name of the 'deviceRegistryEnrich' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
	// The ARN of the role that allows access to the device's registry information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name of the IoT device whose registry information is added to the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-thingname
	//
	ThingName *string `field:"optional" json:"thingName" yaml:"thingName"`
}

