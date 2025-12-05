package previewawsconnectmixins


// The lock version of the Data Table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lockVersionProperty := &LockVersionProperty{
//   	DataTable: jsii.String("dataTable"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatable-lockversion.html
//
type CfnDataTablePropsMixin_LockVersionProperty struct {
	// The data table for the lock version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatable-lockversion.html#cfn-connect-datatable-lockversion-datatable
	//
	DataTable *string `field:"optional" json:"dataTable" yaml:"dataTable"`
}

