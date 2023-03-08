package awsssm


// Information about the source of the data included in the resource data sync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syncSourceProperty := &syncSourceProperty{
//   	sourceRegions: []*string{
//   		jsii.String("sourceRegions"),
//   	},
//   	sourceType: jsii.String("sourceType"),
//
//   	// the properties below are optional
//   	awsOrganizationsSource: &awsOrganizationsSourceProperty{
//   		organizationSourceType: jsii.String("organizationSourceType"),
//
//   		// the properties below are optional
//   		organizationalUnits: []*string{
//   			jsii.String("organizationalUnits"),
//   		},
//   	},
//   	includeFutureRegions: jsii.Boolean(false),
//   }
//
type CfnResourceDataSync_SyncSourceProperty struct {
	// The `SyncSource` AWS Regions included in the resource data sync.
	SourceRegions *[]*string `field:"required" json:"sourceRegions" yaml:"sourceRegions"`
	// The type of data source for the resource data sync.
	//
	// `SourceType` is either `AwsOrganizations` (if an organization is present in AWS Organizations ) or `SingleAccountMultiRegions` .
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Information about the AwsOrganizationsSource resource data sync source.
	//
	// A sync source of this type can synchronize data from AWS Organizations .
	AwsOrganizationsSource interface{} `field:"optional" json:"awsOrganizationsSource" yaml:"awsOrganizationsSource"`
	// Whether to automatically synchronize and aggregate data from new AWS Regions when those Regions come online.
	IncludeFutureRegions interface{} `field:"optional" json:"includeFutureRegions" yaml:"includeFutureRegions"`
}

