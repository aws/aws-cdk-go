package awsservicecatalog


// The user-defined preferences that will be applied when updating a provisioned product.
//
// Not all preferences are applicable to all provisioned product type
//
// One or more AWS accounts that will have access to the provisioned product.
//
// Applicable only to a `CFN_STACKSET` provisioned product type.
//
// The AWS accounts specified should be within the list of accounts in the `STACKSET` constraint. To get the list of accounts in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
//
// If no values are specified, the default value is all accounts from the `STACKSET` constraint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisioningPreferencesProperty := &provisioningPreferencesProperty{
//   	stackSetAccounts: []*string{
//   		jsii.String("stackSetAccounts"),
//   	},
//   	stackSetFailureToleranceCount: jsii.Number(123),
//   	stackSetFailureTolerancePercentage: jsii.Number(123),
//   	stackSetMaxConcurrencyCount: jsii.Number(123),
//   	stackSetMaxConcurrencyPercentage: jsii.Number(123),
//   	stackSetOperationType: jsii.String("stackSetOperationType"),
//   	stackSetRegions: []*string{
//   		jsii.String("stackSetRegions"),
//   	},
//   }
//
type CfnCloudFormationProvisionedProduct_ProvisioningPreferencesProperty struct {
	// One or more AWS accounts where the provisioned product will be available.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// The specified accounts should be within the list of accounts from the `STACKSET` constraint. To get the list of accounts in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
	//
	// If no values are specified, the default value is all acounts from the `STACKSET` constraint.
	StackSetAccounts *[]*string `field:"optional" json:"stackSetAccounts" yaml:"stackSetAccounts"`
	// The number of accounts, per Region, for which this operation can fail before AWS Service Catalog stops the operation in that Region.
	//
	// If the operation is stopped in a Region, AWS Service Catalog doesn't attempt the operation in any subsequent Regions.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// Conditional: You must specify either `StackSetFailureToleranceCount` or `StackSetFailureTolerancePercentage` , but not both.
	//
	// The default value is `0` if no value is specified.
	StackSetFailureToleranceCount *float64 `field:"optional" json:"stackSetFailureToleranceCount" yaml:"stackSetFailureToleranceCount"`
	// The percentage of accounts, per Region, for which this stack operation can fail before AWS Service Catalog stops the operation in that Region.
	//
	// If the operation is stopped in a Region, AWS Service Catalog doesn't attempt the operation in any subsequent Regions.
	//
	// When calculating the number of accounts based on the specified percentage, AWS Service Catalog rounds down to the next whole number.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// Conditional: You must specify either `StackSetFailureToleranceCount` or `StackSetFailureTolerancePercentage` , but not both.
	StackSetFailureTolerancePercentage *float64 `field:"optional" json:"stackSetFailureTolerancePercentage" yaml:"stackSetFailureTolerancePercentage"`
	// The maximum number of accounts in which to perform this operation at one time.
	//
	// This is dependent on the value of `StackSetFailureToleranceCount` . `StackSetMaxConcurrentCount` is at most one more than the `StackSetFailureToleranceCount` .
	//
	// Note that this setting lets you specify the maximum for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// Conditional: You must specify either `StackSetMaxConcurrentCount` or `StackSetMaxConcurrentPercentage` , but not both.
	StackSetMaxConcurrencyCount *float64 `field:"optional" json:"stackSetMaxConcurrencyCount" yaml:"stackSetMaxConcurrencyCount"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified percentage, AWS Service Catalog rounds down to the next whole number. This is true except in cases where rounding down would result is zero. In this case, AWS Service Catalog sets the number as `1` instead.
	//
	// Note that this setting lets you specify the maximum for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// Conditional: You must specify either `StackSetMaxConcurrentCount` or `StackSetMaxConcurrentPercentage` , but not both.
	StackSetMaxConcurrencyPercentage *float64 `field:"optional" json:"stackSetMaxConcurrencyPercentage" yaml:"stackSetMaxConcurrencyPercentage"`
	// Determines what action AWS Service Catalog performs to a stack set or a stack instance represented by the provisioned product.
	//
	// The default value is `UPDATE` if nothing is specified.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// - **CREATE** - Creates a new stack instance in the stack set represented by the provisioned product. In this case, only new stack instances are created based on accounts and Regions; if new ProductId or ProvisioningArtifactID are passed, they will be ignored.
	// - **UPDATE** - Updates the stack set represented by the provisioned product and also its stack instances.
	// - **DELETE** - Deletes a stack instance in the stack set represented by the provisioned product.
	StackSetOperationType *string `field:"optional" json:"stackSetOperationType" yaml:"stackSetOperationType"`
	// One or more AWS Regions where the provisioned product will be available.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// The specified Regions should be within the list of Regions from the `STACKSET` constraint. To get the list of Regions in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
	//
	// If no values are specified, the default value is all Regions from the `STACKSET` constraint.
	StackSetRegions *[]*string `field:"optional" json:"stackSetRegions" yaml:"stackSetRegions"`
}

