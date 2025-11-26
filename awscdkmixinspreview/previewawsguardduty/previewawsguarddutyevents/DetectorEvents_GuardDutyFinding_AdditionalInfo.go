package previewawsguarddutyevents


// Type definition for AdditionalInfo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalScannedPorts interface{}
//   var unusual interface{}
//
//   additionalInfo := &AdditionalInfo{
//   	AdditionalScannedPorts: []interface{}{
//   		additionalScannedPorts,
//   	},
//   	Anomalies: &Anomalies{
//   		AnomalousApIs: []*string{
//   			jsii.String("anomalousApIs"),
//   		},
//   	},
//   	ApiCalls: []AdditionalInfoItem{
//   		&AdditionalInfoItem{
//   			Count: []*string{
//   				jsii.String("count"),
//   			},
//   			FirstSeen: []*string{
//   				jsii.String("firstSeen"),
//   			},
//   			LastSeen: []*string{
//   				jsii.String("lastSeen"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   		},
//   	},
//   	Domain: []*string{
//   		jsii.String("domain"),
//   	},
//   	InBytes: []*string{
//   		jsii.String("inBytes"),
//   	},
//   	LocalPort: []*string{
//   		jsii.String("localPort"),
//   	},
//   	NewPolicy: &NewPolicy{
//   		AllowUsersToChangePassword: []*string{
//   			jsii.String("allowUsersToChangePassword"),
//   		},
//   		HardExpiry: []*string{
//   			jsii.String("hardExpiry"),
//   		},
//   		MaxPasswordAge: []*string{
//   			jsii.String("maxPasswordAge"),
//   		},
//   		MinimumPasswordLength: []*string{
//   			jsii.String("minimumPasswordLength"),
//   		},
//   		PasswordReusePrevention: []*string{
//   			jsii.String("passwordReusePrevention"),
//   		},
//   		RequireLowercaseCharacters: []*string{
//   			jsii.String("requireLowercaseCharacters"),
//   		},
//   		RequireNumbers: []*string{
//   			jsii.String("requireNumbers"),
//   		},
//   		RequireSymbols: []*string{
//   			jsii.String("requireSymbols"),
//   		},
//   		RequireUppercaseCharacters: []*string{
//   			jsii.String("requireUppercaseCharacters"),
//   		},
//   	},
//   	OldPolicy: &OldPolicy{
//   		AllowUsersToChangePassword: []*string{
//   			jsii.String("allowUsersToChangePassword"),
//   		},
//   		HardExpiry: []*string{
//   			jsii.String("hardExpiry"),
//   		},
//   		MaxPasswordAge: []*string{
//   			jsii.String("maxPasswordAge"),
//   		},
//   		MinimumPasswordLength: []*string{
//   			jsii.String("minimumPasswordLength"),
//   		},
//   		PasswordReusePrevention: []*string{
//   			jsii.String("passwordReusePrevention"),
//   		},
//   		RequireLowercaseCharacters: []*string{
//   			jsii.String("requireLowercaseCharacters"),
//   		},
//   		RequireNumbers: []*string{
//   			jsii.String("requireNumbers"),
//   		},
//   		RequireSymbols: []*string{
//   			jsii.String("requireSymbols"),
//   		},
//   		RequireUppercaseCharacters: []*string{
//   			jsii.String("requireUppercaseCharacters"),
//   		},
//   	},
//   	OutBytes: []*string{
//   		jsii.String("outBytes"),
//   	},
//   	PortsScannedSample: []*f64{
//   		jsii.Number(123),
//   	},
//   	ProfiledBehavior: &ProfiledBehavior{
//   		FrequentProfiledApIsAccountProfiling: []*string{
//   			jsii.String("frequentProfiledApIsAccountProfiling"),
//   		},
//   		FrequentProfiledApIsUserIdentityProfiling: []*string{
//   			jsii.String("frequentProfiledApIsUserIdentityProfiling"),
//   		},
//   		FrequentProfiledAsNsAccountProfiling: []*string{
//   			jsii.String("frequentProfiledAsNsAccountProfiling"),
//   		},
//   		FrequentProfiledAsNsBucketProfiling: []*string{
//   			jsii.String("frequentProfiledAsNsBucketProfiling"),
//   		},
//   		FrequentProfiledAsNsUserIdentityProfiling: []*string{
//   			jsii.String("frequentProfiledAsNsUserIdentityProfiling"),
//   		},
//   		FrequentProfiledBucketsAccountProfiling: []*string{
//   			jsii.String("frequentProfiledBucketsAccountProfiling"),
//   		},
//   		FrequentProfiledBucketsUserIdentityProfiling: []*string{
//   			jsii.String("frequentProfiledBucketsUserIdentityProfiling"),
//   		},
//   		FrequentProfiledUserAgentsAccountProfiling: []*string{
//   			jsii.String("frequentProfiledUserAgentsAccountProfiling"),
//   		},
//   		FrequentProfiledUserAgentsUserIdentityProfiling: []*string{
//   			jsii.String("frequentProfiledUserAgentsUserIdentityProfiling"),
//   		},
//   		FrequentProfiledUserNamesAccountProfiling: []*string{
//   			jsii.String("frequentProfiledUserNamesAccountProfiling"),
//   		},
//   		FrequentProfiledUserNamesBucketProfiling: []*string{
//   			jsii.String("frequentProfiledUserNamesBucketProfiling"),
//   		},
//   		FrequentProfiledUserTypesAccountProfiling: []*string{
//   			jsii.String("frequentProfiledUserTypesAccountProfiling"),
//   		},
//   		InfrequentProfiledApIsAccountProfiling: []*string{
//   			jsii.String("infrequentProfiledApIsAccountProfiling"),
//   		},
//   		InfrequentProfiledApIsUserIdentityProfiling: []*string{
//   			jsii.String("infrequentProfiledApIsUserIdentityProfiling"),
//   		},
//   		InfrequentProfiledAsNsAccountProfiling: []*string{
//   			jsii.String("infrequentProfiledAsNsAccountProfiling"),
//   		},
//   		InfrequentProfiledAsNsBucketProfiling: []*string{
//   			jsii.String("infrequentProfiledAsNsBucketProfiling"),
//   		},
//   		InfrequentProfiledAsNsUserIdentityProfiling: []*string{
//   			jsii.String("infrequentProfiledAsNsUserIdentityProfiling"),
//   		},
//   		InfrequentProfiledBucketsAccountProfiling: []*string{
//   			jsii.String("infrequentProfiledBucketsAccountProfiling"),
//   		},
//   		InfrequentProfiledBucketsUserIdentityProfiling: []*string{
//   			jsii.String("infrequentProfiledBucketsUserIdentityProfiling"),
//   		},
//   		InfrequentProfiledUserAgentsAccountProfiling: []*string{
//   			jsii.String("infrequentProfiledUserAgentsAccountProfiling"),
//   		},
//   		InfrequentProfiledUserAgentsUserIdentityProfiling: []*string{
//   			jsii.String("infrequentProfiledUserAgentsUserIdentityProfiling"),
//   		},
//   		InfrequentProfiledUserNamesAccountProfiling: []*string{
//   			jsii.String("infrequentProfiledUserNamesAccountProfiling"),
//   		},
//   		InfrequentProfiledUserNamesBucketProfiling: []*string{
//   			jsii.String("infrequentProfiledUserNamesBucketProfiling"),
//   		},
//   		InfrequentProfiledUserTypesAccountProfiling: []*string{
//   			jsii.String("infrequentProfiledUserTypesAccountProfiling"),
//   		},
//   		NumberOfHistoricalDailyAvgApIsBucketProfiling: []*string{
//   			jsii.String("numberOfHistoricalDailyAvgApIsBucketProfiling"),
//   		},
//   		NumberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling: []*string{
//   			jsii.String("numberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling"),
//   		},
//   		NumberOfHistoricalDailyAvgApIsUserIdentityProfiling: []*string{
//   			jsii.String("numberOfHistoricalDailyAvgApIsUserIdentityProfiling"),
//   		},
//   		NumberOfHistoricalDailyMaxApIsBucketProfiling: []*string{
//   			jsii.String("numberOfHistoricalDailyMaxApIsBucketProfiling"),
//   		},
//   		NumberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling: []*string{
//   			jsii.String("numberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling"),
//   		},
//   		NumberOfHistoricalDailyMaxApIsUserIdentityProfiling: []*string{
//   			jsii.String("numberOfHistoricalDailyMaxApIsUserIdentityProfiling"),
//   		},
//   		RareProfiledApIsAccountProfiling: []*string{
//   			jsii.String("rareProfiledApIsAccountProfiling"),
//   		},
//   		RareProfiledApIsUserIdentityProfiling: []*string{
//   			jsii.String("rareProfiledApIsUserIdentityProfiling"),
//   		},
//   		RareProfiledAsNsAccountProfiling: []*string{
//   			jsii.String("rareProfiledAsNsAccountProfiling"),
//   		},
//   		RareProfiledAsNsBucketProfiling: []*string{
//   			jsii.String("rareProfiledAsNsBucketProfiling"),
//   		},
//   		RareProfiledAsNsUserIdentityProfiling: []*string{
//   			jsii.String("rareProfiledAsNsUserIdentityProfiling"),
//   		},
//   		RareProfiledBucketsAccountProfiling: []*string{
//   			jsii.String("rareProfiledBucketsAccountProfiling"),
//   		},
//   		RareProfiledBucketsUserIdentityProfiling: []*string{
//   			jsii.String("rareProfiledBucketsUserIdentityProfiling"),
//   		},
//   		RareProfiledUserAgentsAccountProfiling: []*string{
//   			jsii.String("rareProfiledUserAgentsAccountProfiling"),
//   		},
//   		RareProfiledUserAgentsUserIdentityProfiling: []*string{
//   			jsii.String("rareProfiledUserAgentsUserIdentityProfiling"),
//   		},
//   		RareProfiledUserNamesAccountProfiling: []*string{
//   			jsii.String("rareProfiledUserNamesAccountProfiling"),
//   		},
//   		RareProfiledUserNamesBucketProfiling: []*string{
//   			jsii.String("rareProfiledUserNamesBucketProfiling"),
//   		},
//   		RareProfiledUserTypesAccountProfiling: []*string{
//   			jsii.String("rareProfiledUserTypesAccountProfiling"),
//   		},
//   	},
//   	RecentCredentials: []AdditionalInfoItem1{
//   		&AdditionalInfoItem1{
//   			AccessKeyId: []*string{
//   				jsii.String("accessKeyId"),
//   			},
//   			IpAddressV4: []*string{
//   				jsii.String("ipAddressV4"),
//   			},
//   			PrincipalId: []*string{
//   				jsii.String("principalId"),
//   			},
//   		},
//   	},
//   	Sample: []*string{
//   		jsii.String("sample"),
//   	},
//   	ScannedPort: []*string{
//   		jsii.String("scannedPort"),
//   	},
//   	ThreatListName: []*string{
//   		jsii.String("threatListName"),
//   	},
//   	ThreatName: []*string{
//   		jsii.String("threatName"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   	Unusual: unusual,
//   	UnusualBehavior: &UnusualBehavior{
//   		IsUnusualUserIdentity: []*string{
//   			jsii.String("isUnusualUserIdentity"),
//   		},
//   		NumberOfPast24HoursApIsBucketProfiling: []*string{
//   			jsii.String("numberOfPast24HoursApIsBucketProfiling"),
//   		},
//   		NumberOfPast24HoursApIsBucketUserIdentityProfiling: []*string{
//   			jsii.String("numberOfPast24HoursApIsBucketUserIdentityProfiling"),
//   		},
//   		NumberOfPast24HoursApIsUserIdentityProfiling: []*string{
//   			jsii.String("numberOfPast24HoursApIsUserIdentityProfiling"),
//   		},
//   		UnusualApIsAccountProfiling: []*string{
//   			jsii.String("unusualApIsAccountProfiling"),
//   		},
//   		UnusualApIsUserIdentityProfiling: []*string{
//   			jsii.String("unusualApIsUserIdentityProfiling"),
//   		},
//   		UnusualAsNsAccountProfiling: []*string{
//   			jsii.String("unusualAsNsAccountProfiling"),
//   		},
//   		UnusualAsNsBucketProfiling: []*string{
//   			jsii.String("unusualAsNsBucketProfiling"),
//   		},
//   		UnusualAsNsUserIdentityProfiling: []*string{
//   			jsii.String("unusualAsNsUserIdentityProfiling"),
//   		},
//   		UnusualBucketsAccountProfiling: []*string{
//   			jsii.String("unusualBucketsAccountProfiling"),
//   		},
//   		UnusualBucketsUserIdentityProfiling: []*string{
//   			jsii.String("unusualBucketsUserIdentityProfiling"),
//   		},
//   		UnusualUserAgentsAccountProfiling: []*string{
//   			jsii.String("unusualUserAgentsAccountProfiling"),
//   		},
//   		UnusualUserAgentsUserIdentityProfiling: []*string{
//   			jsii.String("unusualUserAgentsUserIdentityProfiling"),
//   		},
//   		UnusualUserNamesAccountProfiling: []*string{
//   			jsii.String("unusualUserNamesAccountProfiling"),
//   		},
//   		UnusualUserNamesBucketProfiling: []*string{
//   			jsii.String("unusualUserNamesBucketProfiling"),
//   		},
//   		UnusualUserTypesAccountProfiling: []*string{
//   			jsii.String("unusualUserTypesAccountProfiling"),
//   		},
//   	},
//   	UnusualProtocol: []*string{
//   		jsii.String("unusualProtocol"),
//   	},
//   	UserAgent: &UserAgent{
//   		FullUserAgent: []*string{
//   			jsii.String("fullUserAgent"),
//   		},
//   		UserAgentCategory: []*string{
//   			jsii.String("userAgentCategory"),
//   		},
//   	},
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_AdditionalInfo struct {
	// additionalScannedPorts property.
	//
	// Specify an array of string values to match this event if the actual value of additionalScannedPorts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AdditionalScannedPorts *[]interface{} `field:"optional" json:"additionalScannedPorts" yaml:"additionalScannedPorts"`
	// anomalies property.
	//
	// Specify an array of string values to match this event if the actual value of anomalies is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Anomalies *DetectorEvents_GuardDutyFinding_Anomalies `field:"optional" json:"anomalies" yaml:"anomalies"`
	// apiCalls property.
	//
	// Specify an array of string values to match this event if the actual value of apiCalls is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApiCalls *[]*DetectorEvents_GuardDutyFinding_AdditionalInfoItem `field:"optional" json:"apiCalls" yaml:"apiCalls"`
	// domain property.
	//
	// Specify an array of string values to match this event if the actual value of domain is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Domain *[]*string `field:"optional" json:"domain" yaml:"domain"`
	// inBytes property.
	//
	// Specify an array of string values to match this event if the actual value of inBytes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InBytes *[]*string `field:"optional" json:"inBytes" yaml:"inBytes"`
	// localPort property.
	//
	// Specify an array of string values to match this event if the actual value of localPort is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LocalPort *[]*string `field:"optional" json:"localPort" yaml:"localPort"`
	// newPolicy property.
	//
	// Specify an array of string values to match this event if the actual value of newPolicy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NewPolicy *DetectorEvents_GuardDutyFinding_NewPolicy `field:"optional" json:"newPolicy" yaml:"newPolicy"`
	// oldPolicy property.
	//
	// Specify an array of string values to match this event if the actual value of oldPolicy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OldPolicy *DetectorEvents_GuardDutyFinding_OldPolicy `field:"optional" json:"oldPolicy" yaml:"oldPolicy"`
	// outBytes property.
	//
	// Specify an array of string values to match this event if the actual value of outBytes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutBytes *[]*string `field:"optional" json:"outBytes" yaml:"outBytes"`
	// portsScannedSample property.
	//
	// Specify an array of string values to match this event if the actual value of portsScannedSample is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PortsScannedSample *[]*float64 `field:"optional" json:"portsScannedSample" yaml:"portsScannedSample"`
	// profiledBehavior property.
	//
	// Specify an array of string values to match this event if the actual value of profiledBehavior is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProfiledBehavior *DetectorEvents_GuardDutyFinding_ProfiledBehavior `field:"optional" json:"profiledBehavior" yaml:"profiledBehavior"`
	// recentCredentials property.
	//
	// Specify an array of string values to match this event if the actual value of recentCredentials is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RecentCredentials *[]*DetectorEvents_GuardDutyFinding_AdditionalInfoItem1 `field:"optional" json:"recentCredentials" yaml:"recentCredentials"`
	// sample property.
	//
	// Specify an array of string values to match this event if the actual value of sample is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Sample *[]*string `field:"optional" json:"sample" yaml:"sample"`
	// scannedPort property.
	//
	// Specify an array of string values to match this event if the actual value of scannedPort is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScannedPort *[]*string `field:"optional" json:"scannedPort" yaml:"scannedPort"`
	// threatListName property.
	//
	// Specify an array of string values to match this event if the actual value of threatListName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatListName *[]*string `field:"optional" json:"threatListName" yaml:"threatListName"`
	// threatName property.
	//
	// Specify an array of string values to match this event if the actual value of threatName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatName *[]*string `field:"optional" json:"threatName" yaml:"threatName"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
	// unusual property.
	//
	// Specify an array of string values to match this event if the actual value of unusual is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Unusual interface{} `field:"optional" json:"unusual" yaml:"unusual"`
	// unusualBehavior property.
	//
	// Specify an array of string values to match this event if the actual value of unusualBehavior is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualBehavior *DetectorEvents_GuardDutyFinding_UnusualBehavior `field:"optional" json:"unusualBehavior" yaml:"unusualBehavior"`
	// unusualProtocol property.
	//
	// Specify an array of string values to match this event if the actual value of unusualProtocol is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualProtocol *[]*string `field:"optional" json:"unusualProtocol" yaml:"unusualProtocol"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *DetectorEvents_GuardDutyFinding_UserAgent `field:"optional" json:"userAgent" yaml:"userAgent"`
	// value property.
	//
	// Specify an array of string values to match this event if the actual value of value is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

