package previewawsguarddutyevents


// Type definition for Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   action := &Action{
//   	ActionType: []*string{
//   		jsii.String("actionType"),
//   	},
//   	AwsApiCallAction: &AwsApiCallAction1{
//   		AffectedResources: &AffectedResources1{
//   			AwsCloudTrailTrail: []*string{
//   				jsii.String("awsCloudTrailTrail"),
//   			},
//   			AwsEc2Instance: []*string{
//   				jsii.String("awsEc2Instance"),
//   			},
//   			AwsS3Bucket: []*string{
//   				jsii.String("awsS3Bucket"),
//   			},
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
//   		RemoteAccountDetails: &RemoteAccountDetails{
//   			AccountId: []*string{
//   				jsii.String("accountId"),
//   			},
//   			Affiliated: []*string{
//   				jsii.String("affiliated"),
//   			},
//   		},
//   		RemoteIpDetails: &RemoteIpDetails1{
//   			City: &City1{
//   				CityName: []*string{
//   					jsii.String("cityName"),
//   				},
//   			},
//   			Country: &Country1{
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
//   			Organization: &Organization1{
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
//   	DnsRequestAction: &DnsRequestAction{
//   		Blocked: []*string{
//   			jsii.String("blocked"),
//   		},
//   		Domain: []*string{
//   			jsii.String("domain"),
//   		},
//   		Protocol: []*string{
//   			jsii.String("protocol"),
//   		},
//   	},
//   	KubernetesApiCallAction: &KubernetesApiCallAction{
//   		Parameters: []*string{
//   			jsii.String("parameters"),
//   		},
//   		RemoteIpDetails: &RemoteIpDetails2{
//   			City: &City2{
//   				CityName: []*string{
//   					jsii.String("cityName"),
//   				},
//   			},
//   			Country: &Country2{
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
//   			Organization: &Organization2{
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
//   		RequestUri: []*string{
//   			jsii.String("requestUri"),
//   		},
//   		SourceIPs: []*string{
//   			jsii.String("sourceIPs"),
//   		},
//   		StatusCode: []*string{
//   			jsii.String("statusCode"),
//   		},
//   		UserAgent: []*string{
//   			jsii.String("userAgent"),
//   		},
//   		Verb: []*string{
//   			jsii.String("verb"),
//   		},
//   	},
//   	NetworkConnectionAction: &NetworkConnectionAction{
//   		Blocked: []*string{
//   			jsii.String("blocked"),
//   		},
//   		ConnectionDirection: []*string{
//   			jsii.String("connectionDirection"),
//   		},
//   		LocalIpDetails: &LocalIpDetails{
//   			IpAddressV4: []*string{
//   				jsii.String("ipAddressV4"),
//   			},
//   		},
//   		LocalPortDetails: &LocalPortDetails{
//   			Port: []*string{
//   				jsii.String("port"),
//   			},
//   			PortName: []*string{
//   				jsii.String("portName"),
//   			},
//   		},
//   		Protocol: []*string{
//   			jsii.String("protocol"),
//   		},
//   		RemoteIpDetails: &RemoteIpDetails3{
//   			City: &City3{
//   				CityName: []*string{
//   					jsii.String("cityName"),
//   				},
//   			},
//   			Country: &Country3{
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
//   			Organization: &Organization3{
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
//   		RemotePortDetails: &RemotePortDetails{
//   			Port: []*string{
//   				jsii.String("port"),
//   			},
//   			PortName: []*string{
//   				jsii.String("portName"),
//   			},
//   		},
//   	},
//   	PortProbeAction: &PortProbeAction{
//   		Blocked: []*string{
//   			jsii.String("blocked"),
//   		},
//   		PortProbeDetails: []PortProbeActionItem{
//   			&PortProbeActionItem{
//   				LocalIpDetails: &LocalIpDetails1{
//   					IpAddressV4: []*string{
//   						jsii.String("ipAddressV4"),
//   					},
//   				},
//   				LocalPortDetails: &LocalPortDetails1{
//   					Port: []*string{
//   						jsii.String("port"),
//   					},
//   					PortName: []*string{
//   						jsii.String("portName"),
//   					},
//   				},
//   				RemoteIpDetails: &RemoteIpDetails4{
//   					City: &City4{
//   						CityName: []*string{
//   							jsii.String("cityName"),
//   						},
//   					},
//   					Country: &Country4{
//   						CountryName: []*string{
//   							jsii.String("countryName"),
//   						},
//   					},
//   					GeoLocation: &GeoLocation1{
//   						Lat: []*string{
//   							jsii.String("lat"),
//   						},
//   						Lon: []*string{
//   							jsii.String("lon"),
//   						},
//   					},
//   					IpAddressV4: []*string{
//   						jsii.String("ipAddressV4"),
//   					},
//   					Organization: &Organization4{
//   						Asn: []*string{
//   							jsii.String("asn"),
//   						},
//   						AsnOrg: []*string{
//   							jsii.String("asnOrg"),
//   						},
//   						Isp: []*string{
//   							jsii.String("isp"),
//   						},
//   						Org: []*string{
//   							jsii.String("org"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_Action struct {
	// actionType property.
	//
	// Specify an array of string values to match this event if the actual value of actionType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionType *[]*string `field:"optional" json:"actionType" yaml:"actionType"`
	// awsApiCallAction property.
	//
	// Specify an array of string values to match this event if the actual value of awsApiCallAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsApiCallAction *DetectorEvents_GuardDutyFinding_AwsApiCallAction1 `field:"optional" json:"awsApiCallAction" yaml:"awsApiCallAction"`
	// dnsRequestAction property.
	//
	// Specify an array of string values to match this event if the actual value of dnsRequestAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DnsRequestAction *DetectorEvents_GuardDutyFinding_DnsRequestAction `field:"optional" json:"dnsRequestAction" yaml:"dnsRequestAction"`
	// kubernetesApiCallAction property.
	//
	// Specify an array of string values to match this event if the actual value of kubernetesApiCallAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KubernetesApiCallAction *DetectorEvents_GuardDutyFinding_KubernetesApiCallAction `field:"optional" json:"kubernetesApiCallAction" yaml:"kubernetesApiCallAction"`
	// networkConnectionAction property.
	//
	// Specify an array of string values to match this event if the actual value of networkConnectionAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkConnectionAction *DetectorEvents_GuardDutyFinding_NetworkConnectionAction `field:"optional" json:"networkConnectionAction" yaml:"networkConnectionAction"`
	// portProbeAction property.
	//
	// Specify an array of string values to match this event if the actual value of portProbeAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PortProbeAction *DetectorEvents_GuardDutyFinding_PortProbeAction `field:"optional" json:"portProbeAction" yaml:"portProbeAction"`
}

