// Copyright 2025 Redpanda Data, Inc.

package tls

import (
	"github.com/redpanda-data/benthos/v4/internal/docs"
)

// FieldSpec returns a spec for a common TLS field.
func FieldSpec() docs.FieldSpec {
	return baseFieldSpec().WithChildren(
		baseCertType("client_certs", "A list of client certificates to use. For each certificate either the fields `cert` and `key`, or `cert_file` and `key_file` should be specified, but not both."),
	)
}

// ServerFieldSpec returns a spec for a common TLS field used by a server.
func ServerFieldSpec() docs.FieldSpec {
	return baseFieldSpec().WithChildren(
		baseCertType("server_certs", "A list of server certificates to use. For each certificate either the fields `cert` and `key`, or `cert_file` and `key_file` should be specified, but not both."),
		docs.FieldBool("require_mutual_tls", "Whether to require mutual TLS authentication. When enabled the server will require a client certificate to be presented during the handshake.").HasDefault(false),
		docs.FieldString(
			"client_root_cas", "An optional root certificate authority to use for checking client certs when mTLS is required. This is a string, representing a certificate chain from the parent trusted root certificate, to possible intermediate signing certificates, to the host certificate.", "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
		).HasDefault("").Secret(),
		docs.FieldString(
			"client_root_cas_file", "An optional path of a root certificate authority file to use for checking client certs when mTLS is required. This is a file, often with a .pem extension, containing a certificate chain from the parent trusted root certificate, to possible intermediate signing certificates, to the host certificate.", "./root_cas.pem",
		).HasDefault(""),
	)
}

func baseFieldSpec() docs.FieldSpec {
	return docs.FieldObject(
		"tls", "Custom TLS settings can be used to override system defaults.",
	).WithChildren(
		docs.FieldBool(
			"enabled", "Whether custom TLS settings are enabled.",
		).HasDefault(false),

		docs.FieldBool(
			"skip_cert_verify", "Whether to skip server side certificate verification.",
		).HasDefault(false),

		docs.FieldBool(
			"enable_renegotiation", "Whether to allow the remote server to repeatedly request renegotiation. Enable this option if you're seeing the error message `local error: tls: no renegotiation`.",
		).AtVersion("3.45.0").Advanced().HasDefault(false),

		docs.FieldString(
			"root_cas", "An optional root certificate authority to use. This is a string, representing a certificate chain from the parent trusted root certificate, to possible intermediate signing certificates, to the host certificate.", "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
		).HasDefault("").Secret(),

		docs.FieldString(
			"root_cas_file", "An optional path of a root certificate authority file to use. This is a file, often with a .pem extension, containing a certificate chain from the parent trusted root certificate, to possible intermediate signing certificates, to the host certificate.", "./root_cas.pem",
		).HasDefault(""),
	).Advanced()
}

func baseCertType(name, description string) docs.FieldSpec {
	return docs.FieldObject(
		name, description,
		[]any{
			map[string]any{
				"cert": "foo",
				"key":  "bar",
			},
		},
		[]any{
			map[string]any{
				"cert_file": "./example.pem",
				"key_file":  "./example.key",
			},
		},
	).Array().WithChildren(
		docs.FieldString("cert", "A plain text certificate to use.").HasDefault(""),
		docs.FieldString("key", "A plain text certificate key to use.").HasDefault("").Secret(),
		docs.FieldString("cert_file", "The path of a certificate to use.").HasDefault(""),
		docs.FieldString("key_file", "The path of a certificate key to use.").HasDefault(""),
		docs.FieldString("password", `A plain text password for when the private key is password encrypted in PKCS#1 or PKCS#8 format. The obsolete `+"`pbeWithMD5AndDES-CBC`"+` algorithm is not supported for the PKCS#8 format.

Because the obsolete pbeWithMD5AndDES-CBC algorithm does not authenticate the ciphertext, it is vulnerable to padding oracle attacks that can let an attacker recover the plaintext.
`, "foo", "${KEY_PASSWORD}").HasDefault("").Secret(),
	).HasDefault([]any{})

}
