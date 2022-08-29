package awsdynamodb


// Represents the table schema attributes.
//
// Example:
//   var table table
//
//   schema := table.schema()
//   partitionKey := schema.partitionKey
//   sortKey := schema.sortKey
//
type SchemaOptions struct {
	// Partition key attribute definition.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
}

