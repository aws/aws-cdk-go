package previewawsec2events


// Type definition for GroupSet_2Item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   groupSet2Item := &GroupSet2Item{
//   	GroupId: []*string{
//   		jsii.String("groupId"),
//   	},
//   	GroupName: []*string{
//   		jsii.String("groupName"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_GroupSet2Item struct {
	// groupId property.
	//
	// Specify an array of string values to match this event if the actual value of groupId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupId *[]*string `field:"optional" json:"groupId" yaml:"groupId"`
	// groupName property.
	//
	// Specify an array of string values to match this event if the actual value of groupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupName *[]*string `field:"optional" json:"groupName" yaml:"groupName"`
}

