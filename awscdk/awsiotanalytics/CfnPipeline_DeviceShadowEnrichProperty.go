package awsiotanalytics


// An activity that adds information from the AWS IoT Device Shadows service to a message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceShadowEnrichProperty := &DeviceShadowEnrichProperty{
//   	Attribute: jsii.String("attribute"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	ThingName: jsii.String("thingName"),
//
//   	// the properties below are optional
//   	Next: jsii.String("next"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceshadowenrich.html
//
type CfnPipeline_DeviceShadowEnrichProperty struct {
	// The name of the attribute that is added to the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceshadowenrich.html#cfn-iotanalytics-pipeline-deviceshadowenrich-attribute
	//
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The name of the 'deviceShadowEnrich' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceshadowenrich.html#cfn-iotanalytics-pipeline-deviceshadowenrich-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the role that allows access to the device's shadow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceshadowenrich.html#cfn-iotanalytics-pipeline-deviceshadowenrich-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the IoT device whose shadow information is added to the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceshadowenrich.html#cfn-iotanalytics-pipeline-deviceshadowenrich-thingname
	//
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceshadowenrich.html#cfn-iotanalytics-pipeline-deviceshadowenrich-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
}

