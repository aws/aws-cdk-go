package awsredshift


// A list of scopes set up for Lake Formation integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lakeFormationScopeUnionProperty := &LakeFormationScopeUnionProperty{
//   	LakeFormationQuery: &LakeFormationQueryProperty{
//   		Authorization: jsii.String("authorization"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-lakeformationscopeunion.html
//
type CfnRedshiftIdcApplication_LakeFormationScopeUnionProperty struct {
	// The Lake Formation scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-redshiftidcapplication-lakeformationscopeunion.html#cfn-redshift-redshiftidcapplication-lakeformationscopeunion-lakeformationquery
	//
	LakeFormationQuery interface{} `field:"optional" json:"lakeFormationQuery" yaml:"lakeFormationQuery"`
}

