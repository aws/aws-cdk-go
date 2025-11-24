package mixinsawsssm


// Information about the source of the data included in the resource data sync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   syncSourceProperty := &SyncSourceProperty{
//   	AwsOrganizationsSource: &AwsOrganizationsSourceProperty{
//   		OrganizationalUnits: []*string{
//   			jsii.String("organizationalUnits"),
//   		},
//   		OrganizationSourceType: jsii.String("organizationSourceType"),
//   	},
//   	IncludeFutureRegions: jsii.Boolean(false),
//   	SourceRegions: []*string{
//   		jsii.String("sourceRegions"),
//   	},
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html
//
type CfnResourceDataSyncPropsMixin_SyncSourceProperty struct {
	// Information about the AwsOrganizationsSource resource data sync source.
	//
	// A sync source of this type can synchronize data from AWS Organizations .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-awsorganizationssource
	//
	AwsOrganizationsSource interface{} `field:"optional" json:"awsOrganizationsSource" yaml:"awsOrganizationsSource"`
	// Whether to automatically synchronize and aggregate data from new AWS Regions when those Regions come online.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-includefutureregions
	//
	IncludeFutureRegions interface{} `field:"optional" json:"includeFutureRegions" yaml:"includeFutureRegions"`
	// The `SyncSource` AWS Regions included in the resource data sync.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourceregions
	//
	SourceRegions *[]*string `field:"optional" json:"sourceRegions" yaml:"sourceRegions"`
	// The type of data source for the resource data sync.
	//
	// `SourceType` is either `AwsOrganizations` (if an organization is present in AWS Organizations ) or `SingleAccountMultiRegions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-resourcedatasync-syncsource.html#cfn-ssm-resourcedatasync-syncsource-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

