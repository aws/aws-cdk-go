package interfacesawsconnect


// A reference to a DataTableRecord resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataTableRecordReference := &DataTableRecordReference{
//   	DataTableArn: jsii.String("dataTableArn"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	RecordId: jsii.String("recordId"),
//   }
//
type DataTableRecordReference struct {
	// The DataTableArn of the DataTableRecord resource.
	DataTableArn *string `field:"required" json:"dataTableArn" yaml:"dataTableArn"`
	// The InstanceArn of the DataTableRecord resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The RecordId of the DataTableRecord resource.
	RecordId *string `field:"required" json:"recordId" yaml:"recordId"`
}

