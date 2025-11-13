package interfacesawssecuritylake


// A reference to a DataLake resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakeReference := &DataLakeReference{
//   	DataLakeArn: jsii.String("dataLakeArn"),
//   }
//
type DataLakeReference struct {
	// The Arn of the DataLake resource.
	DataLakeArn *string `field:"required" json:"dataLakeArn" yaml:"dataLakeArn"`
}

