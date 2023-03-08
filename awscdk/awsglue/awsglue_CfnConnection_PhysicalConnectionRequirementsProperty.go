package awsglue


// Specifies the physical requirements for a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   physicalConnectionRequirementsProperty := &physicalConnectionRequirementsProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	securityGroupIdList: []*string{
//   		jsii.String("securityGroupIdList"),
//   	},
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnConnection_PhysicalConnectionRequirementsProperty struct {
	// The connection's Availability Zone.
	//
	// This field is redundant because the specified subnet implies the Availability Zone to be used. Currently the field must be populated, but it will be deprecated in the future.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The security group ID list used by the connection.
	SecurityGroupIdList *[]*string `field:"optional" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// The subnet ID used by the connection.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

