package previewawsec2events


// Type definition for TagSpecificationSetItemItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagSpecificationSetItemItem := &TagSpecificationSetItemItem{
//   	Key: []*string{
//   		jsii.String("key"),
//   	},
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_TagSpecificationSetItemItem struct {
	// key property.
	//
	// Specify an array of string values to match this event if the actual value of key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Key *[]*string `field:"optional" json:"key" yaml:"key"`
	// value property.
	//
	// Specify an array of string values to match this event if the actual value of value is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

