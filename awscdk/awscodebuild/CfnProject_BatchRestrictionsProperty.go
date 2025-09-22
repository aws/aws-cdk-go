package awscodebuild


// Specifies restrictions for the batch build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchRestrictionsProperty := &BatchRestrictionsProperty{
//   	ComputeTypesAllowed: []*string{
//   		jsii.String("computeTypesAllowed"),
//   	},
//   	MaximumBuildsAllowed: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-batchrestrictions.html
//
type CfnProject_BatchRestrictionsProperty struct {
	// An array of strings that specify the compute types that are allowed for the batch build.
	//
	// See [Build environment compute types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html) in the *AWS CodeBuild User Guide* for these values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-batchrestrictions.html#cfn-codebuild-project-batchrestrictions-computetypesallowed
	//
	ComputeTypesAllowed *[]*string `field:"optional" json:"computeTypesAllowed" yaml:"computeTypesAllowed"`
	// Specifies the maximum number of builds allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-batchrestrictions.html#cfn-codebuild-project-batchrestrictions-maximumbuildsallowed
	//
	MaximumBuildsAllowed *float64 `field:"optional" json:"maximumBuildsAllowed" yaml:"maximumBuildsAllowed"`
}

