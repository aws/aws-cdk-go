package awscdk


// Example:
//   subZone := route53.NewPublicHostedZone(this, jsii.String("SubZone"), &PublicHostedZoneProps{
//   	ZoneName: jsii.String("sub.someexample.com"),
//   })
//
//   // import the delegation role by constructing the roleArn
//   delegationRoleArn := awscdk.stack_Of(this).FormatArn(&ArnComponents{
//   	Region: jsii.String(""),
//   	 // IAM is global in each partition
//   	Service: jsii.String("iam"),
//   	Account: jsii.String("parent-account-id"),
//   	Resource: jsii.String("role"),
//   	ResourceName: jsii.String("MyDelegationRole"),
//   })
//   delegationRole := iam.Role_FromRoleArn(this, jsii.String("DelegationRole"), delegationRoleArn)
//
//   // create the record
//   // create the record
//   route53.NewCrossAccountZoneDelegationRecord(this, jsii.String("delegate"), &CrossAccountZoneDelegationRecordProps{
//   	DelegatedZone: subZone,
//   	ParentHostedZoneName: jsii.String("someexample.com"),
//   	 // or you can use parentHostedZoneId
//   	DelegationRole: DelegationRole,
//   })
//
type ArnComponents struct {
	// Resource type (e.g. "table", "autoScalingGroup", "certificate"). For some resource types, e.g. S3 buckets, this field defines the bucket name.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The service namespace that identifies the AWS product (for example, 's3', 'iam', 'codepipline').
	Service *string `field:"required" json:"service" yaml:"service"`
	// The ID of the AWS account that owns the resource, without the hyphens.
	//
	// For example, 123456789012. Note that the ARNs for some resources don't
	// require an account number, so this component might be omitted.
	// Default: The account the stack is deployed to.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The specific ARN format to use for this ARN value.
	// Default: - uses value of `sep` as the separator for formatting,
	// `ArnFormat.SLASH_RESOURCE_NAME` if that property was also not provided
	//
	ArnFormat ArnFormat `field:"optional" json:"arnFormat" yaml:"arnFormat"`
	// The partition that the resource is in.
	//
	// For standard AWS regions, the
	// partition is aws. If you have resources in other partitions, the
	// partition is aws-partitionname. For example, the partition for resources
	// in the China (Beijing) region is aws-cn.
	// Default: The AWS partition the stack is deployed to.
	//
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// The region the resource resides in.
	//
	// Note that the ARNs for some resources
	// do not require a region, so this component might be omitted.
	// Default: The region the stack is deployed to.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Resource name or path within the resource (i.e. S3 bucket object key) or a wildcard such as ``"*"``. This is service-dependent.
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
}

