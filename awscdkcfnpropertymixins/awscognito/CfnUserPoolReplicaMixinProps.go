package awscognito


// Properties for CfnUserPoolReplicaPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnUserPoolReplicaMixinProps := &CfnUserPoolReplicaMixinProps{
//   	RegionName: jsii.String("regionName"),
//   	UserPoolId: jsii.String("userPoolId"),
//   	UserPoolTagsAtCreate: map[string]*string{
//   		"userPoolTagsAtCreateKey": jsii.String("userPoolTagsAtCreate"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolreplica.html
//
type CfnUserPoolReplicaMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolreplica.html#cfn-cognito-userpoolreplica-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolreplica.html#cfn-cognito-userpoolreplica-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolreplica.html#cfn-cognito-userpoolreplica-userpooltagsatcreate
	//
	UserPoolTagsAtCreate *map[string]*string `field:"optional" json:"userPoolTagsAtCreate" yaml:"userPoolTagsAtCreate"`
}

