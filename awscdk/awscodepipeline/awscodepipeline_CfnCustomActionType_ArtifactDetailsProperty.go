package awscodepipeline


// Returns information about the details of an artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactDetailsProperty := &artifactDetailsProperty{
//   	maximumCount: jsii.Number(123),
//   	minimumCount: jsii.Number(123),
//   }
//
type CfnCustomActionType_ArtifactDetailsProperty struct {
	// The maximum number of artifacts allowed for the action type.
	MaximumCount *float64 `field:"required" json:"maximumCount" yaml:"maximumCount"`
	// The minimum number of artifacts allowed for the action type.
	MinimumCount *float64 `field:"required" json:"minimumCount" yaml:"minimumCount"`
}

