package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A custom data identifier.
//
// Include a custom data identifier name and regular expression in the JSON policy used to define the data protection policy.
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
type CustomDataIdentifier interface {
	DataIdentifier
	// - the name of the custom data identifier.
	//
	// This cannot share the same name as a managed data identifier.
	Name() *string
	// - the regular expresssion to detect and mask log events for.
	Regex() *string
	// String representation of a CustomDataIdentifier.
	//
	// Returns: the name and RegEx of the custom data identifier.
	ToString() *string
}

// The jsii proxy struct for CustomDataIdentifier
type jsiiProxy_CustomDataIdentifier struct {
	jsiiProxy_DataIdentifier
}

func (j *jsiiProxy_CustomDataIdentifier) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomDataIdentifier) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}


// Create a custom data identifier.
func NewCustomDataIdentifier(name *string, regex *string) CustomDataIdentifier {
	_init_.Initialize()

	if err := validateNewCustomDataIdentifierParameters(name, regex); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomDataIdentifier{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		[]interface{}{name, regex},
		&j,
	)

	return &j
}

// Create a custom data identifier.
func NewCustomDataIdentifier_Override(c CustomDataIdentifier, name *string, regex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		[]interface{}{name, regex},
		c,
	)
}

func CustomDataIdentifier_ADDRESS() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"ADDRESS",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_AWSSECRETKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"AWSSECRETKEY",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_BANKACCOUNTNUMBER_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"BANKACCOUNTNUMBER_DE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_BANKACCOUNTNUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"BANKACCOUNTNUMBER_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_BANKACCOUNTNUMBER_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"BANKACCOUNTNUMBER_FR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_BANKACCOUNTNUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"BANKACCOUNTNUMBER_GB",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_BANKACCOUNTNUMBER_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"BANKACCOUNTNUMBER_IT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_BANKACCOUNTNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"BANKACCOUNTNUMBER_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_CEPCODE_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"CEPCODE_BR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_CNPJ_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"CNPJ_BR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_CPFCODE_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"CPFCODE_BR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_CREDITCARDEXPIRATION() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"CREDITCARDEXPIRATION",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_CREDITCARDNUMBER() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"CREDITCARDNUMBER",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_CREDITCARDSECURITYCODE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"CREDITCARDSECURITYCODE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_AT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_AT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_AU() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_AU",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_BE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_BE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_BG() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_BG",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_CA",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_CY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_CY",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_CZ() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_CZ",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_DE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_DK() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_DK",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_EE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_EE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_FI() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_FI",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_FR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_GB",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_GR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_GR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_HR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_HR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_HU() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_HU",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_IE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_IE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_IT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_LT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_LT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_LU() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_LU",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_LV() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_LV",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_MT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_MT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_NL() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_NL",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_PL() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_PL",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_PT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_PT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_RO() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_RO",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_SE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_SE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_SI() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_SI",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_SK() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_SK",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRIVERSLICENSE_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRIVERSLICENSE_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_DRUGENFORCEMENTAGENCYNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"DRUGENFORCEMENTAGENCYNUMBER_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_ELECTORALROLLNUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"ELECTORALROLLNUMBER_GB",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_EMAILADDRESS() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"EMAILADDRESS",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_HEALTHCAREPROCEDURECODE_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"HEALTHCAREPROCEDURECODE_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_HEALTHINSURANCECARDNUMBER_EU() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"HEALTHINSURANCECARDNUMBER_EU",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_HEALTHINSURANCECLAIMNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"HEALTHINSURANCECLAIMNUMBER_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_HEALTHINSURANCENUMBER_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"HEALTHINSURANCENUMBER_FR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_INDIVIDUALTAXIDENTIFICATIONNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"INDIVIDUALTAXIDENTIFICATIONNUMBER_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_INSEECODE_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"INSEECODE_FR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_IPADDRESS() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"IPADDRESS",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_LATLONG() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"LATLONG",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_MEDICAREBENEFICIARYNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"MEDICAREBENEFICIARYNUMBER_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NAME() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NAME",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NATIONALDRUGCODE_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NATIONALDRUGCODE_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NATIONALIDENTIFICATIONNUMBER_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NATIONALIDENTIFICATIONNUMBER_DE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NATIONALIDENTIFICATIONNUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NATIONALIDENTIFICATIONNUMBER_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NATIONALIDENTIFICATIONNUMBER_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NATIONALIDENTIFICATIONNUMBER_IT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NATIONALINSURANCENUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NATIONALINSURANCENUMBER_GB",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NATIONALPROVIDERID_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NATIONALPROVIDERID_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NHSNUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NHSNUMBER_GB",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NIENUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NIENUMBER_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_NIFNUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"NIFNUMBER_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_OPENSSHPRIVATEKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"OPENSSHPRIVATEKEY",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PASSPORTNUMBER_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PASSPORTNUMBER_CA",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PASSPORTNUMBER_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PASSPORTNUMBER_DE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PASSPORTNUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PASSPORTNUMBER_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PASSPORTNUMBER_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PASSPORTNUMBER_FR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PASSPORTNUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PASSPORTNUMBER_GB",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PASSPORTNUMBER_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PASSPORTNUMBER_IT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PASSPORTNUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PASSPORTNUMBER_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PERMANENTRESIDENCENUMBER_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PERMANENTRESIDENCENUMBER_CA",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PERSONALHEALTHNUMBER_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PERSONALHEALTHNUMBER_CA",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PGPPRIVATEKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PGPPRIVATEKEY",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PHONENUMBER_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PHONENUMBER_BR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PHONENUMBER_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PHONENUMBER_DE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PHONENUMBER_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PHONENUMBER_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PHONENUMBER_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PHONENUMBER_FR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PHONENUMBER_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PHONENUMBER_GB",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PHONENUMBER_IT() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PHONENUMBER_IT",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PHONENUMBER_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PHONENUMBER_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PKCSPRIVATEKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PKCSPRIVATEKEY",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_POSTALCODE_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"POSTALCODE_CA",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_PUTTYPRIVATEKEY() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"PUTTYPRIVATEKEY",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_RGNUMBER_BR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"RGNUMBER_BR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_SOCIALINSURANCENUMBER_CA() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"SOCIALINSURANCENUMBER_CA",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_SSN_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"SSN_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_SSN_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"SSN_US",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_TAXID_DE() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"TAXID_DE",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_TAXID_ES() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"TAXID_ES",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_TAXID_FR() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"TAXID_FR",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_TAXID_GB() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"TAXID_GB",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_VEHICLEIDENTIFICATIONNUMBER() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"VEHICLEIDENTIFICATIONNUMBER",
		&returns,
	)
	return returns
}

func CustomDataIdentifier_ZIPCODE_US() DataIdentifier {
	_init_.Initialize()
	var returns DataIdentifier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_logs.CustomDataIdentifier",
		"ZIPCODE_US",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CustomDataIdentifier) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

