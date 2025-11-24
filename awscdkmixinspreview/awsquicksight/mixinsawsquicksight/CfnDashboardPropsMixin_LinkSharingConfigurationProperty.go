package mixinsawsquicksight


// A structure that contains the configuration of a shareable link to the dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   linkSharingConfigurationProperty := &LinkSharingConfigurationProperty{
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   			Resource: jsii.String("resource"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linksharingconfiguration.html
//
type CfnDashboardPropsMixin_LinkSharingConfigurationProperty struct {
	// A structure that contains the permissions of a shareable link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linksharingconfiguration.html#cfn-quicksight-dashboard-linksharingconfiguration-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
}

