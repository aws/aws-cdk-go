package awskendra


// Specifies additional capacity units configured for your Enterprise Edition index.
//
// You can add and remove capacity units to fit your usage requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityUnitsConfigurationProperty := &capacityUnitsConfigurationProperty{
//   	queryCapacityUnits: jsii.Number(123),
//   	storageCapacityUnits: jsii.Number(123),
//   }
//
type CfnIndex_CapacityUnitsConfigurationProperty struct {
	// The amount of extra query capacity for an index and [GetQuerySuggestions](https://docs.aws.amazon.com/kendra/latest/dg/API_GetQuerySuggestions.html) capacity.
	//
	// A single extra capacity unit for an index provides 0.1 queries per second or approximately 8,000 queries per day.
	//
	// `GetQuerySuggestions` capacity is five times the provisioned query capacity for an index, or the base capacity of 2.5 calls per second, whichever is higher. For example, the base capacity for an index is 0.1 queries per second, and `GetQuerySuggestions` capacity has a base of 2.5 calls per second. If you add another 0.1 queries per second to total 0.2 queries per second for an index, the `GetQuerySuggestions` capacity is 2.5 calls per second (higher than five times 0.2 queries per second).
	QueryCapacityUnits *float64 `field:"required" json:"queryCapacityUnits" yaml:"queryCapacityUnits"`
	// The amount of extra storage capacity for an index.
	//
	// A single capacity unit provides 30 GB of storage space or 100,000 documents, whichever is reached first.
	StorageCapacityUnits *float64 `field:"required" json:"storageCapacityUnits" yaml:"storageCapacityUnits"`
}

