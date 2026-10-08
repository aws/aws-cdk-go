package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRedshiftIdcApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRedshiftIdcApplicationMixinProps := &CfnRedshiftIdcApplicationMixinProps{
//   	ApplicationType: jsii.String("applicationType"),
//   	AuthorizedTokenIssuerList: []interface{}{
//   		&AuthorizedTokenIssuerProperty{
//   			AuthorizedAudiencesList: []*string{
//   				jsii.String("authorizedAudiencesList"),
//   			},
//   			TrustedTokenIssuerArn: jsii.String("trustedTokenIssuerArn"),
//   		},
//   	},
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	IdcDisplayName: jsii.String("idcDisplayName"),
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	IdentityNamespace: jsii.String("identityNamespace"),
//   	RedshiftIdcApplicationName: jsii.String("redshiftIdcApplicationName"),
//   	ServiceIntegrations: []interface{}{
//   		&ServiceIntegrationsUnionProperty{
//   			LakeFormation: []interface{}{
//   				&LakeFormationScopeUnionProperty{
//   					LakeFormationQuery: &LakeFormationQueryProperty{
//   						Authorization: jsii.String("authorization"),
//   					},
//   				},
//   			},
//   			Redshift: []interface{}{
//   				&RedshiftScopeUnionProperty{
//   					Connect: &ConnectProperty{
//   						Authorization: jsii.String("authorization"),
//   					},
//   				},
//   			},
//   			S3AccessGrants: []interface{}{
//   				&S3AccessGrantsScopeUnionProperty{
//   					ReadWriteAccess: &ReadWriteAccessProperty{
//   						Authorization: jsii.String("authorization"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SsoTagKeys: []*string{
//   		jsii.String("ssoTagKeys"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html
//
type CfnRedshiftIdcApplicationMixinProps struct {
	// The type of application being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-applicationtype
	//
	ApplicationType *string `field:"optional" json:"applicationType" yaml:"applicationType"`
	// The token issuer list for the Amazon Redshift IAM Identity Center application instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-authorizedtokenissuerlist
	//
	AuthorizedTokenIssuerList interface{} `field:"optional" json:"authorizedTokenIssuerList" yaml:"authorizedTokenIssuerList"`
	// The IAM role ARN for the Amazon Redshift IAM Identity Center application instance.
	//
	// It has the required permissions to be assumed and invoke the IDC Identity Center API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-iamrolearn
	//
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The display name for the Amazon Redshift IAM Identity Center application instance.
	//
	// It appears in the console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-idcdisplayname
	//
	IdcDisplayName *string `field:"optional" json:"idcDisplayName" yaml:"idcDisplayName"`
	// The Amazon resource name (ARN) of the IAM Identity Center instance where Amazon Redshift creates a new managed application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-idcinstancearn
	//
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The namespace for the Amazon Redshift IAM Identity Center application instance.
	//
	// It determines which managed application verifies the connection token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-identitynamespace
	//
	IdentityNamespace *string `field:"optional" json:"identityNamespace" yaml:"identityNamespace"`
	// The name of the Redshift application in IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-redshiftidcapplicationname
	//
	RedshiftIdcApplicationName *string `field:"optional" json:"redshiftIdcApplicationName" yaml:"redshiftIdcApplicationName"`
	// A collection of service integrations for the Redshift IAM Identity Center application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-serviceintegrations
	//
	ServiceIntegrations interface{} `field:"optional" json:"serviceIntegrations" yaml:"serviceIntegrations"`
	// A list of tag keys that Redshift Identity Center applications copy to IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-ssotagkeys
	//
	SsoTagKeys *[]*string `field:"optional" json:"ssoTagKeys" yaml:"ssoTagKeys"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html#cfn-redshift-redshiftidcapplication-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

