package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   crossAccountRoleConfigurationProperty := &CrossAccountRoleConfigurationProperty{
//   	CrossAccountRoleArn: jsii.String("crossAccountRoleArn"),
//   	ExternalId: jsii.String("externalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-crossaccountroleconfiguration.html
//
type CfnServicePropsMixin_CrossAccountRoleConfigurationProperty struct {
	// ARN of the cross-account IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-crossaccountroleconfiguration.html#cfn-resiliencehubv2-service-crossaccountroleconfiguration-crossaccountrolearn
	//
	CrossAccountRoleArn *string `field:"optional" json:"crossAccountRoleArn" yaml:"crossAccountRoleArn"`
	// External ID for cross-account access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-crossaccountroleconfiguration.html#cfn-resiliencehubv2-service-crossaccountroleconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

