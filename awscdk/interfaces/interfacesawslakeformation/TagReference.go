package interfacesawslakeformation


// A reference to a Tag resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagReference := &TagReference{
//   	TagKey: jsii.String("tagKey"),
//   }
//
type TagReference struct {
	// The TagKey of the Tag resource.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
}

