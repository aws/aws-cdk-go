package previewawsec2events


// Type definition for TagSpecification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagSpecification := &TagSpecification{
//   	ResourceType: []*string{
//   		jsii.String("resourceType"),
//   	},
//   	Tag: &Tag{
//   		Key: []*string{
//   			jsii.String("key"),
//   		},
//   		Tag: []*string{
//   			jsii.String("tag"),
//   		},
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_TagSpecification struct {
	// ResourceType property.
	//
	// Specify an array of string values to match this event if the actual value of ResourceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceType *[]*string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Tag property.
	//
	// Specify an array of string values to match this event if the actual value of Tag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tag *InstanceEvents_AWSAPICallViaCloudTrail_Tag `field:"optional" json:"tag" yaml:"tag"`
}

