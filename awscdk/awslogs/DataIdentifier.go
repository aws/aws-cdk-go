package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A data protection identifier.
//
// If an identifier is supported but not in this class, it can be passed in the constructor instead.
//
// Example:
//   import kinesisfirehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//
//   logGroupDestination := logs.NewLogGroup(this, jsii.String("LogGroupLambdaAudit"), &LogGroupProps{
//   	LogGroupName: jsii.String("auditDestinationForCDK"),
//   })
//
//   bucket := s3.NewBucket(this, jsii.String("audit-bucket"))
//   s3Destination := destinations.NewS3Bucket(bucket)
//
//   deliveryStream := kinesisfirehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		s3Destination,
//   	},
//   })
//
//   dataProtectionPolicy := logs.NewDataProtectionPolicy(&DataProtectionPolicyProps{
//   	Name: jsii.String("data protection policy"),
//   	Description: jsii.String("policy description"),
//   	Identifiers: []dataIdentifier{
//   		logs.*dataIdentifier_DRIVERSLICENSE_US(),
//   		 // managed data identifier
//   		logs.NewDataIdentifier(jsii.String("EmailAddress")),
//   		 // forward compatibility for new managed data identifiers
//   		logs.NewCustomDataIdentifier(jsii.String("EmployeeId"), jsii.String("EmployeeId-\\d{9}")),
//   	},
//   	 // custom data identifier
//   	LogGroupAuditDestination: logGroupDestination,
//   	S3BucketAuditDestination: bucket,
//   	DeliveryStreamNameAuditDestination: deliveryStream.DeliveryStreamName,
//   })
//
//   logs.NewLogGroup(this, jsii.String("LogGroupLambda"), &LogGroupProps{
//   	LogGroupName: jsii.String("cdkIntegLogGroup"),
//   	DataProtectionPolicy: dataProtectionPolicy,
//   })
//
type DataIdentifier interface {
	// - name of the identifier.
	Name() *string
	ToString() *string
}

// The jsii proxy struct for DataIdentifier
type jsiiProxy_DataIdentifier struct {
	_ byte // padding
}

func (j *jsiiProxy_DataIdentifier) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Create a managed data identifier not in the list of static members.
//
// This is used to maintain forward compatibility, in case a new managed identifier is supported but not updated in CDK yet.
func NewDataIdentifier(name *string) DataIdentifier {
	_init_.Initialize()

	if err := validateNewDataIdentifierParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataIdentifier{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		[]interface{}{name},
		&j,
	)

	return &j
}

// Create a managed data identifier not in the list of static members.
//
// This is used to maintain forward compatibility, in case a new managed identifier is supported but not updated in CDK yet.
func NewDataIdentifier_Override(d DataIdentifier, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		[]interface{}{name},
		d,
	)
}

func DataIdentifier_ADDRESS() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"ADDRESS",
		&returns,
	)
	return returns
}

func DataIdentifier_AWSSECRETKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"AWSSECRETKEY",
		&returns,
	)
	return returns
}

func DataIdentifier_BANKACCOUNTNUMBER_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"BANKACCOUNTNUMBER_DE",
		&returns,
	)
	return returns
}

func DataIdentifier_BANKACCOUNTNUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"BANKACCOUNTNUMBER_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_BANKACCOUNTNUMBER_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"BANKACCOUNTNUMBER_FR",
		&returns,
	)
	return returns
}

func DataIdentifier_BANKACCOUNTNUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"BANKACCOUNTNUMBER_GB",
		&returns,
	)
	return returns
}

func DataIdentifier_BANKACCOUNTNUMBER_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"BANKACCOUNTNUMBER_IT",
		&returns,
	)
	return returns
}

func DataIdentifier_BANKACCOUNTNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"BANKACCOUNTNUMBER_US",
		&returns,
	)
	return returns
}

func DataIdentifier_CEPCODE_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"CEPCODE_BR",
		&returns,
	)
	return returns
}

func DataIdentifier_CNPJ_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"CNPJ_BR",
		&returns,
	)
	return returns
}

func DataIdentifier_CPFCODE_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"CPFCODE_BR",
		&returns,
	)
	return returns
}

func DataIdentifier_CREDITCARDEXPIRATION() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"CREDITCARDEXPIRATION",
		&returns,
	)
	return returns
}

func DataIdentifier_CREDITCARDNUMBER() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"CREDITCARDNUMBER",
		&returns,
	)
	return returns
}

func DataIdentifier_CREDITCARDSECURITYCODE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"CREDITCARDSECURITYCODE",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_AT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_AT",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_AU() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_AU",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_BE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_BE",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_BG() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_BG",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_CA",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_CY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_CY",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_CZ() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_CZ",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_DE",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_DK() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_DK",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_EE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_EE",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_FI() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_FI",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_FR",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_GB",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_GR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_GR",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_HR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_HR",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_HU() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_HU",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_IE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_IE",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_IT",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_LT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_LT",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_LU() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_LU",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_LV() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_LV",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_MT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_MT",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_NL() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_NL",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_PL() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_PL",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_PT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_PT",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_RO() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_RO",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_SE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_SE",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_SI() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_SI",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_SK() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_SK",
		&returns,
	)
	return returns
}

func DataIdentifier_DRIVERSLICENSE_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRIVERSLICENSE_US",
		&returns,
	)
	return returns
}

func DataIdentifier_DRUGENFORCEMENTAGENCYNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"DRUGENFORCEMENTAGENCYNUMBER_US",
		&returns,
	)
	return returns
}

func DataIdentifier_ELECTORALROLLNUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"ELECTORALROLLNUMBER_GB",
		&returns,
	)
	return returns
}

func DataIdentifier_EMAILADDRESS() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"EMAILADDRESS",
		&returns,
	)
	return returns
}

func DataIdentifier_HEALTHCAREPROCEDURECODE_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"HEALTHCAREPROCEDURECODE_US",
		&returns,
	)
	return returns
}

func DataIdentifier_HEALTHINSURANCECARDNUMBER_EU() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"HEALTHINSURANCECARDNUMBER_EU",
		&returns,
	)
	return returns
}

func DataIdentifier_HEALTHINSURANCECLAIMNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"HEALTHINSURANCECLAIMNUMBER_US",
		&returns,
	)
	return returns
}

func DataIdentifier_HEALTHINSURANCENUMBER_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"HEALTHINSURANCENUMBER_FR",
		&returns,
	)
	return returns
}

func DataIdentifier_INDIVIDUALTAXIDENTIFICATIONNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"INDIVIDUALTAXIDENTIFICATIONNUMBER_US",
		&returns,
	)
	return returns
}

func DataIdentifier_INSEECODE_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"INSEECODE_FR",
		&returns,
	)
	return returns
}

func DataIdentifier_IPADDRESS() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"IPADDRESS",
		&returns,
	)
	return returns
}

func DataIdentifier_LATLONG() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"LATLONG",
		&returns,
	)
	return returns
}

func DataIdentifier_MEDICAREBENEFICIARYNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"MEDICAREBENEFICIARYNUMBER_US",
		&returns,
	)
	return returns
}

func DataIdentifier_NAME() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NAME",
		&returns,
	)
	return returns
}

func DataIdentifier_NATIONALDRUGCODE_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NATIONALDRUGCODE_US",
		&returns,
	)
	return returns
}

func DataIdentifier_NATIONALIDENTIFICATIONNUMBER_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NATIONALIDENTIFICATIONNUMBER_DE",
		&returns,
	)
	return returns
}

func DataIdentifier_NATIONALIDENTIFICATIONNUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NATIONALIDENTIFICATIONNUMBER_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_NATIONALIDENTIFICATIONNUMBER_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NATIONALIDENTIFICATIONNUMBER_IT",
		&returns,
	)
	return returns
}

func DataIdentifier_NATIONALINSURANCENUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NATIONALINSURANCENUMBER_GB",
		&returns,
	)
	return returns
}

func DataIdentifier_NATIONALPROVIDERID_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NATIONALPROVIDERID_US",
		&returns,
	)
	return returns
}

func DataIdentifier_NHSNUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NHSNUMBER_GB",
		&returns,
	)
	return returns
}

func DataIdentifier_NIENUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NIENUMBER_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_NIFNUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"NIFNUMBER_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_OPENSSHPRIVATEKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"OPENSSHPRIVATEKEY",
		&returns,
	)
	return returns
}

func DataIdentifier_PASSPORTNUMBER_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PASSPORTNUMBER_CA",
		&returns,
	)
	return returns
}

func DataIdentifier_PASSPORTNUMBER_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PASSPORTNUMBER_DE",
		&returns,
	)
	return returns
}

func DataIdentifier_PASSPORTNUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PASSPORTNUMBER_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_PASSPORTNUMBER_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PASSPORTNUMBER_FR",
		&returns,
	)
	return returns
}

func DataIdentifier_PASSPORTNUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PASSPORTNUMBER_GB",
		&returns,
	)
	return returns
}

func DataIdentifier_PASSPORTNUMBER_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PASSPORTNUMBER_IT",
		&returns,
	)
	return returns
}

func DataIdentifier_PASSPORTNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PASSPORTNUMBER_US",
		&returns,
	)
	return returns
}

func DataIdentifier_PERMANENTRESIDENCENUMBER_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PERMANENTRESIDENCENUMBER_CA",
		&returns,
	)
	return returns
}

func DataIdentifier_PERSONALHEALTHNUMBER_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PERSONALHEALTHNUMBER_CA",
		&returns,
	)
	return returns
}

func DataIdentifier_PGPPRIVATEKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PGPPRIVATEKEY",
		&returns,
	)
	return returns
}

func DataIdentifier_PHONENUMBER_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PHONENUMBER_BR",
		&returns,
	)
	return returns
}

func DataIdentifier_PHONENUMBER_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PHONENUMBER_DE",
		&returns,
	)
	return returns
}

func DataIdentifier_PHONENUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PHONENUMBER_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_PHONENUMBER_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PHONENUMBER_FR",
		&returns,
	)
	return returns
}

func DataIdentifier_PHONENUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PHONENUMBER_GB",
		&returns,
	)
	return returns
}

func DataIdentifier_PHONENUMBER_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PHONENUMBER_IT",
		&returns,
	)
	return returns
}

func DataIdentifier_PHONENUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PHONENUMBER_US",
		&returns,
	)
	return returns
}

func DataIdentifier_PKCSPRIVATEKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PKCSPRIVATEKEY",
		&returns,
	)
	return returns
}

func DataIdentifier_POSTALCODE_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"POSTALCODE_CA",
		&returns,
	)
	return returns
}

func DataIdentifier_PUTTYPRIVATEKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"PUTTYPRIVATEKEY",
		&returns,
	)
	return returns
}

func DataIdentifier_RGNUMBER_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"RGNUMBER_BR",
		&returns,
	)
	return returns
}

func DataIdentifier_SOCIALINSURANCENUMBER_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"SOCIALINSURANCENUMBER_CA",
		&returns,
	)
	return returns
}

func DataIdentifier_SSN_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"SSN_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_SSN_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"SSN_US",
		&returns,
	)
	return returns
}

func DataIdentifier_TAXID_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"TAXID_DE",
		&returns,
	)
	return returns
}

func DataIdentifier_TAXID_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"TAXID_ES",
		&returns,
	)
	return returns
}

func DataIdentifier_TAXID_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"TAXID_FR",
		&returns,
	)
	return returns
}

func DataIdentifier_TAXID_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"TAXID_GB",
		&returns,
	)
	return returns
}

func DataIdentifier_VEHICLEIDENTIFICATIONNUMBER() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"VEHICLEIDENTIFICATIONNUMBER",
		&returns,
	)
	return returns
}

func DataIdentifier_ZIPCODE_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.DataIdentifier",
		"ZIPCODE_US",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataIdentifier) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

