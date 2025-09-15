package awsrds


// A reference to a DBSecurityGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBSecurityGroupReference := &DBSecurityGroupReference{
//   	DbSecurityGroupId: jsii.String("dbSecurityGroupId"),
//   }
//
type DBSecurityGroupReference struct {
	// The Id of the DBSecurityGroup resource.
	DbSecurityGroupId *string `field:"required" json:"dbSecurityGroupId" yaml:"dbSecurityGroupId"`
}

