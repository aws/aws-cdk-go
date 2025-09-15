package awsrds


// A reference to a DBInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBInstanceReference := &DBInstanceReference{
//   	DbInstanceArn: jsii.String("dbInstanceArn"),
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   }
//
type DBInstanceReference struct {
	// The ARN of the DBInstance resource.
	DbInstanceArn *string `field:"required" json:"dbInstanceArn" yaml:"dbInstanceArn"`
	// The DBInstanceIdentifier of the DBInstance resource.
	DbInstanceIdentifier *string `field:"required" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
}

