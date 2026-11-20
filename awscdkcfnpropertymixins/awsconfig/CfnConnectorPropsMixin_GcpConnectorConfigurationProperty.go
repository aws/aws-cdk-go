package awsconfig


// Configuration for connecting to Google Cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gcpConnectorConfigurationProperty := &GcpConnectorConfigurationProperty{
//   	GcpOrganizationId: jsii.String("gcpOrganizationId"),
//   	ServiceAccountEmail: jsii.String("serviceAccountEmail"),
//   	WorkloadIdentityPoolProvider: jsii.String("workloadIdentityPoolProvider"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-gcpconnectorconfiguration.html
//
type CfnConnectorPropsMixin_GcpConnectorConfigurationProperty struct {
	// The ID of the Google Cloud organization that contains the connector's management project, which is the project of the service account in ServiceAccountEmail.
	//
	// Required if the management project belongs to an organization. A connector needs an organization ID to record an entire organization or more than one project. Omit it only if the management project doesn't belong to an organization. In that case, the connector can record only the management project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-gcpconnectorconfiguration.html#cfn-config-connector-gcpconnectorconfiguration-gcporganizationid
	//
	GcpOrganizationId *string `field:"optional" json:"gcpOrganizationId" yaml:"gcpOrganizationId"`
	// The email address of the Google Cloud service account that AWS Config uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-gcpconnectorconfiguration.html#cfn-config-connector-gcpconnectorconfiguration-serviceaccountemail
	//
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The full resource name of the workload identity pool provider that AWS Config uses to federate into Google Cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-connector-gcpconnectorconfiguration.html#cfn-config-connector-gcpconnectorconfiguration-workloadidentitypoolprovider
	//
	WorkloadIdentityPoolProvider *string `field:"optional" json:"workloadIdentityPoolProvider" yaml:"workloadIdentityPoolProvider"`
}

