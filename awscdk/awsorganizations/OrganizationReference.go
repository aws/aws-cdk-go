package awsorganizations


// A reference to a Organization resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationReference := &OrganizationReference{
//   	OrganizationArn: jsii.String("organizationArn"),
//   	OrganizationId: jsii.String("organizationId"),
//   }
//
type OrganizationReference struct {
	// The ARN of the Organization resource.
	OrganizationArn *string `field:"required" json:"organizationArn" yaml:"organizationArn"`
	// The Id of the Organization resource.
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
}

