package interfacesawsconnect


// A reference to a DataTableAttribute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataTableAttributeReference := &DataTableAttributeReference{
//   	AttributeId: jsii.String("attributeId"),
//   	DataTableArn: jsii.String("dataTableArn"),
//   	InstanceArn: jsii.String("instanceArn"),
//   }
//
type DataTableAttributeReference struct {
	// The AttributeId of the DataTableAttribute resource.
	AttributeId *string `field:"required" json:"attributeId" yaml:"attributeId"`
	// The DataTableArn of the DataTableAttribute resource.
	DataTableArn *string `field:"required" json:"dataTableArn" yaml:"dataTableArn"`
	// The InstanceArn of the DataTableAttribute resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
}

