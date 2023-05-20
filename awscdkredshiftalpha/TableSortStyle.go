package awscdkredshiftalpha


// The sort style of a table.
//
// Example:
//   awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
//   	TableColumns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			DataType: jsii.String("varchar(4)"),
//   			SortKey: jsii.Boolean(true),
//   		},
//   		&column{
//   			Name: jsii.String("col2"),
//   			DataType: jsii.String("float"),
//   			SortKey: jsii.Boolean(true),
//   		},
//   	},
//   	Cluster: cluster,
//   	DatabaseName: jsii.String("databaseName"),
//   	SortStyle: awscdkredshiftalpha.TableSortStyle_COMPOUND,
//   })
//
// Experimental.
type TableSortStyle string

const (
	// Amazon Redshift assigns an optimal sort key based on the table data.
	// Experimental.
	TableSortStyle_AUTO TableSortStyle = "AUTO"
	// Specifies that the data is sorted using a compound key made up of all of the listed columns, in the order they are listed.
	// Experimental.
	TableSortStyle_COMPOUND TableSortStyle = "COMPOUND"
	// Specifies that the data is sorted using an interleaved sort key.
	// Experimental.
	TableSortStyle_INTERLEAVED TableSortStyle = "INTERLEAVED"
)

