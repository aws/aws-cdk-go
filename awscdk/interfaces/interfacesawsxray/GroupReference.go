package interfacesawsxray


// A reference to a Group resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupReference := &GroupReference{
//   	GroupArn: jsii.String("groupArn"),
//   }
//
type GroupReference struct {
	// The GroupARN of the Group resource.
	GroupArn *string `field:"required" json:"groupArn" yaml:"groupArn"`
}

