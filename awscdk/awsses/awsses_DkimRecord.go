package awsses


// A DKIM record.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dkimRecord := &dkimRecord{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type DkimRecord struct {
	// The name of the record.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the record.
	Value *string `field:"required" json:"value" yaml:"value"`
}

