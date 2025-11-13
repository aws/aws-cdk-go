package interfacesawslocation


// A reference to a Map resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mapReference := &MapReference{
//   	MapArn: jsii.String("mapArn"),
//   	MapName: jsii.String("mapName"),
//   }
//
type MapReference struct {
	// The ARN of the Map resource.
	MapArn *string `field:"required" json:"mapArn" yaml:"mapArn"`
	// The MapName of the Map resource.
	MapName *string `field:"required" json:"mapName" yaml:"mapName"`
}

