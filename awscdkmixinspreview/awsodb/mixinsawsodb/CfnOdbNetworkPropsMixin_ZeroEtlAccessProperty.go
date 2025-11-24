package mixinsawsodb


// The configuration for Zero-ETL access from the ODB network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   zeroEtlAccessProperty := &ZeroEtlAccessProperty{
//   	Cidr: jsii.String("cidr"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-zeroetlaccess.html
//
type CfnOdbNetworkPropsMixin_ZeroEtlAccessProperty struct {
	// The CIDR block for the Zero-ETL access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-zeroetlaccess.html#cfn-odb-odbnetwork-zeroetlaccess-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The status of the Zero-ETL access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-zeroetlaccess.html#cfn-odb-odbnetwork-zeroetlaccess-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

