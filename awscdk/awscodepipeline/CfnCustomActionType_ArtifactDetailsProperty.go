package awscodepipeline


// Returns information about the details of an artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactDetailsProperty := &ArtifactDetailsProperty{
//   	MaximumCount: jsii.Number(123),
//   	MinimumCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html
//
type CfnCustomActionType_ArtifactDetailsProperty struct {
	// The maximum number of artifacts allowed for the action type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-maximumcount
	//
	MaximumCount *float64 `field:"required" json:"maximumCount" yaml:"maximumCount"`
	// The minimum number of artifacts allowed for the action type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-minimumcount
	//
	MinimumCount *float64 `field:"required" json:"minimumCount" yaml:"minimumCount"`
}

