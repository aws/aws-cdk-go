package awseventsv2


// Configuration for transforming events before delivery: the raw payload, the payload with its metadata envelope, or the output of a JSONata expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   transformerProperty := &TransformerProperty{
//   	JsonataConfiguration: &JsonataConfigurationProperty{
//   		Expression: jsii.String("expression"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-transformer.html
//
type CfnSubscriberPropsMixin_TransformerProperty struct {
	// JSONata transform settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-transformer.html#cfn-eventsv2-subscriber-transformer-jsonataconfiguration
	//
	JsonataConfiguration interface{} `field:"optional" json:"jsonataConfiguration" yaml:"jsonataConfiguration"`
	// The transform type: RAW delivers the event payload only;
	//
	// WITH_METADATA delivers the event with its metadata envelope; JSONATA delivers the output of the JSONata expression in JsonataConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-transformer.html#cfn-eventsv2-subscriber-transformer-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

