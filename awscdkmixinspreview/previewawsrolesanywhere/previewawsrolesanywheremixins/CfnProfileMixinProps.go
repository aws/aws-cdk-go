package previewawsrolesanywheremixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProfileMixinProps := &CfnProfileMixinProps{
//   	AcceptRoleSessionName: jsii.Boolean(false),
//   	AttributeMappings: []interface{}{
//   		&AttributeMappingProperty{
//   			CertificateField: jsii.String("certificateField"),
//   			MappingRules: []interface{}{
//   				&MappingRuleProperty{
//   					Specifier: jsii.String("specifier"),
//   				},
//   			},
//   		},
//   	},
//   	DurationSeconds: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   	ManagedPolicyArns: []*string{
//   		jsii.String("managedPolicyArns"),
//   	},
//   	Name: jsii.String("name"),
//   	RequireInstanceProperties: jsii.Boolean(false),
//   	RoleArns: []*string{
//   		jsii.String("roleArns"),
//   	},
//   	SessionPolicy: jsii.String("sessionPolicy"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html
//
type CfnProfileMixinProps struct {
	// Used to determine if a custom role session name will be accepted in a temporary credential request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-acceptrolesessionname
	//
	AcceptRoleSessionName interface{} `field:"optional" json:"acceptRoleSessionName" yaml:"acceptRoleSessionName"`
	// A mapping applied to the authenticating end-entity certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-attributemappings
	//
	AttributeMappings interface{} `field:"optional" json:"attributeMappings" yaml:"attributeMappings"`
	// The number of seconds vended session credentials will be valid for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-durationseconds
	//
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// The enabled status of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of managed policy ARNs.
	//
	// Managed policies identified by this list will be applied to the vended session credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-managedpolicyarns
	//
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// The customer specified name of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies whether instance properties are required in CreateSession requests with this profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-requireinstanceproperties
	//
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-rolearns
	//
	RoleArns *[]*string `field:"optional" json:"roleArns" yaml:"roleArns"`
	// A session policy that will applied to the trust boundary of the vended session credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-sessionpolicy
	//
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// A list of Tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

