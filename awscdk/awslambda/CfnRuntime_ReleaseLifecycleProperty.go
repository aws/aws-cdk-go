package awslambda


// The publisher's lifecycle declaration for the runtime.
//
// Omitting it removes any stored declaration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   releaseLifecycleProperty := &ReleaseLifecycleProperty{
//   	DeprecationDate: jsii.String("deprecationDate"),
//   	FunctionCreateBlockDate: jsii.String("functionCreateBlockDate"),
//   	FunctionUpdateBlockDate: jsii.String("functionUpdateBlockDate"),
//   	ReleaseState: jsii.String("releaseState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtime-releaselifecycle.html
//
type CfnRuntime_ReleaseLifecycleProperty struct {
	// The date on which the runtime becomes deprecated.
	//
	// A UTC calendar date in yyyy-mm-dd format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtime-releaselifecycle.html#cfn-lambda-runtime-releaselifecycle-deprecationdate
	//
	DeprecationDate *string `field:"optional" json:"deprecationDate" yaml:"deprecationDate"`
	// The date on which creating new functions on this runtime is blocked.
	//
	// A UTC calendar date in yyyy-mm-dd format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtime-releaselifecycle.html#cfn-lambda-runtime-releaselifecycle-functioncreateblockdate
	//
	FunctionCreateBlockDate *string `field:"optional" json:"functionCreateBlockDate" yaml:"functionCreateBlockDate"`
	// The date on which updating existing functions on this runtime is blocked.
	//
	// A UTC calendar date in yyyy-mm-dd format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtime-releaselifecycle.html#cfn-lambda-runtime-releaselifecycle-functionupdateblockdate
	//
	FunctionUpdateBlockDate *string `field:"optional" json:"functionUpdateBlockDate" yaml:"functionUpdateBlockDate"`
	// The release state of the runtime.
	//
	// DEPRECATED requires a DeprecationDate on or before today (UTC); PREVIEW and GA require a later DeprecationDate, or none. Reads report DEPRECATED once DeprecationDate is reached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtime-releaselifecycle.html#cfn-lambda-runtime-releaselifecycle-releasestate
	//
	ReleaseState *string `field:"optional" json:"releaseState" yaml:"releaseState"`
}

