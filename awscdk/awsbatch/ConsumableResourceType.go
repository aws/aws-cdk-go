package awsbatch


// The type of consumable resource.
//
// Example:
//   // Create a replenishable consumable resource (e.g., software licenses)
//   licenseResource := batch.NewConsumableResource(this, jsii.String("LicenseResource"), &ConsumableResourceProps{
//   	ConsumableResourceName: jsii.String("my-software-license"),
//   	ResourceType: batch.ConsumableResourceType_REPLENISHABLE,
//   	TotalQuantity: jsii.Number(100),
//   })
//
//   // Create a non-replenishable consumable resource (e.g., API rate limits)
//   apiRateLimitResource := batch.NewConsumableResource(this, jsii.String("ApiRateLimitResource"), &ConsumableResourceProps{
//   	ConsumableResourceName: jsii.String("api-rate-limit"),
//   	ResourceType: batch.ConsumableResourceType_NON_REPLENISHABLE,
//   	TotalQuantity: jsii.Number(1000),
//   })
//
type ConsumableResourceType string

const (
	// Resource can be re-used after a job completes.
	ConsumableResourceType_REPLENISHABLE ConsumableResourceType = "REPLENISHABLE"
	// Resource cannot be re-used after a job completes.
	ConsumableResourceType_NON_REPLENISHABLE ConsumableResourceType = "NON_REPLENISHABLE"
)

