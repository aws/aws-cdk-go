package awssagemaker


// The team this compute allocation rule applies to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeQuotaTargetProperty := &ComputeQuotaTargetProperty{
//   	TeamName: jsii.String("teamName"),
//
//   	// the properties below are optional
//   	FairShareWeight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotatarget.html
//
type CfnComputeQuota_ComputeQuotaTargetProperty struct {
	// The name of the team.
	//
	// Materialized as Kubernetes namespace 'hyperpod-ns-<TeamName>' on the EKS cluster, so the value must satisfy RFC 1123 namespace constraints. Reassignment is not supported (createOnly).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotatarget.html#cfn-sagemaker-computequota-computequotatarget-teamname
	//
	TeamName *string `field:"required" json:"teamName" yaml:"teamName"`
	// Relative fair-share weight when fair-share scheduling is enabled on the parent ClusterSchedulerConfig.
	//
	// Defaults service-side to 0 when omitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-computequotatarget.html#cfn-sagemaker-computequota-computequotatarget-fairshareweight
	//
	FairShareWeight *float64 `field:"optional" json:"fairShareWeight" yaml:"fairShareWeight"`
}

