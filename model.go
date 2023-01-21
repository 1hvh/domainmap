package main

type Subdomain struct {
	domain    string
	openPorts []Port
}

type Port struct {
	port   uint16
	isOpen bool
}

type Cert struct {
	IssuerCaId     int    `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	Id             int    `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
	SerialNumber   string `json:"serial_number"`
}
