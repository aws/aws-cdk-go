package awsssm


// A GCP project with its ID and optional display name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gcpProjectProperty := &GcpProjectProperty{
//   	ProjectDisplayName: jsii.String("projectDisplayName"),
//   	ProjectId: jsii.String("projectId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpproject.html
//
type CfnCloudConnectorPropsMixin_GcpProjectProperty struct {
	// The display name of the GCP project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpproject.html#cfn-ssm-cloudconnector-gcpproject-projectdisplayname
	//
	ProjectDisplayName *string `field:"optional" json:"projectDisplayName" yaml:"projectDisplayName"`
	// The GCP project ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-gcpproject.html#cfn-ssm-cloudconnector-gcpproject-projectid
	//
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

