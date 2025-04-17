package awscloudwatch


// Layout for TableWidget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewTableWidget(&TableWidgetProps{
//   	// ...
//
//   	Layout: cloudwatch.TableLayout_VERTICAL,
//   }))
//
type TableLayout string

const (
	// Data points are laid out in columns.
	TableLayout_HORIZONTAL TableLayout = "HORIZONTAL"
	// Data points are laid out in rows.
	TableLayout_VERTICAL TableLayout = "VERTICAL"
)

