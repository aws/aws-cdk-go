package awsses


// The reputation configuration for the tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   reputationEntityProperty := &ReputationEntityProperty{
//   	Policy: &ReputationPolicyProperty{
//   		ReputationEntityPolicy: jsii.String("reputationEntityPolicy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-reputationentity.html
//
type CfnTenantPropsMixin_ReputationEntityProperty struct {
	// The reputation policy for the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-reputationentity.html#cfn-ses-tenant-reputationentity-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

