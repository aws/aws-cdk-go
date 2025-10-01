package awslocation


// A reference to a PlaceIndex resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placeIndexReference := &PlaceIndexReference{
//   	IndexName: jsii.String("indexName"),
//   	PlaceIndexArn: jsii.String("placeIndexArn"),
//   }
//
type PlaceIndexReference struct {
	// The IndexName of the PlaceIndex resource.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The ARN of the PlaceIndex resource.
	PlaceIndexArn *string `field:"required" json:"placeIndexArn" yaml:"placeIndexArn"`
}

