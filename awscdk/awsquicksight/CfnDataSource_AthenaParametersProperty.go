package awsquicksight


// Parameters for Amazon Athena.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   athenaParametersProperty := &AthenaParametersProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	WorkGroup: jsii.String("workGroup"),
//   }
//
type CfnDataSource_AthenaParametersProperty struct {
	// Use the `RoleArn` structure to override an account-wide role for a specific Athena data source.
	//
	// For example, say an account administrator has turned off all Athena access with an account-wide role. The administrator can then use `RoleArn` to bypass the account-wide role and allow Athena access for the single Athena data source that is specified in the structure, even if the account-wide role forbidding Athena access is still active.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The workgroup that Amazon Athena uses.
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

