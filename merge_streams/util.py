import json
import sys


def write_message(d: dict):
    sys.stdout.write(json.dumps(d) + "\n")
