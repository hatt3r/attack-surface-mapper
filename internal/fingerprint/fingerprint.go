package fingerprint

// simple port number to service name mapping for common services
func IdentifyService(port int) string {

	switch port {

	case 21:
		return "FTP"

	case 22:
		return "SSH"

	case 25:
		return "SMTP"

	case 53:
		return "DNS"

	case 80:
		return "HTTP"

	case 110:
		return "POP3"

	case 143:
		return "IMAP"

	case 443:
		return "HTTPS"

	case 3306:
		return "MySQL"

	default:
		return "Unknown"
	}
}
