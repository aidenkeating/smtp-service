package smtpdetails

import (
	"strconv"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SMTPDetails struct {
	ID       string
	Host     string
	Port     int
	TLS      bool
	Username string
	Password string
}

type Client interface {
	Create(id string) (*SMTPDetails, error)
	Get(id string) (*SMTPDetails, error)
	Delete(id string) error
}

//ConvertSMTPDetailsToSecret Format a standard set of SMTPDetails as a Kubernetes Secret
func ConvertSMTPDetailsToSecret(smtpDetails *SMTPDetails, secretName string) *apiv1.Secret {
	return &apiv1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name: secretName,
		},
		Data: map[string][]byte{
			SecretKeyHost:     []byte(smtpDetails.Host),
			SecretKeyPort:     []byte(strconv.Itoa(smtpDetails.Port)),
			SecretKeyTLS:      []byte(strconv.FormatBool(smtpDetails.TLS)),
			SecretKeyUsername: []byte(smtpDetails.Username),
			SecretKeyPassword: []byte(smtpDetails.Password),
		},
		Type: apiv1.SecretTypeOpaque,
	}
}
