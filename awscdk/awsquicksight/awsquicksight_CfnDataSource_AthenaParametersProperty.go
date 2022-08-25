package awsquicksight


// Parameters for Amazon Athena.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   athenaParametersProperty := &athenaParametersProperty{
//   	workGroup: jsii.String("workGroup"),
//   }
//
type CfnDataSource_AthenaParametersProperty struct {
	// The workgroup that Amazon Athena uses.
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

