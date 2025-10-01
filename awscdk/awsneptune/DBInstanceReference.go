package awsneptune


// A reference to a DBInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBInstanceReference := &DBInstanceReference{
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   }
//
type DBInstanceReference struct {
	// The DBInstanceIdentifier of the DBInstance resource.
	DbInstanceIdentifier *string `field:"required" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
}

