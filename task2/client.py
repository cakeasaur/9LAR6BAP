import subprocess
import json

data = [1, 2, 3, 4, 5]

proc = subprocess.Popen(
    ["main.exe"],
    stdin=subprocess.PIPE,
    stdout=subprocess.PIPE,
    text=True
)

stdout, _ = proc.communicate(input=json.dumps(data))

print("Ответ:", stdout)