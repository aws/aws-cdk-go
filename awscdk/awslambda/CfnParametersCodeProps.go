package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for `CfnParametersCode`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnParameter cfnParameter
//
//   cfnParametersCodeProps := &CfnParametersCodeProps{
//   	BucketNameParam: cfnParameter,
//   	ObjectKeyParam: cfnParameter,
//   }
//
type CfnParametersCodeProps struct {
	// The CloudFormation parameter that represents the name of the S3 Bucket where the Lambda code will be located in.
	//
	// Must be of type 'String'.
	// Default: a new parameter will be created.
	//
	BucketNameParam awscdk.CfnParameter `field:"optional" json:"bucketNameParam" yaml:"bucketNameParam"`
	// The CloudFormation parameter that represents the path inside the S3 Bucket where the Lambda code will be located at.
	//
	// Must be of type 'String'.
	// Default: a new parameter will be created.
	//
	ObjectKeyParam awscdk.CfnParameter `field:"optional" json:"objectKeyParam" yaml:"objectKeyParam"`
}

