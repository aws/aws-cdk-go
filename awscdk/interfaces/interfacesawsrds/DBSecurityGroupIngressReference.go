package interfacesawsrds


// A reference to a DBSecurityGroupIngress resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBSecurityGroupIngressReference := &DBSecurityGroupIngressReference{
//   	DbSecurityGroupIngressId: jsii.String("dbSecurityGroupIngressId"),
//   }
//
type DBSecurityGroupIngressReference struct {
	// The Id of the DBSecurityGroupIngress resource.
	DbSecurityGroupIngressId *string `field:"required" json:"dbSecurityGroupIngressId" yaml:"dbSecurityGroupIngressId"`
}

