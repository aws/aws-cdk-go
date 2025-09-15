package awsgamelift


// A reference to a Location resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationReference := &LocationReference{
//   	LocationArn: jsii.String("locationArn"),
//   	LocationName: jsii.String("locationName"),
//   }
//
type LocationReference struct {
	// The ARN of the Location resource.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
	// The LocationName of the Location resource.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}

