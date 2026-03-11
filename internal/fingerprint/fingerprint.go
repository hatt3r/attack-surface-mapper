package fingerprint

func IdentifyService(port int) string {

	switch port {

	case 22:
		return "SSH"

	case 80:
		return "HTTP"

	case 443:
		return "HTTPS"

	case 3306:
		return "MySQL"

	default:
		return "Unknown"
	}
}
