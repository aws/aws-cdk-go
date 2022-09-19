// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha


// The sort style of a table.
//
// Example:
//   awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &tableProps{
//   	tableColumns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			dataType: jsii.String("varchar(4)"),
//   			sortKey: jsii.Boolean(true),
//   		},
//   		&column{
//   			name: jsii.String("col2"),
//   			dataType: jsii.String("float"),
//   			sortKey: jsii.Boolean(true),
//   		},
//   	},
//   	cluster: cluster,
//   	databaseName: jsii.String("databaseName"),
//   	sortStyle: awscdkredshiftalpha.TableSortStyle_COMPOUND,
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

