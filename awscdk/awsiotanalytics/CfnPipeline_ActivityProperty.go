package awsiotanalytics


// An activity that performs a transformation on a message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activityProperty := &ActivityProperty{
//   	AddAttributes: &AddAttributesProperty{
//   		Attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   	Channel: &ChannelProperty{
//   		ChannelName: jsii.String("channelName"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   	Datastore: &DatastoreProperty{
//   		DatastoreName: jsii.String("datastoreName"),
//   		Name: jsii.String("name"),
//   	},
//   	DeviceRegistryEnrich: &DeviceRegistryEnrichProperty{
//   		Attribute: jsii.String("attribute"),
//   		Name: jsii.String("name"),
//   		RoleArn: jsii.String("roleArn"),
//   		ThingName: jsii.String("thingName"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   	DeviceShadowEnrich: &DeviceShadowEnrichProperty{
//   		Attribute: jsii.String("attribute"),
//   		Name: jsii.String("name"),
//   		RoleArn: jsii.String("roleArn"),
//   		ThingName: jsii.String("thingName"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   	Filter: &FilterProperty{
//   		Filter: jsii.String("filter"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   	Lambda: &LambdaProperty{
//   		BatchSize: jsii.Number(123),
//   		LambdaName: jsii.String("lambdaName"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   	Math: &MathProperty{
//   		Attribute: jsii.String("attribute"),
//   		Math: jsii.String("math"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   	RemoveAttributes: &RemoveAttributesProperty{
//   		Attributes: []*string{
//   			jsii.String("attributes"),
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   	SelectAttributes: &SelectAttributesProperty{
//   		Attributes: []*string{
//   			jsii.String("attributes"),
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Next: jsii.String("next"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html
//
type CfnPipeline_ActivityProperty struct {
	// Adds other attributes based on existing attributes in the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-addattributes
	//
	AddAttributes interface{} `field:"optional" json:"addAttributes" yaml:"addAttributes"`
	// Determines the source of the messages to be processed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-channel
	//
	Channel interface{} `field:"optional" json:"channel" yaml:"channel"`
	// Specifies where to store the processed message data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-datastore
	//
	Datastore interface{} `field:"optional" json:"datastore" yaml:"datastore"`
	// Adds data from the AWS IoT device registry to your message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-deviceregistryenrich
	//
	DeviceRegistryEnrich interface{} `field:"optional" json:"deviceRegistryEnrich" yaml:"deviceRegistryEnrich"`
	// Adds information from the AWS IoT Device Shadows service to a message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-deviceshadowenrich
	//
	DeviceShadowEnrich interface{} `field:"optional" json:"deviceShadowEnrich" yaml:"deviceShadowEnrich"`
	// Filters a message based on its attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Runs a Lambda function to modify the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Computes an arithmetic expression using the message's attributes and adds it to the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-math
	//
	Math interface{} `field:"optional" json:"math" yaml:"math"`
	// Removes attributes from a message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-removeattributes
	//
	RemoveAttributes interface{} `field:"optional" json:"removeAttributes" yaml:"removeAttributes"`
	// Creates a new message using only the specified attributes from the original message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-selectattributes
	//
	SelectAttributes interface{} `field:"optional" json:"selectAttributes" yaml:"selectAttributes"`
}

