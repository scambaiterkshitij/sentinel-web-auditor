def calculate_risk(data):
    score = 100
    findings = []

    headers = data.get("headers", {})

    if "Content-Security-Policy" not in headers:
        score -= 15
        findings.append("Missing Content-Security-Policy")

    if "Strict-Transport-Security" not in headers:
        score -= 20
        findings.append("Missing HSTS")

    if data.get("tls_version") in ["TLS 1.0", "TLS 1.1"]:
        score -= 25
        findings.append("Weak TLS version")

    if data.get("certificate_expiry"):
        # You can add expiry validation logic later
        pass

    if not data.get("port_open"):
        score -= 40
        findings.append("HTTPS Port Closed")

    if score < 0:
        score = 0

    return {
        "score": score,
        "findings": findings
    }
