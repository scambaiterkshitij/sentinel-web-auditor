import subprocess
import json
import sys
from risk_engine import calculate_risk
from pathlib import Path

BASE_DIR = Path(__file__).resolve().parent.parent
NETPROBE_PATH = BASE_DIR / "netprobe" / "netprobe"
OUTPUT_PATH = BASE_DIR / "shared" / "scan_output.json"

def run_scan(target):
    try:
        result = subprocess.run(
            [str(NETPROBE_PATH), target],
            capture_output=True,
            text=True,
            timeout=30
        )

        data = json.loads(result.stdout)
        return data

    except Exception as e:
        print("Error running netprobe:", e)
        sys.exit(1)

def save_output(data):
    OUTPUT_PATH.parent.mkdir(exist_ok=True)
    with open(OUTPUT_PATH, "w") as f:
        json.dump(data, f, indent=4)

def main():
    if len(sys.argv) < 2:
        print("Usage: python3 main.py <target>")
        sys.exit(1)

    target = sys.argv[1]
    print(f"[+] Scanning {target} ...")

    data = run_scan(target)
    risk = calculate_risk(data)

    final_output = {
        "scan_data": data,
        "risk_analysis": risk
    }

    save_output(final_output)

    print("\n=== Risk Score ===")
    print("Score:", risk["score"])
    print("Findings:")
    for f in risk["findings"]:
        print("-", f)

    print("\nResults saved to shared/scan_output.json")

if __name__ == "__main__":
    main()
