package awscdkredshiftalpha


// Possible Node Types to use in the cluster used for defining `ClusterProps.nodeType`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc iVpc
//
//
//   cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	Vpc: Vpc,
//   	NodeType: nodeType_RA3_XLPLUS,
//   	AvailabilityZoneRelocation: jsii.Boolean(true),
//   })
//
// Experimental.
type NodeType string

const (
	// ds2.xlarge.
	// Experimental.
	NodeType_DS2_XLARGE NodeType = "DS2_XLARGE"
	// ds2.8xlarge.
	// Experimental.
	NodeType_DS2_8XLARGE NodeType = "DS2_8XLARGE"
	// dc1.large.
	// Experimental.
	NodeType_DC1_LARGE NodeType = "DC1_LARGE"
	// dc1.8xlarge.
	// Experimental.
	NodeType_DC1_8XLARGE NodeType = "DC1_8XLARGE"
	// dc2.large.
	// Experimental.
	NodeType_DC2_LARGE NodeType = "DC2_LARGE"
	// dc2.8xlarge.
	// Experimental.
	NodeType_DC2_8XLARGE NodeType = "DC2_8XLARGE"
	// ra3.large.
	// Experimental.
	NodeType_RA3_LARGE NodeType = "RA3_LARGE"
	// ra3.xlplus.
	// Experimental.
	NodeType_RA3_XLPLUS NodeType = "RA3_XLPLUS"
	// ra3.4xlarge.
	// Experimental.
	NodeType_RA3_4XLARGE NodeType = "RA3_4XLARGE"
	// ra3.16xlarge.
	// Experimental.
	NodeType_RA3_16XLARGE NodeType = "RA3_16XLARGE"
)

