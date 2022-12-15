package awslicensemanager


// Properties for defining a `CfnGrant`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGrantProps := &cfnGrantProps{
//   	allowedOperations: []*string{
//   		jsii.String("allowedOperations"),
//   	},
//   	grantName: jsii.String("grantName"),
//   	homeRegion: jsii.String("homeRegion"),
//   	licenseArn: jsii.String("licenseArn"),
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	status: jsii.String("status"),
//   }
//
type CfnGrantProps struct {
	// Allowed operations for the grant.
	AllowedOperations *[]*string `field:"optional" json:"allowedOperations" yaml:"allowedOperations"`
	// Grant name.
	GrantName *string `field:"optional" json:"grantName" yaml:"grantName"`
	// Home Region of the grant.
	HomeRegion *string `field:"optional" json:"homeRegion" yaml:"homeRegion"`
	// License ARN.
	LicenseArn *string `field:"optional" json:"licenseArn" yaml:"licenseArn"`
	// The grant principals. You can specify one of the following as an Amazon Resource Name (ARN):.
	//
	// - An AWS account, which includes only the account specified.
	//
	// - An organizational unit (OU), which includes all accounts in the OU.
	//
	// - An organization, which will include all accounts across your organization.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Granted license status.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

