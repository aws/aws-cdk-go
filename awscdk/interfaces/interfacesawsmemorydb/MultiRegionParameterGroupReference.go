package interfacesawsmemorydb


// A reference to a MultiRegionParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiRegionParameterGroupReference := &MultiRegionParameterGroupReference{
//   	MultiRegionParameterGroupArn: jsii.String("multiRegionParameterGroupArn"),
//   }
//
type MultiRegionParameterGroupReference struct {
	// The Arn of the MultiRegionParameterGroup resource.
	MultiRegionParameterGroupArn *string `field:"required" json:"multiRegionParameterGroupArn" yaml:"multiRegionParameterGroupArn"`
}

