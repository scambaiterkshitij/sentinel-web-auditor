# sentinel-web-auditor
Sentinel Web Auditor

Sentinel Web Auditor is a modular web configuration auditing tool built for Linux environments. It inspects HTTPS configuration, TLS negotiation details, DNS resolution, TCP connectivity, and common security header misconfigurations. The tool is designed strictly for defensive analysis and controlled security assessment.

This is not an exploitation framework. It does not perform injection attacks, brute force attempts, denial-of-service activity, or payload execution.

Purpose

The goal of this project is to provide a structured and extensible auditing foundation using a multi-language architecture:
	•	Go handles low-level networking and TLS inspection
	•	Python handles orchestration and risk scoring
	•	JSON is used as the communication layer between components

This separation keeps the design modular, fast, and easy to extend.

Legal Notice

Only use this tool on systems you own or have explicit permission to assess.

Unauthorized scanning or probing may violate local or international laws. The author assumes no responsibility for misuse or damage caused by improper usage. You are solely responsible for ensuring legal compliance.

What It Checks

Network Layer:
	•	DNS resolution
	•	IP address detection
	•	TCP port 443 availability
	•	Basic latency measurement

TLS Layer:
	•	Negotiated TLS version
	•	Certificate issuer information
	•	Certificate expiration date

HTTP Layer:
	•	HTTPS response status code
	•	Response header collection

Risk Engine:
	•	Missing Content-Security-Policy
	•	Missing Strict-Transport-Security
	•	Weak TLS version (1.0 or 1.1)
	•	Closed HTTPS port

A weighted scoring model calculates a final risk score from 0 to 100.

Supported Operating Systems

Designed for Linux systems including:
	•	Ubuntu
	•	Debian
	•	Kali Linux
	•	Arch Linux
	•	Fedora
	•	Parrot OS
	•	Other modern Linux distributions with Go and Python 3

Windows and macOS are not officially supported.

Requirements
	•	Linux operating system
	•	Go 1.18 or newer (recommended)
	•	Python 3.8 or newer

No external Python packages are required. Only the Python standard library is used.

Installing Dependencies

Install Go:

Ubuntu / Debian / Kali:
sudo apt install golang

Arch Linux:
sudo pacman -S go

Fedora:
sudo dnf install golang

Verify:
go version

Verify Python:
python3 –version

Project Structure

sentinel-web-auditor/
netprobe/
probe.go
netprobe (compiled binary)
orchestrator/
main.py
risk_engine.py
shared/
scan_output.json (generated after scan)

Build Instructions

Navigate to the netprobe directory:

cd netprobe

Build the Go binary:

go build -o netprobe

Ensure it is executable if required:

chmod +x netprobe

Running the Tool

Navigate to the orchestrator directory:

cd ../orchestrator

Run a scan:

python3 main.py example.com

Replace example.com with a domain you are authorized to test.

Results will be printed in the terminal and saved to:

shared/scan_output.json

Output

The output JSON contains:
	•	Raw scan data from the Go probe
	•	TLS and HTTP metadata
	•	Calculated risk score
	•	Findings list

This structured output makes it easy to extend the project with dashboards or reporting layers later.

Design Philosophy
	•	Keep networking fast and low-level using Go
	•	Keep scoring flexible using Python
	•	Avoid unnecessary dependencies
	•	Maintain clean separation between probe logic and analysis logic
	•	Make future expansion straightforward

Limitations
	•	Only checks HTTPS (port 443)
	•	No cipher suite enumeration
	•	No HTTP/2 detection
	•	No certificate chain validation
	•	No concurrency scanning

These can be implemented in future revisions.

Disclaimer

This software is provided as-is without warranty. It is intended for educational use, security research, and authorized infrastructure auditing only.

Always test responsibly.


#scambaiterkshitij
