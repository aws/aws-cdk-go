// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   approveStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Approve"),
//   })
//   manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&manualApprovalActionProps{
//   	actionName: jsii.String("Approve"),
//   })
//   approveStage.addAction(manualApprovalAction)
//
//   role := iam.role.fromRoleArn(this, jsii.String("Admin"), awscdk.Arn.format(&arnComponents{
//   	service: jsii.String("iam"),
//   	resource: jsii.String("role"),
//   	resourceName: jsii.String("Admin"),
//   }, this))
//   manualApprovalAction.grantManualApproval(role)
//
// Experimental.
type ArnComponents struct {
	// Resource type (e.g. "table", "autoScalingGroup", "certificate"). For some resource types, e.g. S3 buckets, this field defines the bucket name.
	// Experimental.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The service namespace that identifies the AWS product (for example, 's3', 'iam', 'codepipline').
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The ID of the AWS account that owns the resource, without the hyphens.
	//
	// For example, 123456789012. Note that the ARNs for some resources don't
	// require an account number, so this component might be omitted.
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The specific ARN format to use for this ARN value.
	// Experimental.
	ArnFormat ArnFormat `field:"optional" json:"arnFormat" yaml:"arnFormat"`
	// The partition that the resource is in.
	//
	// For standard AWS regions, the
	// partition is aws. If you have resources in other partitions, the
	// partition is aws-partitionname. For example, the partition for resources
	// in the China (Beijing) region is aws-cn.
	// Experimental.
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The region the resource resides in.
	//
	// Note that the ARNs for some resources
	// do not require a region, so this component might be omitted.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Resource name or path within the resource (i.e. S3 bucket object key) or a wildcard such as ``"*"``. This is service-dependent.
	// Experimental.
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// Separator between resource type and the resource.
	//
	// Can be either '/', ':' or an empty string. Will only be used if resourceName is defined.
	// Deprecated: use arnFormat instead.
	Sep *string `field:"optional" json:"sep" yaml:"sep"`
}

