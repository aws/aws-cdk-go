package interfacesawsconnect


// A reference to a DataLakeAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakeAssociationReference := &DataLakeAssociationReference{
//   	DataSetId: jsii.String("dataSetId"),
//   	InstanceId: jsii.String("instanceId"),
//   	TargetAccountId: jsii.String("targetAccountId"),
//   }
//
type DataLakeAssociationReference struct {
	// The DataSetId of the DataLakeAssociation resource.
	DataSetId *string `field:"required" json:"dataSetId" yaml:"dataSetId"`
	// The InstanceId of the DataLakeAssociation resource.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The TargetAccountId of the DataLakeAssociation resource.
	TargetAccountId *string `field:"required" json:"targetAccountId" yaml:"targetAccountId"`
}

