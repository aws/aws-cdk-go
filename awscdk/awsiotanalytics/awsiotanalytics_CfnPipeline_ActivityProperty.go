package awsiotanalytics


// An activity that performs a transformation on a message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activityProperty := &activityProperty{
//   	addAttributes: &addAttributesProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	channel: &channelProperty{
//   		channelName: jsii.String("channelName"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	datastore: &datastoreProperty{
//   		datastoreName: jsii.String("datastoreName"),
//   		name: jsii.String("name"),
//   	},
//   	deviceRegistryEnrich: &deviceRegistryEnrichProperty{
//   		attribute: jsii.String("attribute"),
//   		name: jsii.String("name"),
//   		roleArn: jsii.String("roleArn"),
//   		thingName: jsii.String("thingName"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	deviceShadowEnrich: &deviceShadowEnrichProperty{
//   		attribute: jsii.String("attribute"),
//   		name: jsii.String("name"),
//   		roleArn: jsii.String("roleArn"),
//   		thingName: jsii.String("thingName"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	filter: &filterProperty{
//   		filter: jsii.String("filter"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	lambda: &lambdaProperty{
//   		batchSize: jsii.Number(123),
//   		lambdaName: jsii.String("lambdaName"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	math: &mathProperty{
//   		attribute: jsii.String("attribute"),
//   		math: jsii.String("math"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	removeAttributes: &removeAttributesProperty{
//   		attributes: []*string{
//   			jsii.String("attributes"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   	selectAttributes: &selectAttributesProperty{
//   		attributes: []*string{
//   			jsii.String("attributes"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		next: jsii.String("next"),
//   	},
//   }
//
type CfnPipeline_ActivityProperty struct {
	// Adds other attributes based on existing attributes in the message.
	AddAttributes interface{} `field:"optional" json:"addAttributes" yaml:"addAttributes"`
	// Determines the source of the messages to be processed.
	Channel interface{} `field:"optional" json:"channel" yaml:"channel"`
	// Specifies where to store the processed message data.
	Datastore interface{} `field:"optional" json:"datastore" yaml:"datastore"`
	// Adds data from the AWS IoT device registry to your message.
	DeviceRegistryEnrich interface{} `field:"optional" json:"deviceRegistryEnrich" yaml:"deviceRegistryEnrich"`
	// Adds information from the AWS IoT Device Shadows service to a message.
	DeviceShadowEnrich interface{} `field:"optional" json:"deviceShadowEnrich" yaml:"deviceShadowEnrich"`
	// Filters a message based on its attributes.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Runs a Lambda function to modify the message.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Computes an arithmetic expression using the message's attributes and adds it to the message.
	Math interface{} `field:"optional" json:"math" yaml:"math"`
	// Removes attributes from a message.
	RemoveAttributes interface{} `field:"optional" json:"removeAttributes" yaml:"removeAttributes"`
	// Creates a new message using only the specified attributes from the original message.
	SelectAttributes interface{} `field:"optional" json:"selectAttributes" yaml:"selectAttributes"`
}

