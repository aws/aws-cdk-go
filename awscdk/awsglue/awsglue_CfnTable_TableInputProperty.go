package awsglue


// A structure used to define a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   tableInputProperty := &tableInputProperty{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	owner: jsii.String("owner"),
//   	parameters: parameters,
//   	partitionKeys: []interface{}{
//   		&columnProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			comment: jsii.String("comment"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	retention: jsii.Number(123),
//   	storageDescriptor: &storageDescriptorProperty{
//   		bucketColumns: []*string{
//   			jsii.String("bucketColumns"),
//   		},
//   		columns: []interface{}{
//   			&columnProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				comment: jsii.String("comment"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		compressed: jsii.Boolean(false),
//   		inputFormat: jsii.String("inputFormat"),
//   		location: jsii.String("location"),
//   		numberOfBuckets: jsii.Number(123),
//   		outputFormat: jsii.String("outputFormat"),
//   		parameters: parameters,
//   		schemaReference: &schemaReferenceProperty{
//   			schemaId: &schemaIdProperty{
//   				registryName: jsii.String("registryName"),
//   				schemaArn: jsii.String("schemaArn"),
//   				schemaName: jsii.String("schemaName"),
//   			},
//   			schemaVersionId: jsii.String("schemaVersionId"),
//   			schemaVersionNumber: jsii.Number(123),
//   		},
//   		serdeInfo: &serdeInfoProperty{
//   			name: jsii.String("name"),
//   			parameters: parameters,
//   			serializationLibrary: jsii.String("serializationLibrary"),
//   		},
//   		skewedInfo: &skewedInfoProperty{
//   			skewedColumnNames: []*string{
//   				jsii.String("skewedColumnNames"),
//   			},
//   			skewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   			skewedColumnValues: []*string{
//   				jsii.String("skewedColumnValues"),
//   			},
//   		},
//   		sortColumns: []interface{}{
//   			&orderProperty{
//   				column: jsii.String("column"),
//   				sortOrder: jsii.Number(123),
//   			},
//   		},
//   		storedAsSubDirectories: jsii.Boolean(false),
//   	},
//   	tableType: jsii.String("tableType"),
//   	targetTable: &tableIdentifierProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   	},
//   	viewExpandedText: jsii.String("viewExpandedText"),
//   	viewOriginalText: jsii.String("viewOriginalText"),
//   }
//
type CfnTable_TableInputProperty struct {
	// A description of the table.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The table name.
	//
	// For Hive compatibility, this is folded to lowercase when it is stored.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The table owner.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// These key-value pairs define properties associated with the table.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A list of columns by which the table is partitioned. Only primitive types are supported as partition keys.
	//
	// When you create a table used by Amazon Athena, and you do not specify any `partitionKeys` , you must at least set the value of `partitionKeys` to an empty list. For example:
	//
	// `"PartitionKeys": []`.
	PartitionKeys interface{} `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// The retention time for this table.
	Retention *float64 `field:"optional" json:"retention" yaml:"retention"`
	// A storage descriptor containing information about the physical storage of this table.
	StorageDescriptor interface{} `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
	// The type of this table ( `EXTERNAL_TABLE` , `VIRTUAL_VIEW` , etc.).
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
	// A `TableIdentifier` structure that describes a target table for resource linking.
	TargetTable interface{} `field:"optional" json:"targetTable" yaml:"targetTable"`
	// If the table is a view, the expanded text of the view;
	//
	// otherwise `null` .
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// If the table is a view, the original text of the view;
	//
	// otherwise `null` .
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

