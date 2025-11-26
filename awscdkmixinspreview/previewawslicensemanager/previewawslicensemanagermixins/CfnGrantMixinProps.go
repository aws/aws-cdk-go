package previewawslicensemanagermixins


// Properties for CfnGrantPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGrantMixinProps := &CfnGrantMixinProps{
//   	AllowedOperations: []*string{
//   		jsii.String("allowedOperations"),
//   	},
//   	GrantName: jsii.String("grantName"),
//   	HomeRegion: jsii.String("homeRegion"),
//   	LicenseArn: jsii.String("licenseArn"),
//   	Principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html
//
type CfnGrantMixinProps struct {
	// Allowed operations for the grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-allowedoperations
	//
	AllowedOperations *[]*string `field:"optional" json:"allowedOperations" yaml:"allowedOperations"`
	// Grant name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantname
	//
	GrantName *string `field:"optional" json:"grantName" yaml:"grantName"`
	// Home Region of the grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-homeregion
	//
	HomeRegion *string `field:"optional" json:"homeRegion" yaml:"homeRegion"`
	// License ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-licensearn
	//
	LicenseArn *string `field:"optional" json:"licenseArn" yaml:"licenseArn"`
	// The grant principals. You can specify one of the following as an Amazon Resource Name (ARN):.
	//
	// - An AWS account, which includes only the account specified.
	//
	// - An organizational unit (OU), which includes all accounts in the OU.
	//
	// - An organization, which will include all accounts across your organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-principals
	//
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Granted license status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

