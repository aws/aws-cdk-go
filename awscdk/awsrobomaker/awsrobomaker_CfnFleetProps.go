package awsrobomaker


// Properties for defining a `CfnFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetProps := &cfnFleetProps{
//   	name: jsii.String("name"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnFleetProps struct {
	// The name of the fleet.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of all tags added to the fleet.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

