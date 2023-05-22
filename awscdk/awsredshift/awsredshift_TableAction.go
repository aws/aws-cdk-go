package awsredshift


// An action that a Redshift user can be granted privilege to perform on a table.
//
// Example:
//   databaseName := "databaseName"
//   username := "myuser"
//   tableName := "mytable"
//
//   user := awscdk.NewUser(this, jsii.String("User"), &UserProps{
//   	Username: username,
//   	Cluster: cluster,
//   	DatabaseName: databaseName,
//   })
//   table := awscdk.NewTable(this, jsii.String("Table"), &TableProps{
//   	TableColumns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			DataType: jsii.String("varchar(4)"),
//   		},
//   		&column{
//   			Name: jsii.String("col2"),
//   			DataType: jsii.String("float"),
//   		},
//   	},
//   	Cluster: cluster,
//   	DatabaseName: databaseName,
//   })
//   table.grant(user, awscdk.TableAction_INSERT)
//
// Experimental.
type TableAction string

const (
	// Grants privilege to select data from a table or view using a SELECT statement.
	// Experimental.
	TableAction_SELECT TableAction = "SELECT"
	// Grants privilege to load data into a table using an INSERT statement or a COPY statement.
	// Experimental.
	TableAction_INSERT TableAction = "INSERT"
	// Grants privilege to update a table column using an UPDATE statement.
	// Experimental.
	TableAction_UPDATE TableAction = "UPDATE"
	// Grants privilege to delete a data row from a table.
	// Experimental.
	TableAction_DELETE TableAction = "DELETE"
	// Grants privilege to drop a table.
	// Experimental.
	TableAction_DROP TableAction = "DROP"
	// Grants privilege to create a foreign key constraint.
	//
	// You need to grant this privilege on both the referenced table and the referencing table; otherwise, the user can't create the constraint.
	// Experimental.
	TableAction_REFERENCES TableAction = "REFERENCES"
	// Grants all available privileges at once to the specified user or user group.
	// Experimental.
	TableAction_ALL TableAction = "ALL"
)

