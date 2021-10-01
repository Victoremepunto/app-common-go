// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package v1

import "fmt"
import "encoding/json"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *TopicConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	if v, ok := raw["requestedName"]; !ok || v == nil {
		return fmt.Errorf("field requestedName: required")
	}
	type Plain TopicConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TopicConfig(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DatabaseConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["adminPassword"]; !ok || v == nil {
		return fmt.Errorf("field adminPassword: required")
	}
	if v, ok := raw["adminUsername"]; !ok || v == nil {
		return fmt.Errorf("field adminUsername: required")
	}
	if v, ok := raw["hostname"]; !ok || v == nil {
		return fmt.Errorf("field hostname: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	if v, ok := raw["password"]; !ok || v == nil {
		return fmt.Errorf("field password: required")
	}
	if v, ok := raw["port"]; !ok || v == nil {
		return fmt.Errorf("field port: required")
	}
	if v, ok := raw["sslMode"]; !ok || v == nil {
		return fmt.Errorf("field sslMode: required")
	}
	if v, ok := raw["username"]; !ok || v == nil {
		return fmt.Errorf("field username: required")
	}
	type Plain DatabaseConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DatabaseConfig(plain)
	return nil
}

// ClowdApp deployment configuration for Clowder enabled apps.
type AppConfig struct {
	// Database corresponds to the JSON schema field "database".
	Database *DatabaseConfig `json:"database,omitempty"`

	// Endpoints corresponds to the JSON schema field "endpoints".
	Endpoints []DependencyEndpoint `json:"endpoints,omitempty"`

	// FeatureFlags corresponds to the JSON schema field "featureFlags".
	FeatureFlags *FeatureFlagsConfig `json:"featureFlags,omitempty"`

	// InMemoryDb corresponds to the JSON schema field "inMemoryDb".
	InMemoryDb *InMemoryDBConfig `json:"inMemoryDb,omitempty"`

	// Kafka corresponds to the JSON schema field "kafka".
	Kafka *KafkaConfig `json:"kafka,omitempty"`

	// Logging corresponds to the JSON schema field "logging".
	Logging LoggingConfig `json:"logging"`

	// Metadata corresponds to the JSON schema field "metadata".
	Metadata *AppMetadata `json:"metadata,omitempty"`

	// Defines the path to the metrics server that the app should be configured to
	// listen on for metric traffic.
	MetricsPath string `json:"metricsPath"`

	// Defines the metrics port that the app should be configured to listen on for
	// metric traffic.
	MetricsPort int `json:"metricsPort"`

	// ObjectStore corresponds to the JSON schema field "objectStore".
	ObjectStore *ObjectStoreConfig `json:"objectStore,omitempty"`

	// PrivateEndpoints corresponds to the JSON schema field "privateEndpoints".
	PrivateEndpoints []PrivateDependencyEndpoint `json:"privateEndpoints,omitempty"`

	// Defines the private port that the app should be configured to listen on for API
	// traffic.
	PrivatePort *int `json:"privatePort,omitempty"`

	// Defines the public port that the app should be configured to listen on for API
	// traffic.
	PublicPort *int `json:"publicPort,omitempty"`

	// Deprecated: Use 'publicPort' instead.
	WebPort *int `json:"webPort,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DependencyEndpoint) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["app"]; !ok || v == nil {
		return fmt.Errorf("field app: required")
	}
	if v, ok := raw["hostname"]; !ok || v == nil {
		return fmt.Errorf("field hostname: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	if v, ok := raw["port"]; !ok || v == nil {
		return fmt.Errorf("field port: required")
	}
	type Plain DependencyEndpoint
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DependencyEndpoint(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PrivateDependencyEndpoint) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["app"]; !ok || v == nil {
		return fmt.Errorf("field app: required")
	}
	if v, ok := raw["hostname"]; !ok || v == nil {
		return fmt.Errorf("field hostname: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	if v, ok := raw["port"]; !ok || v == nil {
		return fmt.Errorf("field port: required")
	}
	type Plain PrivateDependencyEndpoint
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PrivateDependencyEndpoint(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectStoreConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hostname"]; !ok || v == nil {
		return fmt.Errorf("field hostname: required")
	}
	if v, ok := raw["port"]; !ok || v == nil {
		return fmt.Errorf("field port: required")
	}
	if v, ok := raw["tls"]; !ok || v == nil {
		return fmt.Errorf("field tls: required")
	}
	type Plain ObjectStoreConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ObjectStoreConfig(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FeatureFlagsConfigScheme) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_FeatureFlagsConfigScheme {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_FeatureFlagsConfigScheme, v)
	}
	*j = FeatureFlagsConfigScheme(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectStoreBucket) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	if v, ok := raw["requestedName"]; !ok || v == nil {
		return fmt.Errorf("field requestedName: required")
	}
	type Plain ObjectStoreBucket
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ObjectStoreBucket(plain)
	return nil
}

// Arbitrary metadata pertaining to the application application
type AppMetadata struct {
	// Metadata pertaining to an application's deployments
	Deployments []DeploymentMetadata `json:"deployments,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DeploymentMetadata) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["image"]; !ok || v == nil {
		return fmt.Errorf("field image: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain DeploymentMetadata
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DeploymentMetadata(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FeatureFlagsConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hostname"]; !ok || v == nil {
		return fmt.Errorf("field hostname: required")
	}
	if v, ok := raw["port"]; !ok || v == nil {
		return fmt.Errorf("field port: required")
	}
	if v, ok := raw["scheme"]; !ok || v == nil {
		return fmt.Errorf("field scheme: required")
	}
	type Plain FeatureFlagsConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FeatureFlagsConfig(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LoggingConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain LoggingConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = LoggingConfig(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *InMemoryDBConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hostname"]; !ok || v == nil {
		return fmt.Errorf("field hostname: required")
	}
	if v, ok := raw["port"]; !ok || v == nil {
		return fmt.Errorf("field port: required")
	}
	type Plain InMemoryDBConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = InMemoryDBConfig(plain)
	return nil
}

type BrokerConfigAuthtype string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CloudWatchConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["accessKeyId"]; !ok || v == nil {
		return fmt.Errorf("field accessKeyId: required")
	}
	if v, ok := raw["logGroup"]; !ok || v == nil {
		return fmt.Errorf("field logGroup: required")
	}
	if v, ok := raw["region"]; !ok || v == nil {
		return fmt.Errorf("field region: required")
	}
	if v, ok := raw["secretAccessKey"]; !ok || v == nil {
		return fmt.Errorf("field secretAccessKey: required")
	}
	type Plain CloudWatchConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CloudWatchConfig(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BrokerConfigAuthtype) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_BrokerConfigAuthtype {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_BrokerConfigAuthtype, v)
	}
	*j = BrokerConfigAuthtype(v)
	return nil
}

const BrokerConfigAuthtypeMtls BrokerConfigAuthtype = "mtls"
const BrokerConfigAuthtypeSasl BrokerConfigAuthtype = "sasl"

// Cloud Watch configuration
type CloudWatchConfig struct {
	// Defines the access key that the app should use for configuring CloudWatch.
	AccessKeyId string `json:"accessKeyId"`

	// Defines the logGroup that the app should use for configuring CloudWatch.
	LogGroup string `json:"logGroup"`

	// Defines the region that the app should use for configuring CloudWatch.
	Region string `json:"region"`

	// Defines the secret key that the app should use for configuring CloudWatch.
	SecretAccessKey string `json:"secretAccessKey"`
}

// Broker Configuration
type BrokerConfig struct {
	// Authtype corresponds to the JSON schema field "authtype".
	Authtype *BrokerConfigAuthtype `json:"authtype,omitempty"`

	// Cacert corresponds to the JSON schema field "cacert".
	Cacert *string `json:"cacert,omitempty"`

	// Hostname corresponds to the JSON schema field "hostname".
	Hostname string `json:"hostname"`

	// Port corresponds to the JSON schema field "port".
	Port *int `json:"port,omitempty"`

	// Sasl corresponds to the JSON schema field "sasl".
	Sasl *KafkaSASLConfig `json:"sasl,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BrokerConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hostname"]; !ok || v == nil {
		return fmt.Errorf("field hostname: required")
	}
	type Plain BrokerConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BrokerConfig(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KafkaConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["brokers"]; !ok || v == nil {
		return fmt.Errorf("field brokers: required")
	}
	if v, ok := raw["topics"]; !ok || v == nil {
		return fmt.Errorf("field topics: required")
	}
	type Plain KafkaConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = KafkaConfig(plain)
	return nil
}

// Database Configuration
type DatabaseConfig struct {
	// Defines the pgAdmin password.
	AdminPassword string `json:"adminPassword"`

	// Defines the pgAdmin username.
	AdminUsername string `json:"adminUsername"`

	// Defines the hostname of the database configured for the ClowdApp.
	Hostname string `json:"hostname"`

	// Defines the database name.
	Name string `json:"name"`

	// Defines the password for the standard user.
	Password string `json:"password"`

	// Defines the port of the database configured for the ClowdApp.
	Port int `json:"port"`

	// Defines the CA used to access the database.
	RdsCa *string `json:"rdsCa,omitempty"`

	// Defines the postgres SSL mode that should be used.
	SslMode string `json:"sslMode"`

	// Defines a username with standard access to the database.
	Username string `json:"username"`
}

// Dependent service connection info
type DependencyEndpoint struct {
	// The app name of the ClowdApp hosting the service.
	App string `json:"app"`

	// The hostname of the dependent service.
	Hostname string `json:"hostname"`

	// The PodSpec name of the dependent service inside the ClowdApp.
	Name string `json:"name"`

	// The port of the dependent service.
	Port int `json:"port"`
}

// Deployment Metadata
type DeploymentMetadata struct {
	// Image used by deployment
	Image string `json:"image"`

	// Name of deployment
	Name string `json:"name"`
}

// Feature Flags Configuration
type FeatureFlagsConfig struct {
	// Defines the client access token to use when connect to the FeatureFlags server
	ClientAccessToken *string `json:"clientAccessToken,omitempty"`

	// Defines the hostname for the FeatureFlags server
	Hostname string `json:"hostname"`

	// Defines the port for the FeatureFlags server
	Port int `json:"port"`

	// Details the scheme to use for FeatureFlags http/https
	Scheme FeatureFlagsConfigScheme `json:"scheme"`
}

type FeatureFlagsConfigScheme string

const FeatureFlagsConfigSchemeHttp FeatureFlagsConfigScheme = "http"
const FeatureFlagsConfigSchemeHttps FeatureFlagsConfigScheme = "https"

// In Memory DB Configuration
type InMemoryDBConfig struct {
	// Defines the hostname for the In Memory DB server configuration.
	Hostname string `json:"hostname"`

	// Defines the password for the In Memory DB server configuration.
	Password *string `json:"password,omitempty"`

	// Defines the port for the In Memory DB server configuration.
	Port int `json:"port"`

	// Defines the username for the In Memory DB server configuration.
	Username *string `json:"username,omitempty"`
}

// Kafka Configuration
type KafkaConfig struct {
	// Defines the brokers the app should connect to for Kafka services.
	Brokers []BrokerConfig `json:"brokers"`

	// Defines a list of the topic configurations available to the application.
	Topics []TopicConfig `json:"topics"`
}

// SASL Configuration for Kafka
type KafkaSASLConfig struct {
	// Password corresponds to the JSON schema field "password".
	Password *string `json:"password,omitempty"`

	// Username corresponds to the JSON schema field "username".
	Username *string `json:"username,omitempty"`
}

// Logging Configuration
type LoggingConfig struct {
	// Cloudwatch corresponds to the JSON schema field "cloudwatch".
	Cloudwatch *CloudWatchConfig `json:"cloudwatch,omitempty"`

	// Defines the type of logging configuration
	Type string `json:"type"`
}

// Object Storage Bucket
type ObjectStoreBucket struct {
	// Defines the access key for specificed bucket.
	AccessKey *string `json:"accessKey,omitempty"`

	// The actual name of the bucket being accessed.
	Name string `json:"name"`

	// Defines the region for the specified bucket.
	Region *string `json:"region,omitempty"`

	// The name that was requested for the bucket in the ClowdApp.
	RequestedName string `json:"requestedName"`

	// Defines the secret key for the specified bucket.
	SecretKey *string `json:"secretKey,omitempty"`
}

// Object Storage Configuration
type ObjectStoreConfig struct {
	// Defines the access key for the Object Storage server configuration.
	AccessKey *string `json:"accessKey,omitempty"`

	// Buckets corresponds to the JSON schema field "buckets".
	Buckets []ObjectStoreBucket `json:"buckets,omitempty"`

	// Defines the hostname for the Object Storage server configuration.
	Hostname string `json:"hostname"`

	// Defines the port for the Object Storage server configuration.
	Port int `json:"port"`

	// Defines the secret key for the Object Storage server configuration.
	SecretKey *string `json:"secretKey,omitempty"`

	// Details if the Object Server uses TLS.
	Tls bool `json:"tls"`
}

// Dependent service connection info
type PrivateDependencyEndpoint struct {
	// The app name of the ClowdApp hosting the service.
	App string `json:"app"`

	// The hostname of the dependent service.
	Hostname string `json:"hostname"`

	// The PodSpec name of the dependent service inside the ClowdApp.
	Name string `json:"name"`

	// The port of the dependent service.
	Port int `json:"port"`
}

// Topic Configuration
type TopicConfig struct {
	// The name of the actual topic on the Kafka server.
	Name string `json:"name"`

	// The name that the app requested in the ClowdApp definition.
	RequestedName string `json:"requestedName"`
}

var enumValues_BrokerConfigAuthtype = []interface{}{
	"mtls",
	"sasl",
}
var enumValues_FeatureFlagsConfigScheme = []interface{}{
	"http",
	"https",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AppConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["logging"]; !ok || v == nil {
		return fmt.Errorf("field logging: required")
	}
	if v, ok := raw["metricsPath"]; !ok || v == nil {
		return fmt.Errorf("field metricsPath: required")
	}
	if v, ok := raw["metricsPort"]; !ok || v == nil {
		return fmt.Errorf("field metricsPort: required")
	}
	type Plain AppConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AppConfig(plain)
	return nil
}
