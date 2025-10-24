package awsosis


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   resourcePolicyProperty := &ResourcePolicyProperty{
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-resourcepolicy.html
//
type CfnPipeline_ResourcePolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-resourcepolicy.html#cfn-osis-pipeline-resourcepolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
}

