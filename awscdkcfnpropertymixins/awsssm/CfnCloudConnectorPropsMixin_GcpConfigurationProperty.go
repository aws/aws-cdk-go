package awsssm


// Configuration for connecting to Google Cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gcpConfigurationProperty := &GcpConfigurationProperty{
//   	OrganizationId: jsii.String("organizationId"),
//   	ServiceAccountEmail: jsii.String("serviceAccountEmail"),
//   	Targets: &GcpConfigurationTargetsProperty{
//   		Projects: []interface{}{
//   			&GcpProjectProperty{
//   				ProjectDisplayName: jsii.String("projectDisplayName"),
//   				ProjectId: jsii.String("projectId"),
//   			},
//   		},
//   	},
//   	WorkloadIdentityPoolProvider: jsii.String("workloadIdentityPoolProvider"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html
//
type CfnCloudConnectorPropsMixin_GcpConfigurationProperty struct {
	// The GCP organization ID.
	//
	// Required when Targets is omitted, which targets the entire organization. Cannot be changed after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html#cfn-ssm-cloudconnector-gcpconfiguration-organizationid
	//
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// The GCP service account the connector impersonates.
	//
	// Its project is the connector's management project. Cannot be changed after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html#cfn-ssm-cloudconnector-gcpconfiguration-serviceaccountemail
	//
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The targets for the cloud connector.
	//
	// If omitted, the entire organization is targeted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html#cfn-ssm-cloudconnector-gcpconfiguration-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The resource name of the workload identity pool provider that trusts the connector's IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpconfiguration.html#cfn-ssm-cloudconnector-gcpconfiguration-workloadidentitypoolprovider
	//
	WorkloadIdentityPoolProvider *string `field:"optional" json:"workloadIdentityPoolProvider" yaml:"workloadIdentityPoolProvider"`
}

