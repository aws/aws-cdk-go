package interfacesawsstepfunctions


// A reference to a MapRun resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mapRunReference := &MapRunReference{
//   	MapRunArn: jsii.String("mapRunArn"),
//   }
//
type MapRunReference struct {
	// The MapRunArn of the MapRun resource.
	MapRunArn *string `field:"required" json:"mapRunArn" yaml:"mapRunArn"`
}

