package awsssm


// Configuration for connecting to Google Cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gcpConfigurationProperty := &GcpConfigurationProperty{
//   	ServiceAccountEmail: jsii.String("serviceAccountEmail"),
//   	WorkloadIdentityPoolProvider: jsii.String("workloadIdentityPoolProvider"),
//
//   	// the properties below are optional
//   	OrganizationId: jsii.String("organizationId"),
//   	Targets: &GcpConfigurationTargetsProperty{
//   		Projects: []interface{}{
//   			&GcpProjectProperty{
//   				ProjectId: jsii.String("projectId"),
//
//   				// the properties below are optional
//   				ProjectDisplayName: jsii.String("projectDisplayName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html
//
type CfnCloudConnector_GcpConfigurationProperty struct {
	// The GCP service account the connector impersonates.
	//
	// Its project is the connector's management project. Cannot be changed after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html#cfn-ssm-cloudconnector-gcpconfiguration-serviceaccountemail
	//
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The resource name of the workload identity pool provider that trusts the connector's IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html#cfn-ssm-cloudconnector-gcpconfiguration-workloadidentitypoolprovider
	//
	WorkloadIdentityPoolProvider *string `field:"required" json:"workloadIdentityPoolProvider" yaml:"workloadIdentityPoolProvider"`
	// The GCP organization ID.
	//
	// Required when Targets is omitted, which targets the entire organization. Cannot be changed after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html#cfn-ssm-cloudconnector-gcpconfiguration-organizationid
	//
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// The targets for the cloud connector.
	//
	// If omitted, the entire organization is targeted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html#cfn-ssm-cloudconnector-gcpconfiguration-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

