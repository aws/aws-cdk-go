# AWS::MediaPackageV2 Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## AWS Elemental MediaPackage V2

MediaPackage delivers high-quality video without concern for capacity and makes it easier to implement popular DVR features such as start over, pause, and rewind. Your content will be protected with comprehensive support for DRM. The service seamlessly integrates with other AWS media services as a complete set of tools for cloud-based video processing and delivery.

This package contains constructs for working with AWS Elemental MediaPackage V2. Allowing you to define AWS Elemental MediaPackage V2 Channel Groups, Channels, Origin Endpoints, Channel Policies and Origin Endpoint Policies.

For further information on AWS Elemental MediaPackage V2, see [the documentation](https://aws.amazon.com/mediapackage/).

The following example creates an AWS Elemental MediaPackage V2 Channel Group, Channel and Origin Endpoint:

```go
var stack Stack

group := awsmediapackagev2alpha.NewChannelGroup(stack, jsii.String("MyChannelGroup"), &ChannelGroupProps{
	ChannelGroupName: jsii.String("my-test-channel-group"),
})

channel := awsmediapackagev2alpha.NewChannel(stack, jsii.String("MyChannel"), &ChannelProps{
	ChannelGroup: group,
	ChannelName: jsii.String("my-testchannel"),
	Input: awsmediapackagev2alpha.InputConfiguration_Cmaf(),
})

endpoint := awsmediapackagev2alpha.NewOriginEndpoint(stack, jsii.String("MyOriginEndpoint"), &OriginEndpointProps{
	Channel: Channel,
	OriginEndpointName: jsii.String("my-test-endpoint"),
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

## Using Factory Methods

```go
var stack Stack


// Create a channel group
group := awsmediapackagev2alpha.NewChannelGroup(stack, jsii.String("MyChannelGroup"), &ChannelGroupProps{
	ChannelGroupName: jsii.String("my-channel-group"),
})

// Add a channel using the factory method
channel := group.addChannel(jsii.String("MyChannel"), &ChannelOptions{
	ChannelName: jsii.String("my-channel"),
	Input: awsmediapackagev2alpha.InputConfiguration_Cmaf(),
})

// Add an origin endpoint using the factory method
endpoint := channel.addOriginEndpoint(jsii.String("MyEndpoint"), &OriginEndpointOptions{
	OriginEndpointName: jsii.String("my-endpoint"),
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

## Channel Group

A channel group is the top-level resource that consists of channels and origin endpoints associated with it.

The following code creates a Channel Group:

```go
var stack Stack

group := awsmediapackagev2alpha.NewChannelGroup(stack, jsii.String("MyChannelGroup"), &ChannelGroupProps{
	ChannelGroupName: jsii.String("my-test-channel-group"),
})
```

The following code imports an existing channel group using the name attribute:

```go
var stack Stack

channelGroup := awsmediapackagev2alpha.ChannelGroup_FromChannelGroupAttributes(stack, jsii.String("ImportedChannelGroup"), &ChannelGroupAttributes{
	ChannelGroupName: jsii.String("MyChannelGroup"),
})
```

You can also import from an ARN, which automatically extracts the name and region:

```go
var stack Stack

channelGroup := awsmediapackagev2alpha.ChannelGroup_FromChannelGroupArn(stack, jsii.String("ImportedChannelGroup"), jsii.String("arn:aws:mediapackagev2:us-west-2:123456789012:channelGroup/MyChannelGroup"))
```

For cross-region imports, pass the `region` parameter to ensure the correct ARN is constructed:

```go
var stack Stack

channelGroup := awsmediapackagev2alpha.ChannelGroup_FromChannelGroupAttributes(stack, jsii.String("ImportedChannelGroup"), &ChannelGroupAttributes{
	ChannelGroupName: jsii.String("MyChannelGroup"),
	Region: jsii.String("us-west-2"),
})
```

## Channel

A channel is part of a channel group and represents the entry point for a content stream into MediaPackage.

### Input Configuration

Channels support two input types: HLS and CMAF.

```go
var stack Stack
var group ChannelGroup


hlsChannel := awsmediapackagev2alpha.NewChannel(stack, jsii.String("HlsChannel"), &ChannelProps{
	ChannelGroup: group,
	Input: awsmediapackagev2alpha.InputConfiguration_Hls(),
})

cmafChannel := awsmediapackagev2alpha.NewChannel(stack, jsii.String("CmafChannel"), &ChannelProps{
	ChannelGroup: group,
	Input: awsmediapackagev2alpha.InputConfiguration_Cmaf(&CmafInputProps{
		InputSwitchConfiguration: &InputSwitchConfiguration{
			MqcsInputSwitching: jsii.Boolean(true),
		},
		OutputHeaders: []HeadersCMSD{
			awsmediapackagev2alpha.HeadersCMSD_MQCS(),
		},
	}),
})

simpleCmafChannel := awsmediapackagev2alpha.NewChannel(stack, jsii.String("SimpleCmafChannel"), &ChannelProps{
	ChannelGroup: group,
	Input: awsmediapackagev2alpha.InputConfiguration_*Cmaf(&CmafInputProps{
		OutputHeaders: []HeadersCMSD{
			awsmediapackagev2alpha.HeadersCMSD_MQCS(),
		},
	}),
})
```

### Importing an Existing Channel

The following code imports an existing channel using the name attributes:

```go
var stack Stack

channel := awsmediapackagev2alpha.Channel_FromChannelAttributes(stack, jsii.String("ImportedChannel"), &ChannelAttributes{
	ChannelName: jsii.String("MyChannel"),
	ChannelGroupName: jsii.String("MyChannelGroup"),
})
```

You can also import from an ARN:

```go
var stack Stack

channel := awsmediapackagev2alpha.Channel_FromChannelArn(stack, jsii.String("ImportedChannel"), jsii.String("arn:aws:mediapackagev2:us-west-2:123456789012:channelGroup/MyGroup/channel/MyChannel"))
```

Imported channels expose a `region` property, which is parsed from the ARN or falls back to the importing stack's region.

### Channel Resource Policy

The following code creates a resource policy directly on the channel. This
will automatically create a policy on the first call:

```go
var channel Channel

channel.addToResourcePolicy(awscdk.NewPolicyStatement(&PolicyStatementProps{
	Sid: jsii.String("AllowMediaLiveRoleToAccessEmpChannel"),
	Principals: []IPrincipal{
		awscdk.NewArnPrincipal(jsii.String("arn:aws:iam::AccountID:role/MediaLiveAccessRole")),
	},
	Effect: awscdk.Effect_ALLOW,
	Actions: []*string{
		jsii.String("mediapackagev2:PutObject"),
	},
	Resources: []*string{
		channel.ChannelArn,
	},
}))
```

## Origin Endpoint

```go
var stack Stack
var channel Channel

awsmediapackagev2alpha.NewOriginEndpoint(stack, jsii.String("myendpoint"), &OriginEndpointProps{
	Channel: Channel,
	OriginEndpointName: jsii.String("my-test-endpoint"),
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

The following code imports an existing origin endpoint using the name attributes:

```go
var stack Stack

originEndpoint := awsmediapackagev2alpha.OriginEndpoint_FromOriginEndpointAttributes(stack, jsii.String("ImportedOriginEndpoint"), &OriginEndpointAttributes{
	ChannelGroupName: jsii.String("MyChannelGroup"),
	ChannelName: jsii.String("MyChannel"),
	OriginEndpointName: jsii.String("MyExampleOriginEndpoint"),
})
```

You can also import from an ARN:

```go
var stack Stack

originEndpoint := awsmediapackagev2alpha.OriginEndpoint_FromOriginEndpointArn(stack, jsii.String("ImportedOriginEndpoint"), jsii.String("arn:aws:mediapackagev2:us-west-2:123456789012:channelGroup/MyGroup/channel/MyChannel/originEndpoint/MyEndpoint"))
```

The following code creates a resource policy on the origin endpoint. This
will automatically create a policy on the first call:

```go
var origin OriginEndpoint


origin.addToResourcePolicy(awscdk.NewPolicyStatement(&PolicyStatementProps{
	Sid: jsii.String("AllowRequestsFromCloudFront"),
	Principals: []IPrincipal{
		awscdk.NewServicePrincipal(jsii.String("cloudfront.amazonaws.com")),
	},
	Effect: awscdk.Effect_ALLOW,
	Actions: []*string{
		jsii.String("mediapackagev2:GetHeadObject"),
		jsii.String("mediapackagev2:GetObject"),
	},
	Resources: []*string{
		origin.OriginEndpointArn,
	},
	Conditions: map[string]interface{}{
		"StringEquals": map[string]*string{
			"aws:SourceArn": jsii.String("arn:aws:cloudfront::123456789012:distribution/AAAAAAAAA"),
		},
	},
}))
```

### CDN Authorization

MediaPackage V2 supports two ways to lock an origin endpoint to your CDN:

* **AWS Signature Version 4 (SigV4)** — the CDN signs requests with an IAM
  role. For Amazon CloudFront, see [CloudFront Integration](#cloudfront-integration).
  See the [SigV4 authentication guide](https://docs.aws.amazon.com/mediapackage/latest/userguide/sig-v4-authenticating-requests.html).
* **Header-based CDN authorization** — the CDN attaches a shared secret in
  a request header that MediaPackage validates. Use this when your CDN
  doesn't support SigV4. See the [CDN authorization guide](https://docs.aws.amazon.com/mediapackage/latest/userguide/cdn-auth.html).

To configure header-based authorization, set `cdnAuth` on the `OriginEndpoint`
props. The L2 auto-creates the endpoint policy with:

* a `PolicyStatement` requiring the `mediapackagev2:RequestHasMatchingCdnAuthHeader`
  condition on every `GetObject` request
* the `CdnAuthConfiguration` block that references the secrets and the read role

If you don't supply a role, one is created with the needed Secrets Manager
and KMS permissions.

```go
import secretsmanager "github.com/aws/aws-cdk-go/awscdk"

var channel Channel
var mySecret ISecret


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("OriginEndpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Ts(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
	CdnAuth: &CdnAuthConfiguration{
		Secrets: []ISecret{
			mySecret,
		},
	},
})
```

You can still call `addToResourcePolicy()` to add extra statements (e.g. a
harvester allow); they're appended to the auto-created policy alongside the
gating statement.

## Granting Permissions

### Granting Ingest Access to MediaLive

To allow AWS Elemental MediaLive to ingest content into a MediaPackage channel, use the `grants.ingest()` method:

```go
var channel Channel
var mediaLiveRole IRole


// Grant MediaLive permission to ingest content
channel.grants.Ingest(mediaLiveRole)
```

### CloudFront Integration

MediaPackage origin endpoints are designed to be used with Content Delivery Network (CDN) like Amazon CloudFront distributions. CloudFront provides caching, DDoS protection, and global content delivery for your streaming content.

The simplest way to connect CloudFront to a MediaPackage V2 endpoint is with `MediaPackageV2Origin`, which automatically creates an Origin Access Control (OAC) and wires the endpoint policy:

```go
var endpoint OriginEndpoint
var group ChannelGroup


cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: awsmediapackagev2alpha.NewMediaPackageV2Origin(endpoint, &MediaPackageV2OriginProps{
			ChannelGroup: group,
		}),
	},
})
```

This handles OAC creation, HTTPS-only origin config, and the IAM policy granting CloudFront access to the endpoint (including `GetHeadObject` for MQAR support).

For more control, you can manually configure the policy and OAC:

```go
var originEndpoint OriginEndpoint
var distribution Distribution


originEndpoint.addToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Sid: jsii.String("AllowCloudFrontServicePrincipal"),
	Principals: []IPrincipal{
		iam.NewServicePrincipal(jsii.String("cloudfront.amazonaws.com")),
	},
	Effect: iam.Effect_ALLOW,
	Actions: []*string{
		jsii.String("mediapackagev2:GetObject"),
		jsii.String("mediapackagev2:GetHeadObject"),
	},
	Resources: []*string{
		originEndpoint.OriginEndpointArn,
	},
	Conditions: map[string]interface{}{
		"StringEquals": map[string]*string{
			"aws:SourceArn": distribution.distributionArn,
		},
	},
}))
```

> **Graduation plan:** `MediaPackageV2Origin` currently lives in this alpha module. When MediaPackage V2 graduates to stable, it will move to `aws-cloudfront-origins` alongside `S3BucketOrigin` and other origin helpers.

## Manifest Configuration

MediaPackage V2 supports multiple manifest formats: HLS, Low-Latency HLS (LL-HLS), DASH, and Microsoft Smooth Streaming (MSS).

### HLS Manifests

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
			ManifestWindow: awscdk.Duration_Seconds(jsii.Number(60)),
			ProgramDateTimeInterval: awscdk.Duration_*Seconds(jsii.Number(60)),
			ScteAdMarkerHls: awsmediapackagev2alpha.AdMarkerHls_DATERANGE,
		}),
	},
})
```

### Low-Latency HLS Manifests

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_LowLatencyHLS(&LowLatencyHlsManifestConfiguration{
			ManifestName: jsii.String("index"),
			ManifestWindow: awscdk.Duration_Seconds(jsii.Number(30)),
			ProgramDateTimeInterval: awscdk.Duration_*Seconds(jsii.Number(5)),
			ChildManifestName: jsii.String("child"),
		}),
	},
})
```

### DASH Manifests

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Dash(&DashManifestConfiguration{
			ManifestName: jsii.String("index"),
			ManifestWindow: awscdk.Duration_Seconds(jsii.Number(60)),
			MinBufferTime: awscdk.Duration_*Seconds(jsii.Number(30)),
			MinUpdatePeriod: awscdk.Duration_*Seconds(jsii.Number(10)),
			SegmentTemplateFormat: awsmediapackagev2alpha.SegmentTemplateFormat_NUMBER_WITH_TIMELINE,
			PeriodTriggers: []DashPeriodTriggers{
				awsmediapackagev2alpha.DashPeriodTriggers_AVAILS,
				awsmediapackagev2alpha.DashPeriodTriggers_DRM_KEY_ROTATION,
			},
		}),
	},
})
```

### MSS Manifests

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Ism(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Mss(&MssManifestConfiguration{
			ManifestName: jsii.String("index"),
			ManifestWindow: awscdk.Duration_Seconds(jsii.Number(60)),
			ManifestLayout: awsmediapackagev2alpha.MssManifestLayout_COMPACT,
		}),
	},
})
```

### Multiple Manifests

You can configure multiple manifest formats for a single origin endpoint:

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("hls"),
		}),
		awsmediapackagev2alpha.Manifest_Dash(&DashManifestConfiguration{
			ManifestName: jsii.String("dash"),
		}),
	},
})
```

| Segment type | Supported manifests |
|--------|--------|
| Segment.cmaf() | HLS, LL-HLS, DASH |
| Segment.ts() | HLS, LL-HLS |
| Segment.ism() | MSS |

Each origin endpoint has a single segment configuration. If you need segments with different configurations, use multiple origin endpoints on the same channel.

@see https://docs.aws.amazon.com/mediapackage/latest/userguide/endpoints-create.html

## Manifest Filtering

Manifest filters control which variants are included in the manifest. Filters are type-safe and validated against the [MediaPackage manifest filtering rules](https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html).

| Filter | Method |
|--------|--------|
| Audio / video bitrate | `bitrate()`, `bitrateRange()`, `bitrateCombo()` |
| Audio channels, sample rate, video height, framerate, trickplay height | `numeric()`, `numericList()`, `numericRange()`, `numericCombo()` |
| Audio codec | `audioCodec()`, `audioCodecList()` |
| Video codec | `videoCodec()`, `videoCodecList()` |
| Video dynamic range | `videoDynamicRange()`, `videoDynamicRangeList()` |
| Trickplay type | `trickplayType()`, `trickplayTypeList()` |
| Audio / subtitle language | `text()`, `textList()` |
| Advanced patterns | `custom()` |

The following example creates an HD streaming endpoint that serves only H.264/H.265 content between 1–5 Mbps with stereo audio in English or French:

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
			FilterConfiguration: &FilterConfiguration{
				ManifestFilter: []ManifestFilter{
					awsmediapackagev2alpha.ManifestFilter_BitrateRange(awsmediapackagev2alpha.BitrateFilterKey_VIDEO_BITRATE, awscdk.Bitrate_Mbps(jsii.Number(1)), awscdk.Bitrate_*Mbps(jsii.Number(5))),
					awsmediapackagev2alpha.ManifestFilter_NumericRange(awsmediapackagev2alpha.NumericFilterKey_VIDEO_HEIGHT, jsii.Number(720), jsii.Number(1080)),
					awsmediapackagev2alpha.ManifestFilter_VideoCodecList([]VideoCodec{
						awsmediapackagev2alpha.VideoCodec_H264,
						awsmediapackagev2alpha.VideoCodec_H265,
					}),
					awsmediapackagev2alpha.ManifestFilter_Numeric(awsmediapackagev2alpha.NumericFilterKey_AUDIO_CHANNELS, jsii.Number(2)),
					awsmediapackagev2alpha.ManifestFilter_TextList(awsmediapackagev2alpha.TextFilterKey_AUDIO_LANGUAGE, []*string{
						jsii.String("en-US"),
						jsii.String("fr"),
					}),
				},
				TimeDelay: awscdk.Duration_Seconds(jsii.Number(30)),
			},
		}),
	},
})
```

For advanced patterns that combine ranges and single values, use `numericCombo()` or `bitrateCombo()`:

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
			FilterConfiguration: &FilterConfiguration{
				ManifestFilter: []ManifestFilter{
					awsmediapackagev2alpha.ManifestFilter_NumericCombo(awsmediapackagev2alpha.NumericFilterKey_VIDEO_HEIGHT, []NumericExpression{
						awsmediapackagev2alpha.NumericExpression_Range(jsii.Number(240), jsii.Number(360)),
						awsmediapackagev2alpha.NumericExpression_*Range(jsii.Number(720), jsii.Number(1080)),
						awsmediapackagev2alpha.NumericExpression_Value(jsii.Number(1440)),
					}),
				},
			},
		}),
	},
})
```

### DRM Settings

You can exclude session keys from HLS and LL-HLS multivariant playlists using the `drmSettings` filter configuration. This improves compatibility with legacy HLS clients and provides more granular access control:

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
			FilterConfiguration: &FilterConfiguration{
				DrmSettings: []eXCLUDE_SESSION_KEYS{
					awsmediapackagev2alpha.DrmSettingsKey_*eXCLUDE_SESSION_KEYS,
				},
			},
		}),
	},
})
```

## Start Tag Configuration

Configure where playback should start in HLS and LL-HLS manifests using the EXT-X-START tag:

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
			StartTag: awsmediapackagev2alpha.StartTag_Of(jsii.Number(10)),
		}),
	},
})
```

## Segment Configuration

Configure segment settings for your origin endpoint.

```go
var channel Channel


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("TsEndpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Ts(&TsSegmentProps{
		Duration: awscdk.Duration_Seconds(jsii.Number(6)),
		Name: jsii.String("segment"),
		IncludeDvbSubtitles: jsii.Boolean(true),
		UseAudioRenditionGroup: jsii.Boolean(true),
		IncludeIframeOnlyStreams: jsii.Boolean(false),
		ScteFilter: []ScteMessageType{
			awsmediapackagev2alpha.ScteMessageType_BREAK,
			awsmediapackagev2alpha.ScteMessageType_DISTRIBUTOR_ADVERTISEMENT,
		},
	}),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})

awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("CmafEndpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
		Duration: awscdk.Duration_*Seconds(jsii.Number(6)),
		Name: jsii.String("segment"),
		IncludeIframeOnlyStreams: jsii.Boolean(true),
		ScteFilter: []ScteMessageType{
			awsmediapackagev2alpha.ScteMessageType_DISTRIBUTOR_ADVERTISEMENT,
		},
	}),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_*Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})

awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_*Cmaf(),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_*Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

## Encryption and DRM

Protect your content with encryption using SPEKE (Secure Packager and Encoder Key Exchange). Each container type has its own encryption class with type-safe options:

### CMAF Encryption

```go
var channel Channel
var spekeRole IRole


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
		Encryption: awsmediapackagev2alpha.CmafEncryption_Speke(&CmafSpekeEncryptionProps{
			Method: awsmediapackagev2alpha.CmafEncryptionMethod_CBCS,
			DrmSystems: []CmafDrmSystem{
				awsmediapackagev2alpha.CmafDrmSystem_FAIRPLAY,
				awsmediapackagev2alpha.CmafDrmSystem_WIDEVINE,
			},
			ResourceId: jsii.String("my-content-id"),
			Url: jsii.String("https://example.com/speke"),
			Role: spekeRole,
			KeyRotationInterval: awscdk.Duration_Seconds(jsii.Number(300)),
			AudioPreset: awsmediapackagev2alpha.PresetSpeke20Audio_PRESET_AUDIO_2,
			VideoPreset: awsmediapackagev2alpha.PresetSpeke20Video_PRESET_VIDEO_2,
		}),
	}),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

### TS Encryption

```go
var channel Channel
var spekeRole IRole


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("TsEndpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Ts(&TsSegmentProps{
		Encryption: awsmediapackagev2alpha.TsEncryption_Speke(&TsSpekeEncryptionProps{
			Method: awsmediapackagev2alpha.TsEncryptionMethod_SAMPLE_AES,
			ResourceId: jsii.String("my-content-id"),
			Url: jsii.String("https://example.com/speke"),
			Role: spekeRole,
		}),
	}),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

TS encryption defaults the DRM system based on the method: FairPlay for `SAMPLE_AES`, Clear Key AES 128 for `AES_128`. You can override this with the `drmSystems` property using `TsDrmSystem`.

### Content Key Encryption

You can add content key encryption by providing a certificate imported into AWS Certificate Manager. Your DRM key provider must support content key encryption for this to work:

```go
var channel Channel
var spekeRole IRole
var certificate ICertificate


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
		Encryption: awsmediapackagev2alpha.CmafEncryption_Speke(&CmafSpekeEncryptionProps{
			Method: awsmediapackagev2alpha.CmafEncryptionMethod_CBCS,
			DrmSystems: []CmafDrmSystem{
				awsmediapackagev2alpha.CmafDrmSystem_FAIRPLAY,
			},
			ResourceId: jsii.String("my-content-id"),
			Url: jsii.String("https://example.com/speke"),
			Role: spekeRole,
			Certificate: *Certificate,
		}),
	}),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

### Excluding Segment DRM Metadata

For CMAF content, you can exclude DRM metadata from segments:

```go
var channel Channel
var spekeRole IRole


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
		Encryption: awsmediapackagev2alpha.CmafEncryption_Speke(&CmafSpekeEncryptionProps{
			Method: awsmediapackagev2alpha.CmafEncryptionMethod_CBCS,
			DrmSystems: []CmafDrmSystem{
				awsmediapackagev2alpha.CmafDrmSystem_FAIRPLAY,
			},
			ResourceId: jsii.String("my-content-id"),
			Url: jsii.String("https://example.com/speke"),
			Role: spekeRole,
			ExcludeSegmentDrmMetadata: jsii.Boolean(true),
		}),
	}),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

### ISM (Smooth Streaming) Encryption

ISM endpoints use CENC encryption with PlayReady. Audio and video presets are always `SHARED`, and key rotation is not supported. The DRM system defaults to PlayReady:

```go
var channel Channel
var spekeRole IRole


awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("IsmEndpoint"), &OriginEndpointProps{
	Channel: Channel,
	Segment: awsmediapackagev2alpha.Segment_Ism(&IsmSegmentProps{
		Encryption: awsmediapackagev2alpha.IsmEncryption_Speke(&IsmSpekeEncryptionProps{
			ResourceId: jsii.String("my-content-id"),
			Url: jsii.String("https://example.com/speke"),
			Role: spekeRole,
		}),
	}),
	Manifests: []Manifest{
		awsmediapackagev2alpha.Manifest_Mss(&MssManifestConfiguration{
			ManifestName: jsii.String("index"),
		}),
	},
})
```

## CloudWatch Metrics

MediaPackage V2 resources expose CloudWatch metrics for monitoring. You can create alarms and dashboards using these metrics:

```go
var channelGroup ChannelGroup
var channel Channel
var endpoint OriginEndpoint


// Create a CloudWatch alarm on channel group egress bytes
alarm := channelGroup.metricEgressBytes().CreateAlarm(this, jsii.String("HighEgress"), &CreateAlarmOptions{
	Threshold: jsii.Number(1000),
	EvaluationPeriods: jsii.Number(1),
})

// Monitor channel ingress response time
channel.metricIngressResponseTime().CreateAlarm(this, jsii.String("SlowIngress"), &CreateAlarmOptions{
	Threshold: jsii.Number(1000),
	EvaluationPeriods: jsii.Number(2),
})

// Track origin endpoint request count
requestMetric := endpoint.metricEgressRequestCount(&MetricOptions{
	Statistic: jsii.String("sum"),
	Period: awscdk.Duration_Minutes(jsii.Number(5)),
})
```

Available metrics include:

* `metricIngressBytes()` - Bytes ingested
* `metricEgressBytes()` - Bytes delivered
* `metricIngressResponseTime()` - Ingress response time (average)
* `metricEgressResponseTime()` - Egress response time (average)
* `metricIngressRequestCount()` - Number of ingress requests
* `metricEgressRequestCount()` - Number of egress requests

All metrics support standard CloudWatch metric options for customizing period, statistic, and dimensions.
