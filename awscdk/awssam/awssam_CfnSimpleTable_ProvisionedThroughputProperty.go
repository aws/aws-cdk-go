package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedThroughputProperty := &provisionedThroughputProperty{
//   	writeCapacityUnits: jsii.Number(123),
//
//   	// the properties below are optional
//   	readCapacityUnits: jsii.Number(123),
//   }
//
type CfnSimpleTable_ProvisionedThroughputProperty struct {
	// `CfnSimpleTable.ProvisionedThroughputProperty.WriteCapacityUnits`.
	WriteCapacityUnits *float64 `field:"required" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
	// `CfnSimpleTable.ProvisionedThroughputProperty.ReadCapacityUnits`.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

