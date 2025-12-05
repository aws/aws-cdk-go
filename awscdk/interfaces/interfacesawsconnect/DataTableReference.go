package interfacesawsconnect


// A reference to a DataTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataTableReference := &DataTableReference{
//   	DataTableArn: jsii.String("dataTableArn"),
//   	InstanceArn: jsii.String("instanceArn"),
//   }
//
type DataTableReference struct {
	// The Arn of the DataTable resource.
	DataTableArn *string `field:"required" json:"dataTableArn" yaml:"dataTableArn"`
	// The InstanceArn of the DataTable resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
}

