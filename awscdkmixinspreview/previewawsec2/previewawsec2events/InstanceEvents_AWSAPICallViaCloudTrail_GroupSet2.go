package previewawsec2events


// Type definition for GroupSet_2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   groupSet2 := &GroupSet2{
//   	Items: []GroupSet2Item{
//   		&GroupSet2Item{
//   			GroupId: []*string{
//   				jsii.String("groupId"),
//   			},
//   			GroupName: []*string{
//   				jsii.String("groupName"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_GroupSet2 struct {
	// items property.
	//
	// Specify an array of string values to match this event if the actual value of items is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Items *[]*InstanceEvents_AWSAPICallViaCloudTrail_GroupSet2Item `field:"optional" json:"items" yaml:"items"`
}

