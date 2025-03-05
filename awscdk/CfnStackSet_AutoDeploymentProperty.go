package awscdk


// [ `Service-managed` permissions] Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to a target organizational unit (OU).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   autoDeploymentProperty := &AutoDeploymentProperty{
//   	Enabled: jsii.Boolean(false),
//   	RetainStacksOnAccountRemoval: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-autodeployment.html
//
type CfnStackSet_AutoDeploymentProperty struct {
	// If set to `true` , StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	//
	// If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-autodeployment.html#cfn-cloudformation-stackset-autodeployment-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If set to `true` , stack resources are retained when an account is removed from a target organization or OU.
	//
	// If set to `false` , stack resources are deleted. Specify only if `Enabled` is set to `True` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-autodeployment.html#cfn-cloudformation-stackset-autodeployment-retainstacksonaccountremoval
	//
	RetainStacksOnAccountRemoval interface{} `field:"optional" json:"retainStacksOnAccountRemoval" yaml:"retainStacksOnAccountRemoval"`
}

