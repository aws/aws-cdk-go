package previewawsec2events


// Type definition for TagSpecificationSetItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagSpecificationSetItem := &TagSpecificationSetItem{
//   	ResourceType: []*string{
//   		jsii.String("resourceType"),
//   	},
//   	Tags: []TagSpecificationSetItemItem{
//   		&TagSpecificationSetItemItem{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_TagSpecificationSetItem struct {
	// resourceType property.
	//
	// Specify an array of string values to match this event if the actual value of resourceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceType *[]*string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*InstanceEvents_AWSAPICallViaCloudTrail_TagSpecificationSetItemItem `field:"optional" json:"tags" yaml:"tags"`
}

