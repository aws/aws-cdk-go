package awscdkredshiftalpha


// The data distribution style of a table.
//
// Example:
//   awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
//   	TableColumns: []Column{
//   		&Column{
//   			Name: jsii.String("col1"),
//   			DataType: jsii.String("varchar(4)"),
//   			DistKey: jsii.Boolean(true),
//   		},
//   		&Column{
//   			Name: jsii.String("col2"),
//   			DataType: jsii.String("float"),
//   		},
//   	},
//   	Cluster: cluster,
//   	DatabaseName: jsii.String("databaseName"),
//   	DistStyle: awscdkredshiftalpha.TableDistStyle_KEY,
//   })
//
// Experimental.
type TableDistStyle string

const (
	// Amazon Redshift assigns an optimal distribution style based on the table data.
	// Experimental.
	TableDistStyle_AUTO TableDistStyle = "AUTO"
	// The data in the table is spread evenly across the nodes in a cluster in a round-robin distribution.
	// Experimental.
	TableDistStyle_EVEN TableDistStyle = "EVEN"
	// The data is distributed by the values in the DISTKEY column.
	// Experimental.
	TableDistStyle_KEY TableDistStyle = "KEY"
	// A copy of the entire table is distributed to every node.
	// Experimental.
	TableDistStyle_ALL TableDistStyle = "ALL"
)

