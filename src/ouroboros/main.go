package main

import (
	"fmt"
	"log"
	"ouroboros/internal/api"
	"ouroboros/internal/converter"
	egressv1 "ouroboros/internal/egress/v1"
	egressv2 "ouroboros/internal/egress/v2"
	"ouroboros/internal/ingress"

	"google.golang.org/grpc"

	"github.com/bradylove/envstruct"
	"github.com/cloudfoundry-incubator/uaago"
)

type config struct {
	SubID        string `env:"SUBSCRIPTION_ID,         required"`
	UAAAddr      string `env:"UAA_ADDR,                required"`
	ClientID     string `env:"CLIENT_ID,               required"`
	ClientSecret string `env:"CLIENT_SECRET,           required"`

	LoggregatorEgressAddr     string `env:"LOGGREGATOR_EGRESS_ADDR,     required"`
	LoggregatorIngressPort    int    `env:"LOGGREGATOR_INGRESS_PORT,    required"`
	LoggregatorIngressVersion uint8  `env:"LOGGREGATOR_INGRESS_VERSION, required"`

	TLSCACert           string `env:"LOGGREGATOR_TLS_CA_CERT"`
	TLSClientCert       string `env:"LOGGREGATOR_TLS_CLIENT_CERT"`
	TLSClientKey        string `env:"LOGGREGATOR_TLS_CLIENT_KEY"`
	TLSEgressCommonName string `env:"LOGGREGATOR_TLS_EGRESS_CN"`

	DeploymentName string `env:"DEPLOYMENT_NAME, required"`
	JobName        string `env:"JOB_NAME,        required"`
	InstanceIndex  string `env:"INSTANCE_INDEX,  required"`
	InstanceIP     string `env:"INSTANCE_IP,     required"`
}

func main() {
	conf := loadConfig()
	token := fetchUaaToken(conf)

	log.Println("Starting ouroboros ingress")
	ingress.Consume(conf.LoggregatorEgressAddr, conf.SubID, token, buildWriter(conf))
}

func loadConfig() config {
	var conf config
	if err := envstruct.Load(&conf); err != nil {
		log.Fatalf("ouroboros is not happy with your environment: %s", err)
	}

	return conf
}

func fetchUaaToken(conf config) string {
	uaaClient, err := uaago.NewClient(conf.UAAAddr)
	if err != nil {
		log.Panicf("Error creating uaa client: %s", err)
	}

	token, err := uaaClient.GetAuthToken(conf.ClientID, conf.ClientSecret, true)
	if err != nil {
		log.Panicf("Error getting token from uaa: %s", err)
	}

	return token
}

func buildWriter(conf config) ingress.EnvelopeWriter {
	var writer ingress.EnvelopeWriter

	if conf.LoggregatorIngressVersion == 1 {
		log.Println("Starting ouroboros V1 egress")
		writer = egressv1.NewWriter(fmt.Sprintf("localhost:%d", conf.LoggregatorIngressPort))
	}

	if conf.LoggregatorIngressVersion == 2 {
		log.Println("Starting ouroboros V2 egress")

		creds := api.NewCredentials(
			conf.TLSClientCert,
			conf.TLSClientKey,
			conf.TLSCACert,
			conf.TLSEgressCommonName,
		)

		writer = egressv2.NewWriter(
			fmt.Sprintf("localhost:%d", conf.LoggregatorIngressPort),
			converter.NewConverter(false),
			grpc.WithTransportCredentials(creds),
		)
	}

	if writer == nil {
		log.Fatal("Invalid LOGGREGATOR_INGRESS_VERSION")
	}

	return ingress.NewMetricCounter(
		conf.DeploymentName,
		conf.JobName,
		conf.InstanceIndex,
		conf.InstanceIP,
		1000,
		writer,
	)
}
