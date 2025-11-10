package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceRetentionPolicyProperty := &ResourceRetentionPolicyProperty{
//   	SkipDeregisterOnUpdate: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resourceretentionpolicy.html
//
type CfnJobDefinition_ResourceRetentionPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resourceretentionpolicy.html#cfn-batch-jobdefinition-resourceretentionpolicy-skipderegisteronupdate
	//
	// Default: - false.
	//
	SkipDeregisterOnUpdate interface{} `field:"optional" json:"skipDeregisterOnUpdate" yaml:"skipDeregisterOnUpdate"`
}

