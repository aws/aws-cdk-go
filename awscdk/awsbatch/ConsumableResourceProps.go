package awsbatch


// Properties for defining a Batch Consumable Resource.
//
// Example:
//   licenseResource := batch.NewConsumableResource(this, jsii.String("LicenseResource"), &ConsumableResourceProps{
//   	TotalQuantity: jsii.Number(100),
//   })
//
//   user := iam.NewUser(this, jsii.String("MyUser"))
//
//   // Describe the resource and list the jobs that use it
//   licenseResource.grants.Read(user)
//
//   // Change the total quantity, e.g. to replenish a NON_REPLENISHABLE pool
//   licenseResource.grants.Update(user)
//
type ConsumableResourceProps struct {
	// The total quantity of the consumable resource.
	//
	// Must be a non-negative integer.
	TotalQuantity *float64 `field:"required" json:"totalQuantity" yaml:"totalQuantity"`
	// The name of the consumable resource.
	//
	// The name must be unique within the account and Region. Changing it replaces
	// the resource.
	// Default: - CloudFormation-generated name.
	//
	ConsumableResourceName *string `field:"optional" json:"consumableResourceName" yaml:"consumableResourceName"`
	// The type of consumable resource.
	//
	// Changing this property replaces the resource. Because the name must be unique
	// within the account and Region, changing the type of a resource that has a
	// `consumableResourceName` fails the deployment unless the name is changed as well.
	// Default: ConsumableResourceType.REPLENISHABLE
	//
	ResourceType ConsumableResourceType `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Tags to apply to the consumable resource.
	// Default: - no tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

