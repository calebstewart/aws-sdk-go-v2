// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

type BusinessReportFailureCode string

// Enum values for BusinessReportFailureCode
const (
	BusinessReportFailureCodeAccessDenied    BusinessReportFailureCode = "ACCESS_DENIED"
	BusinessReportFailureCodeNoSuchBucket    BusinessReportFailureCode = "NO_SUCH_BUCKET"
	BusinessReportFailureCodeInternalFailure BusinessReportFailureCode = "INTERNAL_FAILURE"
)

func (enum BusinessReportFailureCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BusinessReportFailureCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BusinessReportFormat string

// Enum values for BusinessReportFormat
const (
	BusinessReportFormatCsv    BusinessReportFormat = "CSV"
	BusinessReportFormatCsvZip BusinessReportFormat = "CSV_ZIP"
)

func (enum BusinessReportFormat) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BusinessReportFormat) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BusinessReportInterval string

// Enum values for BusinessReportInterval
const (
	BusinessReportIntervalOneDay  BusinessReportInterval = "ONE_DAY"
	BusinessReportIntervalOneWeek BusinessReportInterval = "ONE_WEEK"
)

func (enum BusinessReportInterval) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BusinessReportInterval) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BusinessReportStatus string

// Enum values for BusinessReportStatus
const (
	BusinessReportStatusRunning   BusinessReportStatus = "RUNNING"
	BusinessReportStatusSucceeded BusinessReportStatus = "SUCCEEDED"
	BusinessReportStatusFailed    BusinessReportStatus = "FAILED"
)

func (enum BusinessReportStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BusinessReportStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CommsProtocol string

// Enum values for CommsProtocol
const (
	CommsProtocolSip  CommsProtocol = "SIP"
	CommsProtocolSips CommsProtocol = "SIPS"
	CommsProtocolH323 CommsProtocol = "H323"
)

func (enum CommsProtocol) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CommsProtocol) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConferenceProviderType string

// Enum values for ConferenceProviderType
const (
	ConferenceProviderTypeChime            ConferenceProviderType = "CHIME"
	ConferenceProviderTypeBluejeans        ConferenceProviderType = "BLUEJEANS"
	ConferenceProviderTypeFuze             ConferenceProviderType = "FUZE"
	ConferenceProviderTypeGoogleHangouts   ConferenceProviderType = "GOOGLE_HANGOUTS"
	ConferenceProviderTypePolycom          ConferenceProviderType = "POLYCOM"
	ConferenceProviderTypeRingcentral      ConferenceProviderType = "RINGCENTRAL"
	ConferenceProviderTypeSkypeForBusiness ConferenceProviderType = "SKYPE_FOR_BUSINESS"
	ConferenceProviderTypeWebex            ConferenceProviderType = "WEBEX"
	ConferenceProviderTypeZoom             ConferenceProviderType = "ZOOM"
	ConferenceProviderTypeCustom           ConferenceProviderType = "CUSTOM"
)

func (enum ConferenceProviderType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConferenceProviderType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConnectionStatus string

// Enum values for ConnectionStatus
const (
	ConnectionStatusOnline  ConnectionStatus = "ONLINE"
	ConnectionStatusOffline ConnectionStatus = "OFFLINE"
)

func (enum ConnectionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceEventType string

// Enum values for DeviceEventType
const (
	DeviceEventTypeConnectionStatus DeviceEventType = "CONNECTION_STATUS"
	DeviceEventTypeDeviceStatus     DeviceEventType = "DEVICE_STATUS"
)

func (enum DeviceEventType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceEventType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceStatus string

// Enum values for DeviceStatus
const (
	DeviceStatusReady        DeviceStatus = "READY"
	DeviceStatusPending      DeviceStatus = "PENDING"
	DeviceStatusWasOffline   DeviceStatus = "WAS_OFFLINE"
	DeviceStatusDeregistered DeviceStatus = "DEREGISTERED"
	DeviceStatusFailed       DeviceStatus = "FAILED"
)

func (enum DeviceStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceStatusDetailCode string

// Enum values for DeviceStatusDetailCode
const (
	DeviceStatusDetailCodeDeviceSoftwareUpdateNeeded      DeviceStatusDetailCode = "DEVICE_SOFTWARE_UPDATE_NEEDED"
	DeviceStatusDetailCodeDeviceWasOffline                DeviceStatusDetailCode = "DEVICE_WAS_OFFLINE"
	DeviceStatusDetailCodeCredentialsAccessFailure        DeviceStatusDetailCode = "CREDENTIALS_ACCESS_FAILURE"
	DeviceStatusDetailCodeTlsVersionMismatch              DeviceStatusDetailCode = "TLS_VERSION_MISMATCH"
	DeviceStatusDetailCodeAssociationRejection            DeviceStatusDetailCode = "ASSOCIATION_REJECTION"
	DeviceStatusDetailCodeAuthenticationFailure           DeviceStatusDetailCode = "AUTHENTICATION_FAILURE"
	DeviceStatusDetailCodeDhcpFailure                     DeviceStatusDetailCode = "DHCP_FAILURE"
	DeviceStatusDetailCodeInternetUnavailable             DeviceStatusDetailCode = "INTERNET_UNAVAILABLE"
	DeviceStatusDetailCodeDnsFailure                      DeviceStatusDetailCode = "DNS_FAILURE"
	DeviceStatusDetailCodeUnknownFailure                  DeviceStatusDetailCode = "UNKNOWN_FAILURE"
	DeviceStatusDetailCodeCertificateIssuingLimitExceeded DeviceStatusDetailCode = "CERTIFICATE_ISSUING_LIMIT_EXCEEDED"
	DeviceStatusDetailCodeInvalidCertificateAuthority     DeviceStatusDetailCode = "INVALID_CERTIFICATE_AUTHORITY"
	DeviceStatusDetailCodeNetworkProfileNotFound          DeviceStatusDetailCode = "NETWORK_PROFILE_NOT_FOUND"
	DeviceStatusDetailCodeInvalidPasswordState            DeviceStatusDetailCode = "INVALID_PASSWORD_STATE"
	DeviceStatusDetailCodePasswordNotFound                DeviceStatusDetailCode = "PASSWORD_NOT_FOUND"
)

func (enum DeviceStatusDetailCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceStatusDetailCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceUsageType string

// Enum values for DeviceUsageType
const (
	DeviceUsageTypeVoice DeviceUsageType = "VOICE"
)

func (enum DeviceUsageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceUsageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DistanceUnit string

// Enum values for DistanceUnit
const (
	DistanceUnitMetric   DistanceUnit = "METRIC"
	DistanceUnitImperial DistanceUnit = "IMPERIAL"
)

func (enum DistanceUnit) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DistanceUnit) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EnablementType string

// Enum values for EnablementType
const (
	EnablementTypeEnabled EnablementType = "ENABLED"
	EnablementTypePending EnablementType = "PENDING"
)

func (enum EnablementType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EnablementType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EnablementTypeFilter string

// Enum values for EnablementTypeFilter
const (
	EnablementTypeFilterEnabled EnablementTypeFilter = "ENABLED"
	EnablementTypeFilterPending EnablementTypeFilter = "PENDING"
)

func (enum EnablementTypeFilter) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EnablementTypeFilter) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EnrollmentStatus string

// Enum values for EnrollmentStatus
const (
	EnrollmentStatusInitialized    EnrollmentStatus = "INITIALIZED"
	EnrollmentStatusPending        EnrollmentStatus = "PENDING"
	EnrollmentStatusRegistered     EnrollmentStatus = "REGISTERED"
	EnrollmentStatusDisassociating EnrollmentStatus = "DISASSOCIATING"
	EnrollmentStatusDeregistering  EnrollmentStatus = "DEREGISTERING"
)

func (enum EnrollmentStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EnrollmentStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Feature string

// Enum values for Feature
const (
	FeatureBluetooth      Feature = "BLUETOOTH"
	FeatureVolume         Feature = "VOLUME"
	FeatureNotifications  Feature = "NOTIFICATIONS"
	FeatureLists          Feature = "LISTS"
	FeatureSkills         Feature = "SKILLS"
	FeatureNetworkProfile Feature = "NETWORK_PROFILE"
	FeatureSettings       Feature = "SETTINGS"
	FeatureAll            Feature = "ALL"
)

func (enum Feature) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Feature) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Locale string

// Enum values for Locale
const (
	LocaleEnUs Locale = "en-US"
)

func (enum Locale) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Locale) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NetworkEapMethod string

// Enum values for NetworkEapMethod
const (
	NetworkEapMethodEapTls NetworkEapMethod = "EAP_TLS"
)

func (enum NetworkEapMethod) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NetworkEapMethod) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NetworkSecurityType string

// Enum values for NetworkSecurityType
const (
	NetworkSecurityTypeOpen           NetworkSecurityType = "OPEN"
	NetworkSecurityTypeWep            NetworkSecurityType = "WEP"
	NetworkSecurityTypeWpaPsk         NetworkSecurityType = "WPA_PSK"
	NetworkSecurityTypeWpa2Psk        NetworkSecurityType = "WPA2_PSK"
	NetworkSecurityTypeWpa2Enterprise NetworkSecurityType = "WPA2_ENTERPRISE"
)

func (enum NetworkSecurityType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NetworkSecurityType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RequirePin string

// Enum values for RequirePin
const (
	RequirePinYes      RequirePin = "YES"
	RequirePinNo       RequirePin = "NO"
	RequirePinOptional RequirePin = "OPTIONAL"
)

func (enum RequirePin) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RequirePin) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SkillType string

// Enum values for SkillType
const (
	SkillTypePublic  SkillType = "PUBLIC"
	SkillTypePrivate SkillType = "PRIVATE"
)

func (enum SkillType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SkillType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SkillTypeFilter string

// Enum values for SkillTypeFilter
const (
	SkillTypeFilterPublic  SkillTypeFilter = "PUBLIC"
	SkillTypeFilterPrivate SkillTypeFilter = "PRIVATE"
	SkillTypeFilterAll     SkillTypeFilter = "ALL"
)

func (enum SkillTypeFilter) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SkillTypeFilter) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SortValue string

// Enum values for SortValue
const (
	SortValueAsc  SortValue = "ASC"
	SortValueDesc SortValue = "DESC"
)

func (enum SortValue) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SortValue) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TemperatureUnit string

// Enum values for TemperatureUnit
const (
	TemperatureUnitFahrenheit TemperatureUnit = "FAHRENHEIT"
	TemperatureUnitCelsius    TemperatureUnit = "CELSIUS"
)

func (enum TemperatureUnit) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TemperatureUnit) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WakeWord string

// Enum values for WakeWord
const (
	WakeWordAlexa    WakeWord = "ALEXA"
	WakeWordAmazon   WakeWord = "AMAZON"
	WakeWordEcho     WakeWord = "ECHO"
	WakeWordComputer WakeWord = "COMPUTER"
)

func (enum WakeWord) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WakeWord) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
