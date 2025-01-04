package awsvpclattice


// Identifies the resource configuration in one of the following ways:.
//
// - *Amazon Resource Name (ARN)* - Supported resource-types that are provisioned by AWS services, such as RDS databases, can be identified by their ARN.
// - *Domain name* - Any domain name that is publicly resolvable.
// - *IP address* - For IPv4 and IPv6, only IP addresses in the VPC are supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceConfigurationDefinitionProperty := &ResourceConfigurationDefinitionProperty{
//   	IpResource: jsii.String("ipResource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-resourceconfigurationdefinition.html
//
type CfnResourceConfiguration_ResourceConfigurationDefinitionProperty struct {
	// The IP address of the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-resourceconfigurationdefinition.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationdefinition-ipresource
	//
	IpResource *string `field:"required" json:"ipResource" yaml:"ipResource"`
}

