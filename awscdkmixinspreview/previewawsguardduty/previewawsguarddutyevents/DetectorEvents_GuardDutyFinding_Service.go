package previewawsguarddutyevents


// Type definition for Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalScannedPorts interface{}
//   var unusual interface{}
//
//   service := &Service{
//   	Action: &Action{
//   		ActionType: []*string{
//   			jsii.String("actionType"),
//   		},
//   		AwsApiCallAction: &AwsApiCallAction1{
//   			AffectedResources: &AffectedResources1{
//   				AwsCloudTrailTrail: []*string{
//   					jsii.String("awsCloudTrailTrail"),
//   				},
//   				AwsEc2Instance: []*string{
//   					jsii.String("awsEc2Instance"),
//   				},
//   				AwsS3Bucket: []*string{
//   					jsii.String("awsS3Bucket"),
//   				},
//   			},
//   			Api: []*string{
//   				jsii.String("api"),
//   			},
//   			CallerType: []*string{
//   				jsii.String("callerType"),
//   			},
//   			ErrorCode: []*string{
//   				jsii.String("errorCode"),
//   			},
//   			RemoteAccountDetails: &RemoteAccountDetails{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Affiliated: []*string{
//   					jsii.String("affiliated"),
//   				},
//   			},
//   			RemoteIpDetails: &RemoteIpDetails1{
//   				City: &City1{
//   					CityName: []*string{
//   						jsii.String("cityName"),
//   					},
//   				},
//   				Country: &Country1{
//   					CountryName: []*string{
//   						jsii.String("countryName"),
//   					},
//   				},
//   				GeoLocation: &GeoLocation{
//   					Lat: []*string{
//   						jsii.String("lat"),
//   					},
//   					Lon: []*string{
//   						jsii.String("lon"),
//   					},
//   				},
//   				IpAddressV4: []*string{
//   					jsii.String("ipAddressV4"),
//   				},
//   				Organization: &Organization1{
//   					Asn: []*string{
//   						jsii.String("asn"),
//   					},
//   					AsnOrg: []*string{
//   						jsii.String("asnOrg"),
//   					},
//   					Isp: []*string{
//   						jsii.String("isp"),
//   					},
//   					Org: []*string{
//   						jsii.String("org"),
//   					},
//   				},
//   			},
//   			ServiceName: []*string{
//   				jsii.String("serviceName"),
//   			},
//   		},
//   		DnsRequestAction: &DnsRequestAction{
//   			Blocked: []*string{
//   				jsii.String("blocked"),
//   			},
//   			Domain: []*string{
//   				jsii.String("domain"),
//   			},
//   			Protocol: []*string{
//   				jsii.String("protocol"),
//   			},
//   		},
//   		KubernetesApiCallAction: &KubernetesApiCallAction{
//   			Parameters: []*string{
//   				jsii.String("parameters"),
//   			},
//   			RemoteIpDetails: &RemoteIpDetails2{
//   				City: &City2{
//   					CityName: []*string{
//   						jsii.String("cityName"),
//   					},
//   				},
//   				Country: &Country2{
//   					CountryName: []*string{
//   						jsii.String("countryName"),
//   					},
//   				},
//   				GeoLocation: &GeoLocation{
//   					Lat: []*string{
//   						jsii.String("lat"),
//   					},
//   					Lon: []*string{
//   						jsii.String("lon"),
//   					},
//   				},
//   				IpAddressV4: []*string{
//   					jsii.String("ipAddressV4"),
//   				},
//   				Organization: &Organization2{
//   					Asn: []*string{
//   						jsii.String("asn"),
//   					},
//   					AsnOrg: []*string{
//   						jsii.String("asnOrg"),
//   					},
//   					Isp: []*string{
//   						jsii.String("isp"),
//   					},
//   					Org: []*string{
//   						jsii.String("org"),
//   					},
//   				},
//   			},
//   			RequestUri: []*string{
//   				jsii.String("requestUri"),
//   			},
//   			SourceIPs: []*string{
//   				jsii.String("sourceIPs"),
//   			},
//   			StatusCode: []*string{
//   				jsii.String("statusCode"),
//   			},
//   			UserAgent: []*string{
//   				jsii.String("userAgent"),
//   			},
//   			Verb: []*string{
//   				jsii.String("verb"),
//   			},
//   		},
//   		NetworkConnectionAction: &NetworkConnectionAction{
//   			Blocked: []*string{
//   				jsii.String("blocked"),
//   			},
//   			ConnectionDirection: []*string{
//   				jsii.String("connectionDirection"),
//   			},
//   			LocalIpDetails: &LocalIpDetails{
//   				IpAddressV4: []*string{
//   					jsii.String("ipAddressV4"),
//   				},
//   			},
//   			LocalPortDetails: &LocalPortDetails{
//   				Port: []*string{
//   					jsii.String("port"),
//   				},
//   				PortName: []*string{
//   					jsii.String("portName"),
//   				},
//   			},
//   			Protocol: []*string{
//   				jsii.String("protocol"),
//   			},
//   			RemoteIpDetails: &RemoteIpDetails3{
//   				City: &City3{
//   					CityName: []*string{
//   						jsii.String("cityName"),
//   					},
//   				},
//   				Country: &Country3{
//   					CountryName: []*string{
//   						jsii.String("countryName"),
//   					},
//   				},
//   				GeoLocation: &GeoLocation{
//   					Lat: []*string{
//   						jsii.String("lat"),
//   					},
//   					Lon: []*string{
//   						jsii.String("lon"),
//   					},
//   				},
//   				IpAddressV4: []*string{
//   					jsii.String("ipAddressV4"),
//   				},
//   				Organization: &Organization3{
//   					Asn: []*string{
//   						jsii.String("asn"),
//   					},
//   					AsnOrg: []*string{
//   						jsii.String("asnOrg"),
//   					},
//   					Isp: []*string{
//   						jsii.String("isp"),
//   					},
//   					Org: []*string{
//   						jsii.String("org"),
//   					},
//   				},
//   			},
//   			RemotePortDetails: &RemotePortDetails{
//   				Port: []*string{
//   					jsii.String("port"),
//   				},
//   				PortName: []*string{
//   					jsii.String("portName"),
//   				},
//   			},
//   		},
//   		PortProbeAction: &PortProbeAction{
//   			Blocked: []*string{
//   				jsii.String("blocked"),
//   			},
//   			PortProbeDetails: []PortProbeActionItem{
//   				&PortProbeActionItem{
//   					LocalIpDetails: &LocalIpDetails1{
//   						IpAddressV4: []*string{
//   							jsii.String("ipAddressV4"),
//   						},
//   					},
//   					LocalPortDetails: &LocalPortDetails1{
//   						Port: []*string{
//   							jsii.String("port"),
//   						},
//   						PortName: []*string{
//   							jsii.String("portName"),
//   						},
//   					},
//   					RemoteIpDetails: &RemoteIpDetails4{
//   						City: &City4{
//   							CityName: []*string{
//   								jsii.String("cityName"),
//   							},
//   						},
//   						Country: &Country4{
//   							CountryName: []*string{
//   								jsii.String("countryName"),
//   							},
//   						},
//   						GeoLocation: &GeoLocation1{
//   							Lat: []*string{
//   								jsii.String("lat"),
//   							},
//   							Lon: []*string{
//   								jsii.String("lon"),
//   							},
//   						},
//   						IpAddressV4: []*string{
//   							jsii.String("ipAddressV4"),
//   						},
//   						Organization: &Organization4{
//   							Asn: []*string{
//   								jsii.String("asn"),
//   							},
//   							AsnOrg: []*string{
//   								jsii.String("asnOrg"),
//   							},
//   							Isp: []*string{
//   								jsii.String("isp"),
//   							},
//   							Org: []*string{
//   								jsii.String("org"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	AdditionalInfo: &AdditionalInfo{
//   		AdditionalScannedPorts: []interface{}{
//   			additionalScannedPorts,
//   		},
//   		Anomalies: &Anomalies{
//   			AnomalousApIs: []*string{
//   				jsii.String("anomalousApIs"),
//   			},
//   		},
//   		ApiCalls: []AdditionalInfoItem{
//   			&AdditionalInfoItem{
//   				Count: []*string{
//   					jsii.String("count"),
//   				},
//   				FirstSeen: []*string{
//   					jsii.String("firstSeen"),
//   				},
//   				LastSeen: []*string{
//   					jsii.String("lastSeen"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   		Domain: []*string{
//   			jsii.String("domain"),
//   		},
//   		InBytes: []*string{
//   			jsii.String("inBytes"),
//   		},
//   		LocalPort: []*string{
//   			jsii.String("localPort"),
//   		},
//   		NewPolicy: &NewPolicy{
//   			AllowUsersToChangePassword: []*string{
//   				jsii.String("allowUsersToChangePassword"),
//   			},
//   			HardExpiry: []*string{
//   				jsii.String("hardExpiry"),
//   			},
//   			MaxPasswordAge: []*string{
//   				jsii.String("maxPasswordAge"),
//   			},
//   			MinimumPasswordLength: []*string{
//   				jsii.String("minimumPasswordLength"),
//   			},
//   			PasswordReusePrevention: []*string{
//   				jsii.String("passwordReusePrevention"),
//   			},
//   			RequireLowercaseCharacters: []*string{
//   				jsii.String("requireLowercaseCharacters"),
//   			},
//   			RequireNumbers: []*string{
//   				jsii.String("requireNumbers"),
//   			},
//   			RequireSymbols: []*string{
//   				jsii.String("requireSymbols"),
//   			},
//   			RequireUppercaseCharacters: []*string{
//   				jsii.String("requireUppercaseCharacters"),
//   			},
//   		},
//   		OldPolicy: &OldPolicy{
//   			AllowUsersToChangePassword: []*string{
//   				jsii.String("allowUsersToChangePassword"),
//   			},
//   			HardExpiry: []*string{
//   				jsii.String("hardExpiry"),
//   			},
//   			MaxPasswordAge: []*string{
//   				jsii.String("maxPasswordAge"),
//   			},
//   			MinimumPasswordLength: []*string{
//   				jsii.String("minimumPasswordLength"),
//   			},
//   			PasswordReusePrevention: []*string{
//   				jsii.String("passwordReusePrevention"),
//   			},
//   			RequireLowercaseCharacters: []*string{
//   				jsii.String("requireLowercaseCharacters"),
//   			},
//   			RequireNumbers: []*string{
//   				jsii.String("requireNumbers"),
//   			},
//   			RequireSymbols: []*string{
//   				jsii.String("requireSymbols"),
//   			},
//   			RequireUppercaseCharacters: []*string{
//   				jsii.String("requireUppercaseCharacters"),
//   			},
//   		},
//   		OutBytes: []*string{
//   			jsii.String("outBytes"),
//   		},
//   		PortsScannedSample: []*f64{
//   			jsii.Number(123),
//   		},
//   		ProfiledBehavior: &ProfiledBehavior{
//   			FrequentProfiledApIsAccountProfiling: []*string{
//   				jsii.String("frequentProfiledApIsAccountProfiling"),
//   			},
//   			FrequentProfiledApIsUserIdentityProfiling: []*string{
//   				jsii.String("frequentProfiledApIsUserIdentityProfiling"),
//   			},
//   			FrequentProfiledAsNsAccountProfiling: []*string{
//   				jsii.String("frequentProfiledAsNsAccountProfiling"),
//   			},
//   			FrequentProfiledAsNsBucketProfiling: []*string{
//   				jsii.String("frequentProfiledAsNsBucketProfiling"),
//   			},
//   			FrequentProfiledAsNsUserIdentityProfiling: []*string{
//   				jsii.String("frequentProfiledAsNsUserIdentityProfiling"),
//   			},
//   			FrequentProfiledBucketsAccountProfiling: []*string{
//   				jsii.String("frequentProfiledBucketsAccountProfiling"),
//   			},
//   			FrequentProfiledBucketsUserIdentityProfiling: []*string{
//   				jsii.String("frequentProfiledBucketsUserIdentityProfiling"),
//   			},
//   			FrequentProfiledUserAgentsAccountProfiling: []*string{
//   				jsii.String("frequentProfiledUserAgentsAccountProfiling"),
//   			},
//   			FrequentProfiledUserAgentsUserIdentityProfiling: []*string{
//   				jsii.String("frequentProfiledUserAgentsUserIdentityProfiling"),
//   			},
//   			FrequentProfiledUserNamesAccountProfiling: []*string{
//   				jsii.String("frequentProfiledUserNamesAccountProfiling"),
//   			},
//   			FrequentProfiledUserNamesBucketProfiling: []*string{
//   				jsii.String("frequentProfiledUserNamesBucketProfiling"),
//   			},
//   			FrequentProfiledUserTypesAccountProfiling: []*string{
//   				jsii.String("frequentProfiledUserTypesAccountProfiling"),
//   			},
//   			InfrequentProfiledApIsAccountProfiling: []*string{
//   				jsii.String("infrequentProfiledApIsAccountProfiling"),
//   			},
//   			InfrequentProfiledApIsUserIdentityProfiling: []*string{
//   				jsii.String("infrequentProfiledApIsUserIdentityProfiling"),
//   			},
//   			InfrequentProfiledAsNsAccountProfiling: []*string{
//   				jsii.String("infrequentProfiledAsNsAccountProfiling"),
//   			},
//   			InfrequentProfiledAsNsBucketProfiling: []*string{
//   				jsii.String("infrequentProfiledAsNsBucketProfiling"),
//   			},
//   			InfrequentProfiledAsNsUserIdentityProfiling: []*string{
//   				jsii.String("infrequentProfiledAsNsUserIdentityProfiling"),
//   			},
//   			InfrequentProfiledBucketsAccountProfiling: []*string{
//   				jsii.String("infrequentProfiledBucketsAccountProfiling"),
//   			},
//   			InfrequentProfiledBucketsUserIdentityProfiling: []*string{
//   				jsii.String("infrequentProfiledBucketsUserIdentityProfiling"),
//   			},
//   			InfrequentProfiledUserAgentsAccountProfiling: []*string{
//   				jsii.String("infrequentProfiledUserAgentsAccountProfiling"),
//   			},
//   			InfrequentProfiledUserAgentsUserIdentityProfiling: []*string{
//   				jsii.String("infrequentProfiledUserAgentsUserIdentityProfiling"),
//   			},
//   			InfrequentProfiledUserNamesAccountProfiling: []*string{
//   				jsii.String("infrequentProfiledUserNamesAccountProfiling"),
//   			},
//   			InfrequentProfiledUserNamesBucketProfiling: []*string{
//   				jsii.String("infrequentProfiledUserNamesBucketProfiling"),
//   			},
//   			InfrequentProfiledUserTypesAccountProfiling: []*string{
//   				jsii.String("infrequentProfiledUserTypesAccountProfiling"),
//   			},
//   			NumberOfHistoricalDailyAvgApIsBucketProfiling: []*string{
//   				jsii.String("numberOfHistoricalDailyAvgApIsBucketProfiling"),
//   			},
//   			NumberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling: []*string{
//   				jsii.String("numberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling"),
//   			},
//   			NumberOfHistoricalDailyAvgApIsUserIdentityProfiling: []*string{
//   				jsii.String("numberOfHistoricalDailyAvgApIsUserIdentityProfiling"),
//   			},
//   			NumberOfHistoricalDailyMaxApIsBucketProfiling: []*string{
//   				jsii.String("numberOfHistoricalDailyMaxApIsBucketProfiling"),
//   			},
//   			NumberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling: []*string{
//   				jsii.String("numberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling"),
//   			},
//   			NumberOfHistoricalDailyMaxApIsUserIdentityProfiling: []*string{
//   				jsii.String("numberOfHistoricalDailyMaxApIsUserIdentityProfiling"),
//   			},
//   			RareProfiledApIsAccountProfiling: []*string{
//   				jsii.String("rareProfiledApIsAccountProfiling"),
//   			},
//   			RareProfiledApIsUserIdentityProfiling: []*string{
//   				jsii.String("rareProfiledApIsUserIdentityProfiling"),
//   			},
//   			RareProfiledAsNsAccountProfiling: []*string{
//   				jsii.String("rareProfiledAsNsAccountProfiling"),
//   			},
//   			RareProfiledAsNsBucketProfiling: []*string{
//   				jsii.String("rareProfiledAsNsBucketProfiling"),
//   			},
//   			RareProfiledAsNsUserIdentityProfiling: []*string{
//   				jsii.String("rareProfiledAsNsUserIdentityProfiling"),
//   			},
//   			RareProfiledBucketsAccountProfiling: []*string{
//   				jsii.String("rareProfiledBucketsAccountProfiling"),
//   			},
//   			RareProfiledBucketsUserIdentityProfiling: []*string{
//   				jsii.String("rareProfiledBucketsUserIdentityProfiling"),
//   			},
//   			RareProfiledUserAgentsAccountProfiling: []*string{
//   				jsii.String("rareProfiledUserAgentsAccountProfiling"),
//   			},
//   			RareProfiledUserAgentsUserIdentityProfiling: []*string{
//   				jsii.String("rareProfiledUserAgentsUserIdentityProfiling"),
//   			},
//   			RareProfiledUserNamesAccountProfiling: []*string{
//   				jsii.String("rareProfiledUserNamesAccountProfiling"),
//   			},
//   			RareProfiledUserNamesBucketProfiling: []*string{
//   				jsii.String("rareProfiledUserNamesBucketProfiling"),
//   			},
//   			RareProfiledUserTypesAccountProfiling: []*string{
//   				jsii.String("rareProfiledUserTypesAccountProfiling"),
//   			},
//   		},
//   		RecentCredentials: []AdditionalInfoItem1{
//   			&AdditionalInfoItem1{
//   				AccessKeyId: []*string{
//   					jsii.String("accessKeyId"),
//   				},
//   				IpAddressV4: []*string{
//   					jsii.String("ipAddressV4"),
//   				},
//   				PrincipalId: []*string{
//   					jsii.String("principalId"),
//   				},
//   			},
//   		},
//   		Sample: []*string{
//   			jsii.String("sample"),
//   		},
//   		ScannedPort: []*string{
//   			jsii.String("scannedPort"),
//   		},
//   		ThreatListName: []*string{
//   			jsii.String("threatListName"),
//   		},
//   		ThreatName: []*string{
//   			jsii.String("threatName"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   		Unusual: unusual,
//   		UnusualBehavior: &UnusualBehavior{
//   			IsUnusualUserIdentity: []*string{
//   				jsii.String("isUnusualUserIdentity"),
//   			},
//   			NumberOfPast24HoursApIsBucketProfiling: []*string{
//   				jsii.String("numberOfPast24HoursApIsBucketProfiling"),
//   			},
//   			NumberOfPast24HoursApIsBucketUserIdentityProfiling: []*string{
//   				jsii.String("numberOfPast24HoursApIsBucketUserIdentityProfiling"),
//   			},
//   			NumberOfPast24HoursApIsUserIdentityProfiling: []*string{
//   				jsii.String("numberOfPast24HoursApIsUserIdentityProfiling"),
//   			},
//   			UnusualApIsAccountProfiling: []*string{
//   				jsii.String("unusualApIsAccountProfiling"),
//   			},
//   			UnusualApIsUserIdentityProfiling: []*string{
//   				jsii.String("unusualApIsUserIdentityProfiling"),
//   			},
//   			UnusualAsNsAccountProfiling: []*string{
//   				jsii.String("unusualAsNsAccountProfiling"),
//   			},
//   			UnusualAsNsBucketProfiling: []*string{
//   				jsii.String("unusualAsNsBucketProfiling"),
//   			},
//   			UnusualAsNsUserIdentityProfiling: []*string{
//   				jsii.String("unusualAsNsUserIdentityProfiling"),
//   			},
//   			UnusualBucketsAccountProfiling: []*string{
//   				jsii.String("unusualBucketsAccountProfiling"),
//   			},
//   			UnusualBucketsUserIdentityProfiling: []*string{
//   				jsii.String("unusualBucketsUserIdentityProfiling"),
//   			},
//   			UnusualUserAgentsAccountProfiling: []*string{
//   				jsii.String("unusualUserAgentsAccountProfiling"),
//   			},
//   			UnusualUserAgentsUserIdentityProfiling: []*string{
//   				jsii.String("unusualUserAgentsUserIdentityProfiling"),
//   			},
//   			UnusualUserNamesAccountProfiling: []*string{
//   				jsii.String("unusualUserNamesAccountProfiling"),
//   			},
//   			UnusualUserNamesBucketProfiling: []*string{
//   				jsii.String("unusualUserNamesBucketProfiling"),
//   			},
//   			UnusualUserTypesAccountProfiling: []*string{
//   				jsii.String("unusualUserTypesAccountProfiling"),
//   			},
//   		},
//   		UnusualProtocol: []*string{
//   			jsii.String("unusualProtocol"),
//   		},
//   		UserAgent: &UserAgent{
//   			FullUserAgent: []*string{
//   				jsii.String("fullUserAgent"),
//   			},
//   			UserAgentCategory: []*string{
//   				jsii.String("userAgentCategory"),
//   			},
//   		},
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   	Archived: []*string{
//   		jsii.String("archived"),
//   	},
//   	AwsApiCallAction: &AwsApiCallAction{
//   		AffectedResources: []*string{
//   			jsii.String("affectedResources"),
//   		},
//   		Api: []*string{
//   			jsii.String("api"),
//   		},
//   		CallerType: []*string{
//   			jsii.String("callerType"),
//   		},
//   		ErrorCode: []*string{
//   			jsii.String("errorCode"),
//   		},
//   		RemoteIpDetails: &RemoteIpDetails{
//   			City: &City{
//   				CityName: []*string{
//   					jsii.String("cityName"),
//   				},
//   			},
//   			Country: &Country{
//   				CountryName: []*string{
//   					jsii.String("countryName"),
//   				},
//   			},
//   			GeoLocation: &GeoLocation{
//   				Lat: []*string{
//   					jsii.String("lat"),
//   				},
//   				Lon: []*string{
//   					jsii.String("lon"),
//   				},
//   			},
//   			IpAddressV4: []*string{
//   				jsii.String("ipAddressV4"),
//   			},
//   			Organization: &Organization{
//   				Asn: []*string{
//   					jsii.String("asn"),
//   				},
//   				AsnOrg: []*string{
//   					jsii.String("asnOrg"),
//   				},
//   				Isp: []*string{
//   					jsii.String("isp"),
//   				},
//   				Org: []*string{
//   					jsii.String("org"),
//   				},
//   			},
//   		},
//   		ServiceName: []*string{
//   			jsii.String("serviceName"),
//   		},
//   	},
//   	Count: []*string{
//   		jsii.String("count"),
//   	},
//   	DetectorId: []*string{
//   		jsii.String("detectorId"),
//   	},
//   	EbsVolumeScanDetails: &EbsVolumeScanDetails{
//   		ScanCompletedAt: []*string{
//   			jsii.String("scanCompletedAt"),
//   		},
//   		ScanDetections: &ScanDetections{
//   			HighestSeverityThreatDetails: &HighestSeverityThreatDetails{
//   				Count: []*string{
//   					jsii.String("count"),
//   				},
//   				Severity: []*string{
//   					jsii.String("severity"),
//   				},
//   				ThreatName: []*string{
//   					jsii.String("threatName"),
//   				},
//   			},
//   			ScannedItemCount: &ScannedItemCount{
//   				Files: []*string{
//   					jsii.String("files"),
//   				},
//   				TotalGb: []*string{
//   					jsii.String("totalGb"),
//   				},
//   				Volumes: []*string{
//   					jsii.String("volumes"),
//   				},
//   			},
//   			ThreatDetectedByName: &ThreatDetectedByName{
//   				ItemCount: []*string{
//   					jsii.String("itemCount"),
//   				},
//   				Shortened: []*string{
//   					jsii.String("shortened"),
//   				},
//   				ThreatNames: []ThreatDetectedByNameItem{
//   					&ThreatDetectedByNameItem{
//   						FilePaths: []ThreatDetectedByNameItemItem{
//   							&ThreatDetectedByNameItemItem{
//   								FileName: []*string{
//   									jsii.String("fileName"),
//   								},
//   								FilePath: []*string{
//   									jsii.String("filePath"),
//   								},
//   								Hash: []*string{
//   									jsii.String("hash"),
//   								},
//   								VolumeArn: []*string{
//   									jsii.String("volumeArn"),
//   								},
//   							},
//   						},
//   						ItemCount: []*string{
//   							jsii.String("itemCount"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						Severity: []*string{
//   							jsii.String("severity"),
//   						},
//   					},
//   				},
//   				UniqueThreatNameCount: []*string{
//   					jsii.String("uniqueThreatNameCount"),
//   				},
//   			},
//   			ThreatsDetectedItemCount: &ThreatsDetectedItemCount{
//   				Files: []*string{
//   					jsii.String("files"),
//   				},
//   			},
//   		},
//   		ScanId: []*string{
//   			jsii.String("scanId"),
//   		},
//   		ScanStartedAt: []*string{
//   			jsii.String("scanStartedAt"),
//   		},
//   		Sources: []*string{
//   			jsii.String("sources"),
//   		},
//   		TriggerFindingId: []*string{
//   			jsii.String("triggerFindingId"),
//   		},
//   	},
//   	EventFirstSeen: []*string{
//   		jsii.String("eventFirstSeen"),
//   	},
//   	EventLastSeen: []*string{
//   		jsii.String("eventLastSeen"),
//   	},
//   	Evidence: &Evidence{
//   		ThreatIntelligenceDetails: []EvidenceItem{
//   			&EvidenceItem{
//   				ThreatListName: []*string{
//   					jsii.String("threatListName"),
//   				},
//   				ThreatNames: []*string{
//   					jsii.String("threatNames"),
//   				},
//   			},
//   		},
//   	},
//   	FeatureName: []*string{
//   		jsii.String("featureName"),
//   	},
//   	ResourceRole: []*string{
//   		jsii.String("resourceRole"),
//   	},
//   	ServiceName: []*string{
//   		jsii.String("serviceName"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_Service struct {
	// action property.
	//
	// Specify an array of string values to match this event if the actual value of action is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Action *DetectorEvents_GuardDutyFinding_Action `field:"optional" json:"action" yaml:"action"`
	// additionalInfo property.
	//
	// Specify an array of string values to match this event if the actual value of additionalInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AdditionalInfo *DetectorEvents_GuardDutyFinding_AdditionalInfo `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// archived property.
	//
	// Specify an array of string values to match this event if the actual value of archived is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Archived *[]*string `field:"optional" json:"archived" yaml:"archived"`
	// awsApiCallAction property.
	//
	// Specify an array of string values to match this event if the actual value of awsApiCallAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsApiCallAction *DetectorEvents_GuardDutyFinding_AwsApiCallAction `field:"optional" json:"awsApiCallAction" yaml:"awsApiCallAction"`
	// count property.
	//
	// Specify an array of string values to match this event if the actual value of count is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Count *[]*string `field:"optional" json:"count" yaml:"count"`
	// detectorId property.
	//
	// Specify an array of string values to match this event if the actual value of detectorId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Detector reference.
	//
	// Experimental.
	DetectorId *[]*string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// ebsVolumeScanDetails property.
	//
	// Specify an array of string values to match this event if the actual value of ebsVolumeScanDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EbsVolumeScanDetails *DetectorEvents_GuardDutyFinding_EbsVolumeScanDetails `field:"optional" json:"ebsVolumeScanDetails" yaml:"ebsVolumeScanDetails"`
	// eventFirstSeen property.
	//
	// Specify an array of string values to match this event if the actual value of eventFirstSeen is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventFirstSeen *[]*string `field:"optional" json:"eventFirstSeen" yaml:"eventFirstSeen"`
	// eventLastSeen property.
	//
	// Specify an array of string values to match this event if the actual value of eventLastSeen is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventLastSeen *[]*string `field:"optional" json:"eventLastSeen" yaml:"eventLastSeen"`
	// evidence property.
	//
	// Specify an array of string values to match this event if the actual value of evidence is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Evidence *DetectorEvents_GuardDutyFinding_Evidence `field:"optional" json:"evidence" yaml:"evidence"`
	// featureName property.
	//
	// Specify an array of string values to match this event if the actual value of featureName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FeatureName *[]*string `field:"optional" json:"featureName" yaml:"featureName"`
	// resourceRole property.
	//
	// Specify an array of string values to match this event if the actual value of resourceRole is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceRole *[]*string `field:"optional" json:"resourceRole" yaml:"resourceRole"`
	// serviceName property.
	//
	// Specify an array of string values to match this event if the actual value of serviceName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServiceName *[]*string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

