package awsssm


// The targets for the cloud connector.
//
// If omitted, the entire organization is targeted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gcpConfigurationTargetsProperty := &GcpConfigurationTargetsProperty{
//   	Projects: []interface{}{
//   		&GcpProjectProperty{
//   			ProjectId: jsii.String("projectId"),
//
//   			// the properties below are optional
//   			ProjectDisplayName: jsii.String("projectDisplayName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfigurationtargets.html
//
type CfnCloudConnector_GcpConfigurationTargetsProperty struct {
	// List of GCP projects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfigurationtargets.html#cfn-ssm-cloudconnector-gcpconfigurationtargets-projects
	//
	Projects interface{} `field:"required" json:"projects" yaml:"projects"`
}

