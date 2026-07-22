package interfacesawsdataexchange


// A reference to a EntitledDataSets resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entitledDataSetsReference := &EntitledDataSetsReference{
//   	EntitledDataSetsArn: jsii.String("entitledDataSetsArn"),
//   	EntitledDataSetsId: jsii.String("entitledDataSetsId"),
//   }
//
type EntitledDataSetsReference struct {
	// The ARN of the EntitledDataSets resource.
	EntitledDataSetsArn *string `field:"required" json:"entitledDataSetsArn" yaml:"entitledDataSetsArn"`
	// The Id of the EntitledDataSets resource.
	EntitledDataSetsId *string `field:"required" json:"entitledDataSetsId" yaml:"entitledDataSetsId"`
}

