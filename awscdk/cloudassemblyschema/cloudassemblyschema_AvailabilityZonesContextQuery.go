package cloudassemblyschema


// Query to availability zone context provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availabilityZonesContextQuery := &availabilityZonesContextQuery{
//   	account: jsii.String("account"),
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
type AvailabilityZonesContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

