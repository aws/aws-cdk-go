package awsssm


// Information about the `AwsOrganizationsSource` resource data sync source.
//
// A sync source of this type can synchronize data from AWS Organizations or, if an AWS organization isn't present, from multiple AWS Regions .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsOrganizationsSourceProperty := &AwsOrganizationsSourceProperty{
//   	OrganizationSourceType: jsii.String("organizationSourceType"),
//
//   	// the properties below are optional
//   	OrganizationalUnits: []*string{
//   		jsii.String("organizationalUnits"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html
//
type CfnResourceDataSync_AwsOrganizationsSourceProperty struct {
	// If an AWS organization is present, this is either `OrganizationalUnits` or `EntireOrganization` .
	//
	// For `OrganizationalUnits` , the data is aggregated from a set of organization units. For `EntireOrganization` , the data is aggregated from the entire AWS organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationsourcetype
	//
	OrganizationSourceType *string `field:"required" json:"organizationSourceType" yaml:"organizationSourceType"`
	// The AWS Organizations organization units included in the sync.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-awsorganizationssource.html#cfn-ssm-resourcedatasync-awsorganizationssource-organizationalunits
	//
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
}

