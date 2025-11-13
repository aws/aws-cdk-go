package interfacesawsiot


// A reference to a ThingGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thingGroupReference := &ThingGroupReference{
//   	ThingGroupArn: jsii.String("thingGroupArn"),
//   	ThingGroupName: jsii.String("thingGroupName"),
//   }
//
type ThingGroupReference struct {
	// The ARN of the ThingGroup resource.
	ThingGroupArn *string `field:"required" json:"thingGroupArn" yaml:"thingGroupArn"`
	// The ThingGroupName of the ThingGroup resource.
	ThingGroupName *string `field:"required" json:"thingGroupName" yaml:"thingGroupName"`
}

