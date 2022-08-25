package awsiotanalytics


// An activity that adds data from the AWS IoT device registry to your message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceRegistryEnrichProperty := &deviceRegistryEnrichProperty{
//   	attribute: jsii.String("attribute"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	thingName: jsii.String("thingName"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_DeviceRegistryEnrichProperty struct {
	// The name of the attribute that is added to the message.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The name of the 'deviceRegistryEnrich' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the role that allows access to the device's registry information.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the IoT device whose registry information is added to the message.
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

