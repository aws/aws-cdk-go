package awsiotanalytics


// An activity that adds data from the AWS IoT device registry to your message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceRegistryEnrichProperty := &DeviceRegistryEnrichProperty{
//   	Attribute: jsii.String("attribute"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	ThingName: jsii.String("thingName"),
//
//   	// the properties below are optional
//   	Next: jsii.String("next"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html
//
type CfnPipeline_DeviceRegistryEnrichProperty struct {
	// The name of the attribute that is added to the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-attribute
	//
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The name of the 'deviceRegistryEnrich' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the role that allows access to the device's registry information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the IoT device whose registry information is added to the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-thingname
	//
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html#cfn-iotanalytics-pipeline-deviceregistryenrich-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
}

