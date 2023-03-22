package awscodebuild


// Specifies restrictions for the batch build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchRestrictionsProperty := &batchRestrictionsProperty{
//   	computeTypesAllowed: []*string{
//   		jsii.String("computeTypesAllowed"),
//   	},
//   	maximumBuildsAllowed: jsii.Number(123),
//   }
//
type CfnProject_BatchRestrictionsProperty struct {
	// An array of strings that specify the compute types that are allowed for the batch build.
	//
	// See [Build environment compute types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html) in the *AWS CodeBuild User Guide* for these values.
	ComputeTypesAllowed *[]*string `field:"optional" json:"computeTypesAllowed" yaml:"computeTypesAllowed"`
	// Specifies the maximum number of builds allowed.
	MaximumBuildsAllowed *float64 `field:"optional" json:"maximumBuildsAllowed" yaml:"maximumBuildsAllowed"`
}

