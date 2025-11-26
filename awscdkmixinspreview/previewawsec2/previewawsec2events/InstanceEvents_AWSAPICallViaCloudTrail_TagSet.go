package previewawsec2events


// Type definition for TagSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagSet := &TagSet{
//   	Items: []TagSpecificationSetItemItem{
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
type InstanceEvents_AWSAPICallViaCloudTrail_TagSet struct {
	// items property.
	//
	// Specify an array of string values to match this event if the actual value of items is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Items *[]*InstanceEvents_AWSAPICallViaCloudTrail_TagSpecificationSetItemItem `field:"optional" json:"items" yaml:"items"`
}

