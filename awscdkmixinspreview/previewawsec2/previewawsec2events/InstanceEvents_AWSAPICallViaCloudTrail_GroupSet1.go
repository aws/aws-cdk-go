package previewawsec2events


// Type definition for GroupSet_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   groupSet1 := &GroupSet1{
//   	Items: []GroupSet1Item{
//   		&GroupSet1Item{
//   			GroupId: []*string{
//   				jsii.String("groupId"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_GroupSet1 struct {
	// items property.
	//
	// Specify an array of string values to match this event if the actual value of items is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Items *[]*InstanceEvents_AWSAPICallViaCloudTrail_GroupSet1Item `field:"optional" json:"items" yaml:"items"`
}

