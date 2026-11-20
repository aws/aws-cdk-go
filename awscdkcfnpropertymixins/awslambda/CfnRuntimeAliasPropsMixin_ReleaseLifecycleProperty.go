package awslambda


// The publisher's lifecycle declaration for the alias.
//
// Omitting it removes any stored declaration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   releaseLifecycleProperty := &ReleaseLifecycleProperty{
//   	DeprecationDate: jsii.String("deprecationDate"),
//   	FunctionCreateBlockDate: jsii.String("functionCreateBlockDate"),
//   	FunctionUpdateBlockDate: jsii.String("functionUpdateBlockDate"),
//   	ReleaseState: jsii.String("releaseState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-releaselifecycle.html
//
type CfnRuntimeAliasPropsMixin_ReleaseLifecycleProperty struct {
	// The date on which the alias becomes deprecated.
	//
	// A UTC calendar date in yyyy-mm-dd format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-releaselifecycle.html#cfn-lambda-runtimealias-releaselifecycle-deprecationdate
	//
	DeprecationDate *string `field:"optional" json:"deprecationDate" yaml:"deprecationDate"`
	// The date on which creating new functions on this alias is blocked.
	//
	// A UTC calendar date in yyyy-mm-dd format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-releaselifecycle.html#cfn-lambda-runtimealias-releaselifecycle-functioncreateblockdate
	//
	FunctionCreateBlockDate *string `field:"optional" json:"functionCreateBlockDate" yaml:"functionCreateBlockDate"`
	// The date on which updating existing functions on this alias is blocked.
	//
	// A UTC calendar date in yyyy-mm-dd format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-releaselifecycle.html#cfn-lambda-runtimealias-releaselifecycle-functionupdateblockdate
	//
	FunctionUpdateBlockDate *string `field:"optional" json:"functionUpdateBlockDate" yaml:"functionUpdateBlockDate"`
	// The release state of the alias.
	//
	// DEPRECATED requires a DeprecationDate on or before today (UTC); PREVIEW and GA require a later DeprecationDate, or none. Reads report DEPRECATED once DeprecationDate is reached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimealias-releaselifecycle.html#cfn-lambda-runtimealias-releaselifecycle-releasestate
	//
	ReleaseState *string `field:"optional" json:"releaseState" yaml:"releaseState"`
}

