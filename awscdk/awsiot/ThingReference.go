package awsiot


// A reference to a Thing resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thingReference := &ThingReference{
//   	ThingArn: jsii.String("thingArn"),
//   	ThingName: jsii.String("thingName"),
//   }
//
type ThingReference struct {
	// The ARN of the Thing resource.
	ThingArn *string `field:"required" json:"thingArn" yaml:"thingArn"`
	// The ThingName of the Thing resource.
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
}

