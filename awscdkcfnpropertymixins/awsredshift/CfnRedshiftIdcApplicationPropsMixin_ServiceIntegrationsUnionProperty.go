package awsredshift


// A list of service integrations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   serviceIntegrationsUnionProperty := &ServiceIntegrationsUnionProperty{
//   	LakeFormation: []interface{}{
//   		&LakeFormationScopeUnionProperty{
//   			LakeFormationQuery: &LakeFormationQueryProperty{
//   				Authorization: jsii.String("authorization"),
//   			},
//   		},
//   	},
//   	Redshift: []interface{}{
//   		&RedshiftScopeUnionProperty{
//   			Connect: &ConnectProperty{
//   				Authorization: jsii.String("authorization"),
//   			},
//   		},
//   	},
//   	S3AccessGrants: []interface{}{
//   		&S3AccessGrantsScopeUnionProperty{
//   			ReadWriteAccess: &ReadWriteAccessProperty{
//   				Authorization: jsii.String("authorization"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-serviceintegrationsunion.html
//
type CfnRedshiftIdcApplicationPropsMixin_ServiceIntegrationsUnionProperty struct {
	// A list of scopes set up for Lake Formation integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-serviceintegrationsunion.html#cfn-redshift-redshiftidcapplication-serviceintegrationsunion-lakeformation
	//
	LakeFormation interface{} `field:"optional" json:"lakeFormation" yaml:"lakeFormation"`
	// A list of scopes set up for Amazon Redshift integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-serviceintegrationsunion.html#cfn-redshift-redshiftidcapplication-serviceintegrationsunion-redshift
	//
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// A list of scopes set up for S3 Access Grants integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-serviceintegrationsunion.html#cfn-redshift-redshiftidcapplication-serviceintegrationsunion-s3accessgrants
	//
	S3AccessGrants interface{} `field:"optional" json:"s3AccessGrants" yaml:"s3AccessGrants"`
}

