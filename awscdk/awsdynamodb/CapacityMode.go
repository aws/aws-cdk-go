package awsdynamodb


// Capacity modes.
type CapacityMode string

const (
	// Fixed.
	CapacityMode_FIXED CapacityMode = "FIXED"
	// Autoscaled.
	CapacityMode_AUTOSCALED CapacityMode = "AUTOSCALED"
)

