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
// Experimental.
type SchemaOptions struct {
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Experimental.
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
}

