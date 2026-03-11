package risk

// simple risk scoring based on service type switch statement
func Score(service string) string {

	switch service {

	case "FTP":
		return "HIGH (unencrypted file transfer)"

	case "SSH":
		return "MEDIUM (secure but often targeted)"

	case "HTTP":
		return "MEDIUM (unencrypted web traffic)"

	case "HTTPS":
		return "LOW"

	case "MySQL":
		return "HIGH (database exposure risk)"

	default:
		return "UNKNOWN"
	}
}
