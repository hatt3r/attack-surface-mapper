package risk

func Score(service string) string {

	switch service {

	case "SSH":
		return "MEDIUM"

	case "HTTP":
		return "LOW"

	case "MySQL":
		return "HIGH"

	default:
		return "UNKNOWN"
	}
}
